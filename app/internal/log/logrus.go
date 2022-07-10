package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()

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
	return logger
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
