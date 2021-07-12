package handler

import (
	"consolidated/base"
	"consolidated/config"
	"consolidated/feature/auth/service"
	"consolidated/util"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userRequest util.UserRequest
	code, err := service.Login(c, &userRequest)
	if err != nil {
		util.JsonResponse(c, code, "", err)
	} else {
		token := jwtGenerate(userRequest)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func jwtGenerate(payload util.UserRequest) string {
	atClaims := jwt.MapClaims{}
	atClaims[base.UserRequest] = util.JsonSerialize(payload)
	atClaims["exp"] = time.Now().Add(time.Hour * time.Duration(config.Server().TokenExpire)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(config.Server().SecretKey))
	return token
}
