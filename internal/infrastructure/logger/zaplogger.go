package logger

import (
	"github.com/andyj29/wannabet/internal/application/common"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	InfraLoggerConfig = Config{
		Layer:       "Infra",
		LogFileName: "infra.log",
	}
	InfraLogger = NewZapLogger(InfraLoggerConfig)
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

type Config struct {
	ServiceName, ServiceHost, Layer, LogFileName string
}

func NewZapLogger(cfg Config) common.Logger {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "@timestamp"
	config.MessageKey = "message"
	config.LevelKey = "log.level"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, _ := os.OpenFile(cfg.LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	logFields := zap.Fields(
		zap.String("service.name", cfg.ServiceName),
		zap.String("service.host", cfg.ServiceHost),
		zap.String("log.layer", cfg.Layer),
	)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel), logFields)
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
