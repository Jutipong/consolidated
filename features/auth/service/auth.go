package service

import (
	"consolidated/config"
	repo "consolidated/features/auth/repository"
	"consolidated/utils"
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func BasicAuth(c *gin.Context) (userAuth repo.Auth, statusCode string, message string, errs interface{}) {
	_rule := config.GetRule(1)
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return userAuth, _rule["ErrorCode"], _rule["Message"], ("Unauthorized")
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return userAuth, _rule["ErrorCode"], _rule["Message"], ("Unauthorized")
	}

	//## Init data
	userAuth.Username = pair[0]
	userAuth.Password = pair[1]

	v := validate.Struct(userAuth)
	if !v.Validate() {
		errs := utils.GetFieldNameError(v)
		return userAuth, _rule["ErrorCode"], _rule["Message"], errs
	}

	// //## Validate
	// errs = utils.ValidField(userAuth, "Username", []utils.Rule{
	// 	{Id: 1},
	// })
	// if errs != nil {
	// 	return userAuth, errs
	// }

	// errs = utils.ValidField(userAuth, "Password", []utils.Rule{
	// 	{Id: 1},
	// })
	// if errs != nil {
	// 	return userAuth, errs
	// }
	// //## End Validate

	return userAuth, "", "", nil
}
