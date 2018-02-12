package logger

import (
	"go.uber.org/zap"
)

type Loggable interface {
	Info(message string)
}

type ZapLogger struct {
	LogProvider *zap.Logger
}

func Logger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}

func (z *ZapLogger) Info(message string) {
	z.LogProvider.Info(message)
}
