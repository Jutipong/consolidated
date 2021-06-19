package helper

import (
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
	path := "./logs/"
	//## Config file
	writer, _ := rotatelogs.New(
		path+"%Y%m%d.log",
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

//## Any
func LogInfo(objs interface{}) {
	log.Info(objs)
}

//## Request
func LogInfoReqquest(c *gin.Context, objs interface{}) {
	cu := GetUserRequest(c)
	b, _ := json.Marshal(
		model.LoggerTrasection{
			Type:         "Request",
			TransationId: cu.TransationId,
			Data:         objs,
		})
	log.Info(string(b))
}

//## Response
func LogInfoResponse(c *gin.Context, objs interface{}) {
	cu := GetUserRequest(c)
	b, _ := json.Marshal(
		model.LoggerTrasection{
			Type:         "Response",
			TransationId: cu.TransationId,
			Data:         objs,
		})
	log.Info(string(b))
}

func LogError(objs interface{}) {
	log.Error(objs)
}

// func main() {
// 	i := 0
// 	for {
// 		i = i + 1
// 		log.Info(fmt.Sprint("Count: ", i))
// 		// time.Sleep(time.Duration(1) * time.Second)
// 		log.Error(fmt.Sprint("Count: ", i+2))
// 		log.Warning(fmt.Sprint("Count: ", i+5))
// 		time.Sleep(time.Duration(1) * time.Second)
// 	}
// }
