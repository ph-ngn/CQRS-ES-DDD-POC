package infrastructure

import (
	"github.com/andyj29/wannabet/internal/application/common"
	"go.uber.org/zap"
)

var infraLogger = NewZapLogger()

type zapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() common.Logger {
	logger, _ := zap.NewProduction()
	return &zapLogger{
		logger: logger.Sugar(),
	}
}

func (zl *zapLogger) Debugf(tmp string, args ...interface{}) {
	zl.logger.Debugf(tmp, args...)
}

func (zl *zapLogger) Infof(tmp string, args ...interface{}) {
	zl.logger.Infof(tmp, args...)
}

func (zl *zapLogger) Warnf(tmp string, args ...interface{}) {
	zl.logger.Warnf(tmp, args...)
}

func (zl *zapLogger) Errorf(tmp string, args ...interface{}) {
	zl.logger.Errorf(tmp, args...)
}

func (zl *zapLogger) Fatalf(tmp string, args ...interface{}) {
	zl.logger.Fatalf(tmp, args...)
}

func (zl *zapLogger) Panicf(tmp string, args ...interface{}) {
	zl.logger.Panicf(tmp, args...)
}
