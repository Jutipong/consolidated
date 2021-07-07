package utils

import (
	"consolidated/config"
	"consolidated/enum"
	"consolidated/model"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//## No Logger File
func JwtGenerate(payload model.UserRequest) string {
	atClaims := jwt.MapClaims{}
	atClaims[enum.UserRequest] = JsonSerialize(payload)
	atClaims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(config.Config.Server.SecretKey))
	return token
}

func JwtVerify(c *gin.Context) {
	const BEARER_SCHEMA = enum.Bearer
	authHeader := c.GetHeader(enum.Authorization)
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
		return []byte(config.Config.Server.SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set(enum.UserRequest, fmt.Sprint(claims[enum.UserRequest]))
		c.Next()

	} else {
		JsonResult(c, http.StatusUnauthorized, err.Error(), nil)
		c.Abort()
		return
	}
}

func GetUserRequest(ctx *gin.Context) model.UserRequest {
	str := ctx.GetString(enum.UserRequest)
	var userRequest model.UserRequest
	JsonDeserialize(str, &userRequest)
	return userRequest
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
