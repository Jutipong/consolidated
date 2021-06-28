package controller

import (
	repo "consolidated/features/authentication/repository"
	"consolidated/utils"
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
		utils.JsonResult(c, http.StatusUnauthorized,
			"username or password is not authenticated",
			utils.GetErrShouldBind(err))
	} else {
		token := utils.JwtGenerate(login)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func basicAuth(c *gin.Context) (userAuth repo.Auth, err error) {
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		err := errors.New("Unauthorized")
		return userAuth, err
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return userAuth, errors.New("Unauthorized")
	}
	userAuth.Username = pair[0]
	userAuth.Password = pair[1]
	return userAuth, nil
}
