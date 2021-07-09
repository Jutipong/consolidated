package utils

import (
	"consolidated/config"
	"encoding/json"
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	logFile "github.com/sirupsen/logrus"
)

// type LoggerTrasection struct {
// 	Type         string
// 	TransationId string
// 	Data         interface{}
// }

type PlainFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

func (f *PlainFormatter) Format(entry *logFile.Entry) ([]byte, error) {
	timestamp := fmt.Sprint(entry.Time.Format(f.TimestampFormat))
	return []byte(fmt.Sprintf("%s %s %s\n", f.LevelDesc[entry.Level], timestamp, entry.Message)), nil
}

func SetupLogger() string {
	//## Config File
	writer, err := rotatelogs.New(
		config.Config.LoggerFile.RootPath+"%Y%m%d.log",
		rotatelogs.WithRotationSize((10 * 1048576)), //10MB
		rotatelogs.WithMaxAge(-1),
		// rotatelogs.WithRotationCount(7),
		rotatelogs.WithRotationTime(time.Duration(time.Now().Local().Day())),
	)

	if err != nil {
		return fmt.Sprintf("LoggerFile Setup fail: %v", err.Error())
	}

	//## Format Logger
	plainFormatter := new(PlainFormatter)
	plainFormatter.TimestampFormat = "[2006-01-02 15:04:05]"
	plainFormatter.LevelDesc = []string{"PANC", "FATL", "ERROR", "WARN", "INFO", "DEBUG"}
	logFile.SetFormatter(plainFormatter)
	logFile.SetOutput(writer)

	return ""
}

//## ==================================================
//## ====================== Info ======================
//## ==================================================

//## Any
func LogInfo(payload interface{}) {
	// b, _ := json.Marshal(payload)
	// fmt.Println(string(b))
	// log.Info(string(b))
	logFile.Info(payload)
}

////## Request
// func LogInfoReqquest(c *gin.Context, payload interface{}) {
// 	cu := GetUserRequest(c)
// 	b, _ := json.Marshal(LoggerTrasection{
// 		Type:         "Request",
// 		TransationId: cu.TransationId,
// 		Data:         payload,
// 	})
// 	fmt.Println(string(b))
// 	log.Info(string(b))
// }

// //## Response
// func LogInfoResponse(c *gin.Context, payload interface{}) {
// 	cu := GetUserRequest(c)
// 	b, _ := json.Marshal(LoggerTrasection{
// 		Type:         "Response",
// 		TransationId: cu.TransationId,
// 		Data:         payload,
// 	})
// 	fmt.Println(string(b))
// 	log.Info(string(b))
// }

//## ==================================================
//## ====================== Warn ======================
//## ==================================================

//## Any
func LogWarn(payload interface{}) {
	b, _ := json.Marshal(payload)
	fmt.Println(string(b))
	logFile.Warn(string(b))
}

// //## Request
// func LogWarnReqquest(c *gin.Context, payload interface{}) {
// 	cu := GetUserRequest(c)
// 	b, _ := json.Marshal(LoggerTrasection{
// 		Type:         "Request",
// 		TransationId: cu.TransationId,
// 		Data:         payload,
// 	})
// 	fmt.Println(string(b))
// 	log.Warn(string(b))
// }

// //## Response
// func LogWarnResponse(c *gin.Context, payload interface{}) {
// 	cu := GetUserRequest(c)
// 	b, _ := json.Marshal(LoggerTrasection{
// 		Type:         "Response",
// 		TransationId: cu.TransationId,
// 		Data:         payload,
// 	})
// 	fmt.Println(string(b))
// 	log.Warn(string(b))
// }

//## ===================================================
//## ====================== Error ======================
//## ===================================================

//## Any
func LogError(payload interface{}) {
	b, _ := json.Marshal(payload)
	fmt.Println(string(b))
	logFile.Error(string(b))
}

// //## Request
// func LogErrorReqquest(c *gin.Context, payload interface{}) {
// 	cu := GetUserRequest(c)
// 	b, _ := json.Marshal(LoggerTrasection{
// 		Type:         "Request",
// 		TransationId: cu.TransationId,
// 		Data:         payload,
// 	})
// 	fmt.Println(string(b))
// 	log.Error(string(b))
// }

// //## Response
// func LogErrorResponse(c *gin.Context, payload interface{}) {
// 	cu := GetUserRequest(c)
// 	b, _ := json.Marshal(LoggerTrasection{
// 		Type:         "Response",
// 		TransationId: cu.TransationId,
// 		Data:         payload,
// 	})
// 	fmt.Println(string(b))
// 	log.Error(string(b))
// }
