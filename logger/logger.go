package logger

import (
	"log"
	"time"

	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LOG_PATH             = "./logs"
	LOG_FILE_NAME        = "logs.log"
	OUTPUT_LOG_FILE_NAME = "logs.txt"
	FULL_LOG_PATH        = LOG_PATH + "/" + LOG_FILE_NAME
)

var Log *zap.Logger

func init() {
	logFilePath := FULL_LOG_PATH

	err := os.MkdirAll(LOG_PATH, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	config := zap.NewProductionConfig()
	config.OutputPaths = []string{
		logFilePath,
	}
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	Log, err = config.Build()
	if err != nil {
		panic(err)
	}
}
