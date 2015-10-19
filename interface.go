// Copyright 2015 Yury Kozyrev
// Proprietary license.
package logger


type LoggerInterface interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Notice(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Critical(format string, args ...interface{})
}
