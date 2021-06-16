package service

import (
	"consolidated/config"
	"consolidated/helper"
	"consolidated/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var _secretKey = config.GetSecretKey()

func JwtGenerate(payload model.Login) string {
	atClaims := jwt.MapClaims{}

	// Payload begin
	// atClaims["id"] = payload.ID
	// atClaims["username"] = payload.Username
	// atClaims["level"] = payload.Level
	userRequest := &model.UserRequest{
		EmpCode:     "C0001",
		User:        "Jutipong Subin",
		Permissions: "Admin",
	}

	b, err := json.Marshal(userRequest)
	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf("generate token err: %s", err.Error())
	}

	atClaims["UserRequest"] = string(b)
	atClaims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Payload end

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(_secretKey))
	return token
}

func JwtVerify(c *gin.Context) {

	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) == 0 {
		helper.RespondJSON(c, http.StatusUnauthorized, "authorization is empty", nil)
		c.Abort()
	}

	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(_secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// staffID := fmt.Sprintf("%v", claims["id"])
		// username := fmt.Sprintf("%v", claims["jwt_username"])
		// level := fmt.Sprintf("%v", claims["jwt_level"])
		// c.Set("jwt_staff_id", staffID)
		// c.Set("jwt_username", username)
		// c.Set("jwt_level", level)
		// UserRequest := fmt.Sprintf("%v", claims["UserRequest"])
		c.Set("UserRequest", claims["UserRequest"])
		c.Next()
	} else {
		helper.RespondJSON(c, http.StatusUnauthorized, err.Error(), nil)
		c.Abort()
	}
}
