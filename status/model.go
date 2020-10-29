package status

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

// OutputType 标准输出类型
type OutputType int

const (
	// OutputTypePrint 屏幕打印
	OutputTypePrint OutputType = 1
	// OutputTypeLog 日志输出
	OutputTypeLog OutputType = 2
)

// StatusOutput 标准输出
type StatusOutput struct {
	Logger *logrus.Logger
}

// InitStatusOutput 日志初始化
func InitStatusOutput(level int) (*StatusOutput, error) {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter)                     //default
	log.Formatter.(*logrus.TextFormatter).DisableColors = false   // remove colors
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	log.Level = logrus.InfoLevel
	if OutputType(level) == OutputTypePrint {
		log.Out = os.Stdout
	} else if OutputType(level) == OutputTypeLog {
		filename, err := GetLogFileName()
		if err != nil {
			return nil, errors.New("获取日志文件名错误！")
		}
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, errors.New("打开日志文件错误！")
		}
		log.Out = file
	} else {
		return nil, errors.New("标准输出类型错误！")
	}
	s := StatusOutput{
		Logger: log,
	}
	return &s, nil
}
