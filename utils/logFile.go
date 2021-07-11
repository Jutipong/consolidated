package utils

import (
	"consolidated/config"
	"fmt"
	"log"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	logFile "github.com/sirupsen/logrus"
)

type plainFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

func (f *plainFormatter) Format(entry *logFile.Entry) ([]byte, error) {
	timestamp := fmt.Sprint(entry.Time.Format(f.TimestampFormat))
	return []byte(fmt.Sprintf("%s %s %s\n", f.LevelDesc[entry.Level], timestamp, entry.Message)), nil
}

func SetupLogger() {

	//## Config File
	writer, err := rotatelogs.New(
		// config.Config.LoggerFile.RootPath+"%Y%m%d.log",
		config.LogFile().RootPath+"%Y%m%d.log",
		rotatelogs.WithRotationSize((10 * 1048576)), //10MB
		rotatelogs.WithMaxAge(-1),
		// rotatelogs.WithRotationCount(7),
		rotatelogs.WithRotationTime(time.Duration(time.Now().Local().Day())),
	)

	if err != nil {
		panic(err)
	}

	//## Format Logger
	plainFormatter := new(plainFormatter)
	plainFormatter.TimestampFormat = "[2006-01-02 15:04:05]"
	plainFormatter.LevelDesc = []string{"PANC", "FATL", "ERROR", "WARN", "INFO", "DEBUG"}
	logFile.SetFormatter(plainFormatter)
	logFile.SetOutput(writer)
}

func LogInfo(payload interface{}) {
	log.Println(payload)
	logFile.Info(payload)
}

func LogWarn(payload interface{}) {
	log.Println(payload)
	logFile.Warn(payload)
}

func LogError(payload interface{}) {
	log.Println(payload)
	logFile.Error(payload)
}
