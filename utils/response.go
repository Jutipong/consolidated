package utils

import (
	"consolidated/base"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ResponseCode   string      `json:"responseCode"`
	ResponseDesc   string      `json:"responseDesc"`
	ResponseDetail interface{} `json:"responseDetail"`
}

func JsonResult(c *gin.Context, ruleId float32, message string, payload interface{}) {
	var res Response
	_rule := base.GetRule(ruleId)

	res.ResponseCode = _rule["Code"]
	res.ResponseDetail = payload

	if len(message) == 0 {
		res.ResponseDesc = _rule["Message"]
	} else {
		res.ResponseDesc = message
	}

	c.JSON(http.StatusOK, res)
}
