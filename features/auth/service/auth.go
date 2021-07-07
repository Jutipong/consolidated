package service

import (
	"consolidated/features/auth/model"
	"consolidated/features/auth/repository"
	"consolidated/utils"
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func Login(c *gin.Context, userRequest *model.UserRequest) (code float32, err interface{}) {
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return float32(400), []string{}
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return float32(400), []string{}
	}

	//## Validate Struct
	userLogin := model.UserLogin{
		Username: pair[0],
		Password: pair[1],
	}
	v := validate.Struct(userLogin)
	if !v.Validate() {
		errs := utils.GetFieldNameError(v)
		return float32(1), errs
	}

	//Check in db
	//## 1.Check Username
	userAuth := repository.AuthUser{Username: userLogin.Username}
	if err := userAuth.FindByUserName(); err != nil {
		return float32(400), []string{}
	}
	//## 2.Get Salt in db
	userSalt := repository.AuthSalt{ID: userAuth.ID}
	if err := userSalt.FindById(); err != nil {
		return float32(400), []string{}
	}

	//## Validate password
	if !utils.CheckPasswordHash(userLogin.Password+userSalt.Salt, userAuth.Password) {
		return  float32(400), []string{}
	}


	userRequest.UserId = userAuth.Username
	userRequest.UserName =  userAuth.Username;

	return  float32(0000), nil
}
