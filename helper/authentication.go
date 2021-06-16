package helper

import (
	"consolidated/model"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetUserRequest(ctx *gin.Context) model.UserRequest {
	jsonData := ctx.GetString("UserRequest")
	var userRequest model.UserRequest
	json.Unmarshal([]byte(jsonData), &userRequest)
	return userRequest
}
