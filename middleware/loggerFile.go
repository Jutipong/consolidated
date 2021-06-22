package middleware

import (
	"bytes"
	"consolidated/helper"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinBodyLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		requestURI := c.Request.RequestURI
		method := c.Request.Method
		httpStatus := c.Writer.Status()

		//## Request data
		request := string(bodyBytes)
		//## Response Data
		response := blw.body.String()

		strLog := fmt.Sprintf("RequestURI:%v | Method:%v | Request:%v | Status:%v | Response:%v",
			requestURI, method, request, httpStatus, response)

		switch c.Writer.Status() {
		case http.StatusOK:
			helper.LogInfo(strLog)
		case http.StatusBadRequest, http.StatusUnauthorized:
			helper.LogWarn(strLog)
		default:
			helper.LogError(strLog)
		}
	}
}
