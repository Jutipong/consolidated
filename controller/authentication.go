package controller

import (
	"consolidated/helper"
	"consolidated/model"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthenticatedController struct{}

//## No Get Logger File
func (auth *AuthenticatedController) Login(c *gin.Context) {
	login, err := basicAuth(c)
	if err != nil {
		helper.JsonResult(c, http.StatusUnauthorized,
			"username or password is not authenticated",
			helper.GetErrShouldBind(err))
	} else {
		token := helper.JwtGenerate(login)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func basicAuth(c *gin.Context) (userLogin model.Login, err error) {
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		err := errors.New("Unauthorized")
		return userLogin, err
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return userLogin, errors.New("Unauthorized")
	}
	userLogin.Username = pair[0]
	userLogin.Password = pair[1]
	return userLogin, nil
}
