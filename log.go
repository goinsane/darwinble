package darwinble

import (
	"log"
)

type Logger interface {
	Print(v ...any)
	Printf(format string, v ...any)
	Println(v ...any)
}

var logger Logger = log.Default()
var debugOk = false

// SetLogger replaces the logger with a custom one.
func SetLogger(newLogger Logger) {
	logger = newLogger
}

func SetDebug(enabled bool) {
	debugOk = enabled
}

func logPrint(v ...any) {
	if logger == nil {
		return
	}
	logger.Print(v...)
}

func logPrintf(format string, v ...any) {
	if logger == nil {
		return
	}
	logger.Printf(format, v...)
}

func logDebug(v ...any) {
	if logger == nil || !debugOk {
		return
	}
	logger.Print(v...)
}

func logDebugf(format string, v ...any) {
	if logger == nil || !debugOk {
		return
	}
	logger.Printf(format, v...)
}
