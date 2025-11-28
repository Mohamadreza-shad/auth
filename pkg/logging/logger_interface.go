package logging

import (
	"github.com/Mohamadreza-shad/auth/pkg/logging/keyval"
	"unsafe"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(msg string, keyAndValues ...keyval.Pair)
	Info(msg string, keyAndValues ...keyval.Pair)
	Warn(msg string, keyAndValues ...keyval.Pair)
	Error(msg string, keyAndValues ...keyval.Pair)
	Fatal(msg string, keyAndValues ...keyval.Pair)
}

type zapLogger struct {
	zap *zap.Logger
}

func (l *zapLogger) Debug(msg string, pairs ...keyval.Pair) {
	var zapFields = *(*[]zapcore.Field)(unsafe.Pointer(&pairs))
	l.zap.Debug(msg, zapFields...)
}
func (l *zapLogger) Info(msg string, pairs ...keyval.Pair) {
	var zapFields = *(*[]zapcore.Field)(unsafe.Pointer(&pairs))
	l.zap.Info(msg, zapFields...)
}
func (l *zapLogger) Warn(msg string, pairs ...keyval.Pair) {
	var zapFields = *(*[]zapcore.Field)(unsafe.Pointer(&pairs))
	l.zap.Warn(msg, zapFields...)
}
func (l *zapLogger) Error(msg string, pairs ...keyval.Pair) {
	var zapFields = *(*[]zapcore.Field)(unsafe.Pointer(&pairs))
	l.zap.Error(msg, zapFields...)
}
func (l *zapLogger) Fatal(msg string, pairs ...keyval.Pair) {
	var zapFields = *(*[]zapcore.Field)(unsafe.Pointer(&pairs))
	l.zap.Fatal(msg, zapFields...)
}
