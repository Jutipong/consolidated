package util

import (
	"consolidated/base"
	"consolidated/config"
	"fmt"
	"net/http"
	"time"

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

//## No Logger File
func JwtGenerate(payload UserRequest) string {
	atClaims := jwt.MapClaims{}
	atClaims[base.UserRequest] = JsonSerialize(payload)
	atClaims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(config.Server().SecretKey))
	// token, _ := at.SignedString([]byte(config.Config.Server.SecretKey))
	return token
}

func JwtVerify(c *gin.Context) {
	const BEARER_SCHEMA = base.Bearer
	authHeader := c.GetHeader(base.Authorization)
	if len(authHeader) == 0 {
		JsonResponse(c, http.StatusBadRequest, "authorization is empty", nil)
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
		JsonResponse(c, http.StatusUnauthorized, err.Error(), nil)
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

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
