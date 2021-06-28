package service

import (
	repo "consolidated/features/auth/repository"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) (userAuth repo.Auth, err error) {
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
