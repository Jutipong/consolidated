package service

import (
	"consolidated/config"
	"consolidated/helper"
	"consolidated/model"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// type jwtCustomClaims struct {
// 	Name  string `json:"name"`
// 	Admin bool   `json:"admin"`
// 	jwt.StandardClaims
// }

var secretKey = config.GetsecretKey()

func JwtSign(payload model.Login) string {
	atClaims := jwt.MapClaims{}

	// Payload begin
	// atClaims["id"] = payload.ID
	atClaims["username"] = payload.Username
	// atClaims["level"] = payload.Level
	atClaims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	// Payload end

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secretKey))
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
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)

		staffID := fmt.Sprintf("%v", claims["id"])
		username := fmt.Sprintf("%v", claims["jwt_username"])
		level := fmt.Sprintf("%v", claims["jwt_level"])
		c.Set("jwt_staff_id", staffID)
		c.Set("jwt_username", username)
		c.Set("jwt_level", level)

		c.Next()
	} else {
		helper.RespondJSON(c, http.StatusUnauthorized, err.Error(), nil)
		c.Abort()
	}
}
