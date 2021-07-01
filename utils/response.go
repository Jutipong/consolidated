package utils

import (
	"consolidated/config"
	"consolidated/enum"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResultData struct {
	ResponseCode   string      `json:"responseCode"`
	ResponseDesc   string      `json:"responseDesc"`
	ResponseDetail interface{} `json:"responseDetail"`
}

func JsonResult(c *gin.Context, ruleId float32, message string, payload interface{}) {
	var res ResultData
	_rule := config.GetRule(ruleId)
	//## Initial Data

	res.ResponseCode = _rule["Code"]
	res.ResponseDetail = payload
	if res.ResponseCode == "0000" {
		res.ResponseDesc = enum.Success
	} else {
		if len(message) == 0 {
			res.ResponseDesc = _rule["Message"]
		} else {
			res.ResponseDesc = message
		}
	}

	//## Next
	c.JSON(http.StatusOK, res)
}

// type ResponseData struct {
// 	ResponseCode   string      `json:"responseCode"`
// 	ResponseDesc   string      `json:"responseDesc"`
// 	ResponseDetail interface{} `json:"responseDetail"`
// }

// func JsonResponse(c *gin.Context, statusCode string, message string, payload interface{}) {
// 	var res ResponseData

// 	//## Initial Data
// 	res.ResponseCode = statusCode
// 	res.ResponseDesc = message
// 	res.ResponseDetail = payload

// 	// if res.ResponseCode == http.StatusOK {
// 	// 	res.ResponseDesc = enum.Success
// 	// } else {
// 	// 	if len(message) == 0 {
// 	// 		res.ResponseDesc = enum.Error
// 	// 	} else {
// 	// 		res.ResponseDesc = message
// 	// 	}
// 	// }

// 	//## Next
// 	c.JSON(http.StatusOK, res)
// }
