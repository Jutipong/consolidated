package utils

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

func JsonResult(c *gin.Context, status int, message string, payload interface{}) {
	var res ResponseData

	//## Initial Data
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

	//## Next
	c.JSON(http.StatusOK, res)
}
