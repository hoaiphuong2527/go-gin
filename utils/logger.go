package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	logFile, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	fileWriter := zapcore.AddSync(logFile)
	core := zapcore.NewCore(fileEncoder, fileWriter, zapcore.InfoLevel)

	Logger = zap.New(core)
	defer Logger.Sync()
}

func GetLogger() *zap.Logger {
	return Logger
}
