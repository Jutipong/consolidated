package utils

import (
	"consolidated/enum"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResultData struct {
	ResponseCode int         `json:"responseCode"`
	ResponseDesc string      `json:"responseDesc"`
	Data         interface{} `json:"channelID"`
}

func JsonResult(c *gin.Context, status int, message string, payload interface{}) {
	var res ResultData

	//## Initial Data
	res.ResponseCode = status
	res.Data = payload

	if res.ResponseCode == http.StatusOK {
		res.ResponseDesc = enum.Success
	} else {
		if len(message) == 0 {
			res.ResponseDesc = enum.Error
		} else {
			res.ResponseDesc = message
		}
	}

	//## Next
	c.JSON(http.StatusOK, res)
}

type ResponseData struct {
	ResponseCode string      `json:"responseCode"`
	ResponseDesc string      `json:"responseDesc"`
	Data         interface{} `json:"channelID"`
}

func JsonResponse(c *gin.Context, statusCode string, message string, payload interface{}) {
	var res ResponseData

	//## Initial Data
	res.ResponseCode = statusCode
	res.ResponseDesc = message
	res.Data = payload

	// if res.ResponseCode == http.StatusOK {
	// 	res.ResponseDesc = enum.Success
	// } else {
	// 	if len(message) == 0 {
	// 		res.ResponseDesc = enum.Error
	// 	} else {
	// 		res.ResponseDesc = message
	// 	}
	// }

	//## Next
	c.JSON(http.StatusOK, res)
}
