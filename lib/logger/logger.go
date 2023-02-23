package logger

import (
	"github.com/natefinch/lumberjack/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func InitLogger() {
	encoder := getEncoder()
	writerSyncer := getWriterSyncer()
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writerSyncer, zapcore.AddSync(os.Stdout)), zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSyncer() zapcore.WriteSyncer {
	lumberWriteSyncer, _ := lumberjack.NewRoller(
		"./logs/gateway.log",
		10*1024*1024,
		&lumberjack.Options{
			MaxBackups: 10,
			MaxAge:     28 * time.Hour * 24,
			Compress:   false,
		},
	)
	return zapcore.AddSync(lumberWriteSyncer)
}
