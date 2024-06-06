package infra

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupZapLogger() *zap.Logger {
	var paths = []string{"stdout", "./logs/dbg.log"}
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = paths
	cfg.EncoderConfig.CallerKey = zapcore.OmitKey
	cfg.EncoderConfig.EncodeTime = customTimeEncoder
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	return logger
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05:000"))
}
