package helper

import (
	"consolidated/config"
	"consolidated/model"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

type PlainFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

func (f *PlainFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := fmt.Sprint(entry.Time.Format(f.TimestampFormat))
	return []byte(fmt.Sprintf("%s %s %s\n", f.LevelDesc[entry.Level], timestamp, entry.Message)), nil
}

func SetupLogger() {
	//## Config file
	writer, _ := rotatelogs.New(
		config.Config.LoggerFile.RootPath+"%Y%m%d.log",
		rotatelogs.WithRotationSize((10 * 1048576)), //10MB
		rotatelogs.WithMaxAge(-1),
		// rotatelogs.WithRotationCount(7),
		rotatelogs.WithRotationTime(time.Duration(time.Now().Local().Day())),
	)

	//## Format Logger
	plainFormatter := new(PlainFormatter)
	plainFormatter.TimestampFormat = "[2006-01-02 15:04:05]"
	plainFormatter.LevelDesc = []string{"PANC", "FATL", "ERROR", "WARN", "INFO", "DEBUG"}
	log.SetFormatter(plainFormatter)
	log.SetOutput(writer)
}

//## ==================================================
//## ====================== Info ======================
//## ==================================================

//## Any
func LogInfo(payload interface{}) {
	// b, _ := json.Marshal(payload)
	// fmt.Println(string(b))
	// log.Info(string(b))
	log.Info(payload)
}

//## Request
func LogInfoReqquest(c *gin.Context, payload interface{}) {
	cu := GetUserRequest(c)
	b, _ := json.Marshal(model.LoggerTrasection{
		Type:         "Request",
		TransationId: cu.TransationId,
		Data:         payload,
	})
	fmt.Println(string(b))
	log.Info(string(b))
}

//## Response
func LogInfoResponse(c *gin.Context, payload interface{}) {
	cu := GetUserRequest(c)
	b, _ := json.Marshal(model.LoggerTrasection{
		Type:         "Response",
		TransationId: cu.TransationId,
		Data:         payload,
	})
	fmt.Println(string(b))
	log.Info(string(b))
}

//## ==================================================
//## ====================== Warn ======================
//## ==================================================

//## Any
func LogWarn(payload interface{}) {
	b, _ := json.Marshal(payload)
	fmt.Println(string(b))
	log.Warn(string(b))
}

//## Request
func LogWarnReqquest(c *gin.Context, payload interface{}) {
	cu := GetUserRequest(c)
	b, _ := json.Marshal(model.LoggerTrasection{
		Type:         "Request",
		TransationId: cu.TransationId,
		Data:         payload,
	})
	fmt.Println(string(b))
	log.Warn(string(b))
}

//## Response
func LogWarnResponse(c *gin.Context, payload interface{}) {
	cu := GetUserRequest(c)
	b, _ := json.Marshal(model.LoggerTrasection{
		Type:         "Response",
		TransationId: cu.TransationId,
		Data:         payload,
	})
	fmt.Println(string(b))
	log.Warn(string(b))
}

//## ===================================================
//## ====================== Error ======================
//## ===================================================

//## Any
func LogError(payload interface{}) {
	b, _ := json.Marshal(payload)
	fmt.Println(string(b))
	log.Error(string(b))
}

//## Request
func LogErrorReqquest(c *gin.Context, payload interface{}) {
	cu := GetUserRequest(c)
	b, _ := json.Marshal(model.LoggerTrasection{
		Type:         "Request",
		TransationId: cu.TransationId,
		Data:         payload,
	})
	fmt.Println(string(b))
	log.Error(string(b))
}

//## Response
func LogErrorResponse(c *gin.Context, payload interface{}) {
	cu := GetUserRequest(c)
	b, _ := json.Marshal(model.LoggerTrasection{
		Type:         "Response",
		TransationId: cu.TransationId,
		Data:         payload,
	})
	fmt.Println(string(b))
	log.Error(string(b))
}
