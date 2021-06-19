package helper

import (
	"consolidated/config"
	"consolidated/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func JwtGenerate(payload model.Login) string {
	atClaims := jwt.MapClaims{}

	// Payload begin
	// atClaims["id"] = payload.ID
	// atClaims["username"] = payload.Username
	// atClaims["level"] = payload.Level
	userRequest := &model.UserRequest{
		SystemId:    "S0001",
		EmpCode:     "E0001",
		User:        "U0001",
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
	token, _ := at.SignedString([]byte(config.Config.Server.SecretKey))
	return token
}

func JwtVerify(c *gin.Context) {

	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) == 0 {
		RespondJSON(c, http.StatusBadRequest, "authorization is empty", nil)
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
		//## Encode json
		jsonData := fmt.Sprint(claims["UserRequest"])

		if jsonData == "" {
			RespondJSON(c, http.StatusUnauthorized, err.Error(), nil)
			c.Abort()
			return
		} else {
			var uc model.UserRequest
			json.Unmarshal([]byte(string(jsonData)), &uc)

			//## Add TransationId
			uc.TransationId = uuid.New().String()

			//## obj to json
			b, _ := json.Marshal(uc)

			//## set to gin
			c.Set("UserRequest", string(b))
			c.Next()
		}

	} else {
		RespondJSON(c, http.StatusUnauthorized, err.Error(), nil)
		c.Abort()
		return
	}
}

func GetUserRequest(ctx *gin.Context) model.UserRequest {
	jsonData := ctx.GetString("UserRequest")
	var userRequest model.UserRequest
	json.Unmarshal([]byte(jsonData), &userRequest)
	return userRequest
}
