package logrus

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"loginMicroservice/app/internal/logger"
	"os"
	"runtime"
	"strings"
)

const level = logrus.TraceLevel

type Logger struct {
	loger *logrus.Logger
}

// NewLogger create and set-up logger
func NewLogger() *Logger {
	logger := logrus.New()
	logger.Level = level
	logger.Out = os.Stdout
	logger.SetReportCaller(true)

	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, file, line, _ := runtime.Caller(8)
			return "", fmt.Sprintf("%s:%d", getFileName(file), line)
		},
	}
	logger.SetFormatter(formatter)

	return &Logger{
		loger: logger,
	}
}

func getFileName(path string) string {
	splitPath := strings.Split(path, "/")
	if len(splitPath) < 1 {
		return ""
	}

	return splitPath[len(splitPath)-1]
}

// WithField ...
func (l *Logger) WithField(key string, value interface{}) logger.LoggerWithField {
	return &loggerWithFields{
		loger:  l.loger,
		fields: logger.Fields{key: value},
	}
}

// WithFields ...
func (l *Logger) WithFields(fields logger.Fields) logger.LoggerWithField {
	return &loggerWithFields{
		loger:  l.loger,
		fields: fields,
	}
}

// Debug ...
func (l *Logger) Debug(args ...interface{}) {
	l.loger.Debug(args)
}

// Info ...
func (l *Logger) Info(args ...interface{}) {
	l.loger.Info(args)
}

// Error ...
func (l *Logger) Error(args ...interface{}) {
	l.loger.Error(args)
}
