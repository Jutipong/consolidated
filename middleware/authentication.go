package middleware

import (
	"consolidated/service"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("token is empty")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString = authHeader[len(BEARER_SCHEMA):]
		token, err := service.ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func JwtVerifyOut(msg string) {
	fmt.Println(msg)
	// return func(c *gin.Context) {
	// const BEARER_SCHEMA = "Bearer "
	// authHeader := c.GetHeader("Authorization")
	// tokenString := authHeader[len(BEARER_SCHEMA):]

	// token, err := service.NewJWTService().ValidateToken(tokenString)
	// token, err := service.ValidateToken(tokenString)

	// 	if token.Valid {
	// 		claims := token.Claims.(jwt.MapClaims)
	// 		log.Println("Claims[Name]: ", claims["name"])
	// 		log.Println("Claims[Admin]: ", claims["admin"])
	// 		log.Println("Claims[Issuer]: ", claims["iss"])
	// 		log.Println("Claims[IssuedAt]: ", claims["iat"])
	// 		log.Println("Claims[ExpiresAt]: ", claims["exp"])
	// 	} else {
	// 		log.Println(err)
	// 		c.AbortWithStatus(http.StatusUnauthorized)
	// 	}
	// }
}
