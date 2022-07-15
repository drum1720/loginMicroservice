package logrus

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"loginMicroservice/app/internal/logger"
	"os"
	"runtime"
	"strings"
)

type Logger struct {
	loger *logrus.Logger
}

func NewLogger() *Logger {
	logger := logrus.New()

	logger.Out = os.Stdout
	logger.SetReportCaller(true)

	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}

	logger.SetFormatter(formatter)

	return &Logger{
		loger: logger,
	}
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func (l *Logger) WithField(key string, value interface{}) logger.LoggerWithField {
	return &loggerWithFields{
		loger:  l.loger,
		fields: logger.Fields{key: value},
	}
}

func (l *Logger) WithFields(fields logger.Fields) logger.LoggerWithField {
	return &loggerWithFields{
		loger:  l.loger,
		fields: fields,
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.loger.Debug(args)
}

func (l *Logger) Info(args ...interface{}) {
	l.loger.Info(args)
}

func (l *Logger) Error(args ...interface{}) {
	l.loger.Error(args)
}
