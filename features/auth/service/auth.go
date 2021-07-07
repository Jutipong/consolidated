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

func Login(c *gin.Context) (userLogin model.Auth, code float32, err interface{}) {
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return userLogin, float32(400), []string{}
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return userLogin, float32(400), []string{}
	}

	//## Init data
	userLogin.Username = pair[0]
	userLogin.Password = pair[1]

	//## Validate Struct
	v := validate.Struct(userLogin)
	if !v.Validate() {
		errs := utils.GetFieldNameError(v)
		return userLogin, float32(1), errs
	}

	//Check in db
	//## 1.Check Username
	userAuth := repository.AuthUser{Username: userLogin.Username}
	if err := userAuth.FindByUserName(); err != nil {
		return userLogin, float32(400), []string{}
	}
	//## 2.Get Salt in db
	userSalt := repository.AuthSalt{ID: userAuth.ID}
	if err := userSalt.FindById(); err != nil {
		return userLogin, float32(400), []string{}
	}

	//## Validate password
	if !utils.CheckPasswordHash(userLogin.Password+userSalt.Salt, userAuth.Password) {
		return userLogin, float32(400), []string{}
	}

	return userLogin, float32(0000), nil
}
