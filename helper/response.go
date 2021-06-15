package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status  int
	Message string
	Data    interface{}
}

func RespondJSON(ctx *gin.Context, status int, message string, payload interface{}) {
	var res ResponseData
	res.Status = status
	if res.Status == http.StatusOK {
		res.Message = "Success"

	} else {
		res.Message = message
	}

	res.Data = payload

	ctx.JSON(200, res)
}
