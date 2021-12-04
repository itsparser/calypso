package log

import (
	"os"
	"sync"

	_logger "github.com/sirupsen/logrus"
)

var (
	L = new(Logger)
)

type Logger struct {
	once     sync.Once
	_default *_logger.Logger
}

func Init() *_logger.Logger {
	L.once.Do(func() {
		L._default = &_logger.Logger{
			Out:          os.Stderr,
			Formatter:    new(_logger.TextFormatter),
			Hooks:        make(_logger.LevelHooks),
			Level:        _logger.TraceLevel,
			ExitFunc:     os.Exit,
			ReportCaller: false,
		}
	})
	return L._default
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {

	Init().Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	Init().Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	Init().Print(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	Init().Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	Init().Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	Init().Warning(args...)
}

// Error logs a message at level Error on the standard logger.
//
// - Something failed but I'm not quitting.
func Error(args ...interface{}) {
	Init().Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
//
// - Calls panic() after logging
func Panic(args ...interface{}) {
	Init().Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
//
// - Calls os.Exit(1) after logging
func Fatal(args ...interface{}) {
	Init().Fatal(args...)
}
