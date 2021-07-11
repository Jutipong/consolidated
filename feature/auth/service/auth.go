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

func Login(c *gin.Context, userRequest *util.UserRequest) (code float32, err interface{}) {
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return base.BadRequest, []string{}
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return base.BadRequest, []string{}
	}

	//## Validate Struct
	userLogin := util.UserLogin{
		Username: pair[0],
		Password: pair[1],
	}
	v := validate.Struct(userLogin)
	if !v.Validate() {
		errs := util.GetFieldNameError(v)
		return float32(1), errs
	}

	//Check in db
	//## 1.Check Username
	userAuth := repository.AuthUser{Username: userLogin.Username}
	if err := userAuth.FindByUserName(); err != nil {
		return base.Unauthorized, []string{}
	}
	//## 2.Get Salt in db
	userSalt := repository.AuthSalt{ID: userAuth.ID}
	if err := userSalt.FindById(); err != nil {
		return base.Unauthorized, []string{}
	}

	//## Validate password
	if !util.CheckPasswordHash(userLogin.Password+userSalt.Salt, userAuth.Password) {
		return base.Unauthorized, []string{}
	}

	userRequest.UserId = userAuth.Username
	userRequest.UserName = userAuth.Username

	return base.Successfully, nil
}
