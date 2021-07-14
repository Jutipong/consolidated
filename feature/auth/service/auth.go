package service

import (
	"consolidated/base"
	"consolidated/feature/auth/repository"
	"consolidated/util"
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func Login(c *gin.Context, userRequest *util.UserRequest, errs *[]string) float32 {
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return base.BadRequest
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return base.BadRequest
	}

	//## Validate Struct
	userLogin := util.UserLogin{
		Username: pair[0],
		Password: pair[1],
	}
	//config validate
	validate.Config(func(opt *validate.GlobalOption) { opt.StopOnError = false })
	v := validate.Struct(userLogin)
	if !v.Validate() {
		*errs = append(*errs, util.GetFieldNameError(v)...)
		return float32(1)
	}

	//Check in db
	//## 1.Check Username
	userAuth := repository.AuthUser{Username: userLogin.Username}
	if err := userAuth.FindByUserName(); err != nil {
		return base.Unauthorized
	}
	//## 2.Get Salt in db
	userSalt := repository.AuthSalt{ID: userAuth.ID}
	if err := userSalt.FindById(); err != nil {
		return base.Unauthorized
	}

	//## Validate password
	if !util.CheckPasswordHash(userLogin.Password+userSalt.Salt, userAuth.Password) {
		return base.Unauthorized
	}

	userRequest.UserId = userAuth.Username
	userRequest.UserName = userAuth.Username

	return base.Successfully
}
