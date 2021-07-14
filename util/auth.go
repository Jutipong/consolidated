package util

import (
	"consolidated/base"
	"consolidated/config"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type UserRequest struct {
	UserId   string
	UserName string
}

func JwtVerify(c *gin.Context) {
	const BEARER_SCHEMA = base.Bearer
	authHeader := c.GetHeader(base.Authorization)
	if len(authHeader) == 0 {
		JsonResult(c, http.StatusBadRequest, "authorization is empty", nil)
		c.Abort()
		return
	}

	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Server().SecretKey), nil
		// return []byte(config.Config.Server.SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set(base.UserRequest, fmt.Sprint(claims[base.UserRequest]))
		c.Next()

	} else {
		JsonResult(c, http.StatusUnauthorized, err.Error(), nil)
		c.Abort()
		return
	}
}

func GetUserRequest(ctx *gin.Context) UserRequest {
	str := ctx.GetString(base.UserRequest)
	var userRequest UserRequest
	JsonDeserialize(str, &userRequest)
	return userRequest
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
