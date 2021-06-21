package helper

import (
	"fmt"
	"io/ioutil"
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

	//##Logger
	LoggerResponse(c, res)

	//## Next
	c.JSON(http.StatusOK, res)
}

//##Logger Request
func LoggerRequest(c *gin.Context) {
	req, err := ioutil.ReadAll(c.Request.Body)
	if err == nil {
		LogInfoReqquest(c, req)
	} else {
		fmt.Println("Request *error: ", err.Error())
	}
	c.Next()
}

//##Logger response
func LoggerResponse(c *gin.Context, payload ResponseData) {

	// fmt.Println(res)

	switch payload.Status {
	case http.StatusOK:
		LogInfoResponse(c, payload)
	case http.StatusBadRequest:
		LogWarnResponse(c, payload)
	case http.StatusUnauthorized:
		LogWarnResponse(c, payload)
	default:
		LogErrorResponse(c, payload)
	}
}
