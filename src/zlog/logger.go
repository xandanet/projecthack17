package zlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Logger *zap.Logger

func InitializeLogger() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	// In development, we use the stdout to log
	if os.Getenv("APP_ENV") == "" || os.Getenv("APP_ENV") == "dev" {
		Logger = consoleLogger(config)
		return
	}

	logPath := os.Getenv("ERROR_LOG_PATH")
	if logPath == "" {
		logPath = "application.log"
	}

	_, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("cannot open log fille: %s", err.Error()))
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10, // MB
		MaxAge:     1,  // days
		MaxBackups: 3,
		Compress:   true,
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		zap.DebugLevel,
	)

	Logger = zap.New(core, zap.AddCaller())
}

func consoleLogger(config zapcore.EncoderConfig) *zap.Logger {
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
