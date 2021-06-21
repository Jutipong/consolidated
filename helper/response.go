package helper

import (
	"consolidated/enum"
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
	res.Data = payload

	if res.Status == http.StatusOK {
		res.Message = enum.Success
	} else {
		if len(message) == 0 {
			res.Message = enum.Error
		} else {
			res.Message = message
		}
	}

	//##Logger
	LoggerResponse(c, res)
	//## Next
	c.JSON(http.StatusOK, res)
}

//##Logger Request
func LoggerRequest(c *gin.Context, payload interface{}) {
	LogInfoReqquest(c, payload)
	c.Next()
}

//##Logger response
func LoggerResponse(c *gin.Context, payload ResponseData) {
	switch payload.Status {
	case http.StatusOK:
		LogInfoResponse(c, payload)
	case http.StatusBadRequest, http.StatusUnauthorized:
		LogWarnResponse(c, payload)
	default:
		LogErrorResponse(c, payload)
	}
}
