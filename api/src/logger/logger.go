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
}

func New(fileName string) (*Logger, error) {
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	writer := io.MultiWriter(os.Stdout, logFile)
	var logger Logger = Logger{
		DebugLogger: log.New(writer, "[custodire][DEBUG]", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLogger:  log.New(writer, "[custodire][INFO]", log.Ldate|log.Ltime|log.Lshortfile),
		WarnLogger:  log.New(writer, "[custodire][WARN]", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger: log.New(writer, "[custodire][ERROR]", log.Ldate|log.Ltime|log.Lshortfile),
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
