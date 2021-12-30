package log

import (
	"github.com/workfoxes/calypso/pkg/config"
	"github.com/workfoxes/calypso/pkg/constant"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var (
	L   Logger
	S   *zap.SugaredLogger
	err error
)

func init() {
	env := config.GetDefaultValue(constant.ENV, "DEV")
	if env == "Production" {

	} else {
		//encoderConfig := ecszap.EncoderConfig{
		//	EncodeLevel:    zapcore.CapitalLevelEncoder,
		//	EncodeDuration: zapcore.MillisDurationEncoder,
		//	EncodeCaller:   ecszap.FullCallerEncoder,
		//}
		//core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
		//logger := zap.New(core, zap.AddCaller())
		logger, _ := zap.NewDevelopment()
		L._default = logger
		S = L._default.Sugar()
		if err != nil {
			panic(err)
		}
	}
}

type Logger struct {
	once     sync.Once
	_default *zap.Logger
}

func Init() *zap.Logger {
	logger := zap.NewExample()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	zap.L().Info("replaced zap's global loggers")
	return L._default
}

func NewStackdriverEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "eventTime",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			switch l {
			case zapcore.DebugLevel:
				enc.AppendString("DEBUG")
			case zapcore.InfoLevel:
				enc.AppendString("INFO")
			case zapcore.WarnLevel:
				enc.AppendString("WARNING")
			case zapcore.ErrorLevel:
				enc.AppendString("ERROR")
			case zapcore.DPanicLevel:
				enc.AppendString("CRITICAL")
			case zapcore.PanicLevel:
				enc.AppendString("ALERT")
			case zapcore.FatalLevel:
				enc.AppendString("EMERGENCY")
			}
		},
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {

	L._default.Sugar().Debug(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	L._default.Sugar().Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	L._default.Sugar().Info(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	L._default.Sugar().Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	L._default.Sugar().Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	L._default.Sugar().Warn(args...)
}

// Error logs a message at level Error on the standard logger.
//
// - Something failed but I'm not quitting.
func Error(args ...interface{}) {
	L._default.Sugar().Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
//
// - Calls panic() after logging
func Panic(args ...interface{}) {
	L._default.Sugar().Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
//
// - Calls os.Exit(1) after logging
func Fatal(args ...interface{}) {
	L._default.Sugar().Fatal(args...)
}
