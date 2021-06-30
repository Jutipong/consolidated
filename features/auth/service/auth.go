package service

import (
	repo "consolidated/features/auth/repository"
	"consolidated/utils"
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) (userAuth repo.Auth, errs interface{}) {

	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return userAuth, ("Unauthorized")
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return userAuth, ("Unauthorized")
	}

	//## Init data
	userAuth.Username = pair[0]
	userAuth.Password = pair[1]

	//## Validate
	errs = utils.ValidField(userAuth, "Username", []utils.Rule{
		{Id: 1},
	})
	if errs != nil {
		return userAuth, errs
	}

	errs = utils.ValidField(userAuth, "Password", []utils.Rule{
		{Id: 1},
	})
	if errs != nil {
		return userAuth, errs
	}
	//## End Validate

	return userAuth, nil
}
