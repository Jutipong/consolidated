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

	//## Initial data
	res.Status = status
	if res.Status == http.StatusOK {
		res.Message = "Success"
	} else {
		res.Message = message
	}
	res.Data = payload

	//## Logger response
	switch status {
	case http.StatusOK:
		LogInfoResponse(c, res)
	case http.StatusBadRequest:
		LogWarnResponse(c, res)
	case http.StatusUnauthorized:
		LogWarnResponse(c, res)
	default:
		LogErrorResponse(c, res)
	}

	//## Next
	c.JSON(http.StatusOK, res)
}

func LoggerRequest(c *gin.Context) {
	var res ResponseData
	fmt.Println("req: ", res)
	LogInfoReqquest(c, res)
	c.Next()
}
