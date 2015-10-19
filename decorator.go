// Copyright 2015 Yury Kozzyrev. All rights reserved.

package logger
import "strings"


type Decorator struct {
	driver LoggerInterface
}

func NewDecorator(l LoggerInterface) LoggerInterface {
	return &Decorator{l}
}


func (l *Decorator) Debug(format string, args ...interface{}){
	format =  processFormat(format)
	l.driver.Debug(format, args...)
}

func (l *Decorator) Info(format string, args ...interface{}){
	format =  processFormat(format)
	l.driver.Info(format, args...)
}

func (l *Decorator) Notice(format string, args ...interface{}){
	format =  processFormat(format)
	l.driver.Notice(format, args...)
}

func (l *Decorator) Warning(format string, args ...interface{}){
	format =  processFormat(format)
	l.driver.Warning(format, args...)
}

func (l *Decorator) Error(format string, args ...interface{}){
	format =  processFormat(format)
	l.driver.Error(format, args...)
}

func (l *Decorator) Critical(format string, args ...interface{}){
	format =  processFormat(format)
	l.driver.Critical(format, args...)
}

func processFormat (format string) string {
	return strings.Replace(format, "\n", " - ", -1)
}
