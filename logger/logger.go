package logger

//
// microservices => logger => logger.go
//

import (
	"errors"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log              *zap.Logger
	customTimeFormat string
)

type Config struct {
	LogLevel    zapcore.Level
	Development bool
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(customTimeFormat))
}

func Init(config Config) (*zap.Logger, error) {

	if Log != nil {
		Log.Fatal("Logger já inicializado uma vez, não há necessidade de fazê-lo várias vezess")
		return nil, errors.New("logger já inicializado uma vez")
	}

	globalLevel := config.LogLevel
	logTimeFormat := "Jan 2 15:04:05"

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= globalLevel && lvl < zapcore.ErrorLevel
	})
	consoleInfos := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	var useCustomTimeFormat bool

	var ecfg zapcore.EncoderConfig
	if config.Development {
		ecfg = zap.NewDevelopmentEncoderConfig()
	} else {
		ecfg = zap.NewProductionEncoderConfig()
	}

	if len(logTimeFormat) > 0 {
		customTimeFormat = logTimeFormat
		ecfg.EncodeTime = customTimeEncoder
		useCustomTimeFormat = true
	}
	consoleEncoder := zapcore.NewJSONEncoder(ecfg)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleInfos, lowPriority),
	)

	Log = zap.New(core)
	zap.ReplaceGlobals(Log)
	zap.RedirectStdLog(Log)

	if !useCustomTimeFormat {
		Log.Warn("formato de hora para logger não é fornecido - usando o padrão zap")
	}

	return Log, nil
}
