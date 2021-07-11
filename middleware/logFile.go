package middleware

import (
	"bytes"
	"consolidated/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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

func LoggerFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.RequestURI, "login") {
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

			//## Request Data
			//==============================================
			var request string
			var req interface{}
			json.Unmarshal([]byte(string(string(bodyBytes))), &req)
			if req != nil {
				if b, err := json.Marshal(req); err == nil {
					request = string(b)
				}
			}

			//## Response Data
			//==============================================
			response := blw.body.String()

			//## String Format
			strLog := fmt.Sprintf("RequestURI:%v | Method:%v | Request:%v | Status:%v | Response:%v",
				requestURI, method, request, httpStatus, response)

			switch c.Writer.Status() {
			case http.StatusOK:
				utils.LogInfo(strLog)
			case http.StatusBadRequest, http.StatusUnauthorized:
				utils.LogWarn(strLog)
			default:
				utils.LogError(strLog)
			}
		}
		c.Next()
	}
}
