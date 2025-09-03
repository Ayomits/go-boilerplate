package logger

import "log"

func NewLogger() *log.Logger {
	return log.Default()
}

var Log = NewLogger()