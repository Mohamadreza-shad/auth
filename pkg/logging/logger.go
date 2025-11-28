package logging

import (
	"github.com/Mohamadreza-shad/auth/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewLogger(cfg *config.Config) (Logger, error) {
	var cores []zapcore.Core
	consoleEncoderCfg := zap.NewProductionEncoderConfig()
	consoleEncoderCfg.TimeKey = "timestamp"
	consoleEncoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleEncoderCfg),
		zapcore.Lock(os.Stdout),
		zap.DebugLevel,
	)
	cores = append(cores, consoleCore)

	core := zapcore.NewTee(cores...)
	z := zap.New(core)
	return &zapLogger{zap: z}, nil
}
