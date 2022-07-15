package logger

type Logger interface {
	WithField(key string, value interface{}) *LoggerWithField
	WithFields(fields Fields) *LoggerWithField

	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
}

type LoggerWithField interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
}
