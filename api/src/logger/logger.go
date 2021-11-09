package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
	FatalLogger *log.Logger
}

func New(fileName string) (*Logger, error) {
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	writer := io.MultiWriter(os.Stdout, logFile)
	var logFlags int = log.Ldate | log.Ltime | log.Lshortfile
	var logger Logger = Logger{
		DebugLogger: log.New(writer, "[custodire][DEBUG] ", logFlags),
		InfoLogger:  log.New(writer, "[custodire][INFO] ", logFlags),
		WarnLogger:  log.New(writer, "[custodire][WARN] ", logFlags),
		ErrorLogger: log.New(writer, "[custodire][ERROR] ", logFlags),
		FatalLogger: log.New(writer, "[custodire][FATAL] ", logFlags),
	}
	return &logger, nil
}

func (logger *Logger) Debug(msg interface{}) {
	logger.DebugLogger.Println(msg)
}

func (logger *Logger) Info(msg interface{}) {
	logger.InfoLogger.Println(msg)
}

func (logger *Logger) Warn(msg interface{}) {
	logger.WarnLogger.Println(msg)
}

func (logger *Logger) Error(msg interface{}) {
	logger.ErrorLogger.Println(msg)
}

func (logger *Logger) Fatal(msg interface{}) {
	logger.FatalLogger.Println(msg)
}
