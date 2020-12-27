package logrus

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// LogFormatter 日志自定义格式
type LogFormatter struct{}

// Format 格式详情
func (s *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	var file string
	var len int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
	}
	msg := fmt.Sprintf("%s %s %s:%d %s\n", timestamp, strings.ToUpper(entry.Level.String()), file, len, entry.Message)
	return []byte(msg), nil
}

// New 封装的初始化logrus
func New() *logrus.Logger {
	LogFileName := "logrus.log"
	LogFilePath, err := os.Getwd()
	if err != nil {
		log.Fatal("logrus init: get LogFilePath failure: \n", err)
	}
	fileName := path.Join(LogFilePath, LogFileName)

	fn, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0666)
	if err != nil {
		log.Fatal("logrus init: open LogFile failure: ", err)
	}

	Logger := logrus.New()
	Logger.Out = fn

	Logger.SetReportCaller(true)
	// Logger.SetFormatter(&logrus.TextFormatter{})
	Logger.SetFormatter(new(LogFormatter)) //注册自定义格式
	Logger.SetLevel(logrus.DebugLevel)
	return Logger
}
