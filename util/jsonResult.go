package util

import (
	"consolidated/base"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    string      `json:"responseCode"`
	Message string      `json:"responseDesc"`
	Datas   interface{} `json:"responseDetail"`
} //@name Response

func JsonResult(c *gin.Context, code float32, message string, payload interface{}) {
	var res Response
	_rule := base.GetRule(code)
	res.Code = _rule["Code"]
	res.Datas = payload

	if utf8.RuneCountInString(message) == 0 {
		res.Message = _rule["Message"]
	} else {
		res.Message = message
	}

	c.JSON(http.StatusOK, res)
}
