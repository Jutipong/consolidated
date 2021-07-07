package utils

import (
	"consolidated/config"
	"consolidated/enum"
	"consolidated/features/auth/model"
	repo "consolidated/features/auth/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//## No Logger File
func JwtGenerate(payload model.Auth) string {
	atClaims := jwt.MapClaims{}

	// atClaims["id"] = payload.ID
	// atClaims["username"] = payload.Username
	// atClaims["level"] = payload.Level
	// userRequest := &model.UserRequest{
	userRequest := repo.UserRequest{
		SystemId:    "S0001",
		EmpCode:     "E0001",
		User:        "U0001",
		Permissions: "Admin",
	}

	b, err := json.Marshal(userRequest)
	if err != nil {
		fmt.Println(err)
		msg := fmt.Sprintf("generate token err: %s", err.Error())
		fmt.Println(msg)
		return msg
	}

	atClaims[enum.UserRequest] = string(b)
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

func GetUserRequest(ctx *gin.Context) repo.UserRequest {
	jsonData := ctx.GetString(enum.UserRequest)
	var userRequest repo.UserRequest
	json.Unmarshal([]byte(jsonData), &userRequest)
	return userRequest
}

func getSalt(password *string) {
	*password += "zFA.HNNVQXv]A;^F<'*)xlBl$"
}

func HashPassword(password string) (string, error) {
	getSalt(&password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
