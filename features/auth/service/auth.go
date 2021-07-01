package service

import (
	repo "consolidated/features/auth/repository"
	"consolidated/utils"
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func BasicAuth(c *gin.Context) (userAuth repo.Auth, statusCode float32, errs interface{}) {
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return userAuth, float32(400), nil
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return userAuth, float32(400), nil
	}

	//## Init data
	userAuth.Username = pair[0]
	userAuth.Password = pair[1]

	v := validate.Struct(userAuth)
	if !v.Validate() {
		errs := utils.GetFieldNameError(v)
		return userAuth, float32(1), errs
	}

	return userAuth, float32(0000), nil
}
