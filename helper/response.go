package helper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status  int
	Message string
	Data    interface{}
}

func RespondJSON(c *gin.Context, status int, message string, payload interface{}) {
	var res ResponseData

	res.Status = status
	if res.Status == http.StatusOK {
		res.Message = "Success"
	} else {
		res.Message = message
	}
	res.Data = payload

	LogInfoResponse(c, res)
	c.JSON(http.StatusOK, res)
}

func LoggerRequest(c *gin.Context) {
	var res ResponseData
	fmt.Println("req: ",res)
	LogInfoReqquest(c, res)
	c.Next()
}
