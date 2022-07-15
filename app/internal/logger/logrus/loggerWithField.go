package logrus

import (
	"github.com/sirupsen/logrus"
	"loginMicroservice/app/internal/logger"
)

type loggerWithFields struct {
	loger  *logrus.Logger
	fields logger.Fields
}

func (l *loggerWithFields) Debug(args ...interface{}) {
	l.loger.WithFields(logrus.Fields(l.fields)).Debug(args)
}

func (l *loggerWithFields) Info(args ...interface{}) {
	l.loger.WithFields(logrus.Fields(l.fields)).Info(args)
}

func (l *loggerWithFields) Error(args ...interface{}) {
	l.loger.WithFields(logrus.Fields(l.fields)).Error(args)
}
