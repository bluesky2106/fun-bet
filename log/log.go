package log

import (
	"github.com/bluesky2106/fun-bet/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

// InitLogger : print log to log.txt for production env. Otherwise, print log to console. // params: env
func InitLogger(env string) {
	if env != config.EnvProduction {
		logger, _ = zap.NewDevelopment()
	} else {
		writerSyncer := getLogWriter()
		encoder := getEncoder()
		core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
		logger = zap.New(core, zap.AddCaller())
	}

	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}

// GetLogger : get logger
func GetLogger() *zap.Logger {
	if logger == nil {
		InitLogger(config.EnvDevelopment)
	}
	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "log/log.txt",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
