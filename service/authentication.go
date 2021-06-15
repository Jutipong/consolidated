package service

import (
	"consolidated/config"
	"consolidated/model"
	"fmt"
	"net/http"
	"strings"
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
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	fmt.Println(tokenString)
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
		c.JSON(http.StatusOK, gin.H{"result": "nok", "message": "invalid token", "error": err})
		c.Abort()
	}
}

// func JwtVerify() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var tokenString string

// 		const BEARER_SCHEMA = "Bearer "
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			log.Println("token is empty")
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		tokenString = authHeader[len(BEARER_SCHEMA):]
// 		token, err := validateToken(tokenString)
// 		if token.Valid {
// 			claims := token.Claims.(jwt.MapClaims)
// 			log.Println("Claims[Name]: ", claims["name"])
// 			log.Println("Claims[Admin]: ", claims["admin"])
// 			log.Println("Claims[Issuer]: ", claims["iss"])
// 			log.Println("Claims[IssuedAt]: ", claims["iat"])
// 			log.Println("Claims[ExpiresAt]: ", claims["exp"])
// 		} else {
// 			log.Println(err)
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}
// 	}
// }

// func validateToken(tokenString string) (*jwt.Token, error) {
// 	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Signing method validation
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(getSecretKey()), nil
// 	})
// }
