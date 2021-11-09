package logger

import (
	"os"
	"strings"
	"testing"
)

const (
	testLogFile string = "testLog.txt"
)

func CheckFileContains(fileName string, needle string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err.Error())
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		panic(err.Error())
	}
	return strings.Contains(string(buffer), needle)
}

func CleanTestFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		panic(err.Error())
	}
}

func ClearTestLogFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		panic(err.Error())
	}
	file.Close()
}

func TestLoggerContainsMessage(t *testing.T) {
	ClearTestLogFile(testLogFile)
	logger, err := New(testLogFile)
	if err != nil {
		t.Error("Failed to open log file")
	}
	logger.Debug("Test Log")
	if contains := CheckFileContains(testLogFile, "Test Log"); !contains {
		t.Error("Failed to log message")
	}
	CleanTestFile(testLogFile)
}

func TestLoggerDebugStamp(t *testing.T) {
	ClearTestLogFile(testLogFile)
	logger, err := New(testLogFile)
	if err != nil {
		t.Error("Failed to open log file")
	}
	logger.Debug("Test Log")
	if contains := CheckFileContains(testLogFile, "[DEBUG]"); !contains {
		t.Error("Failed to stamp message with debug")
	}
	CleanTestFile(testLogFile)
}

func TestLoggerInfoStamp(t *testing.T) {
	ClearTestLogFile(testLogFile)
	logger, err := New(testLogFile)
	if err != nil {
		t.Error("Failed to open log file")
	}
	logger.Info("Test Log")
	if contains := CheckFileContains(testLogFile, "[INFO]"); !contains {
		t.Error("Failed to stamp message with info")
	}
	CleanTestFile(testLogFile)
}

func TestLoggerWarnStamp(t *testing.T) {
	ClearTestLogFile(testLogFile)
	logger, err := New(testLogFile)
	if err != nil {
		t.Error("Failed to open log file")
	}
	logger.Warn("Test Log")
	if contains := CheckFileContains(testLogFile, "[WARN]"); !contains {
		t.Error("Failed to stamp message with warn")
	}
	CleanTestFile(testLogFile)
}

func TestLoggerErrorStamp(t *testing.T) {
	ClearTestLogFile(testLogFile)
	logger, err := New(testLogFile)
	if err != nil {
		t.Error("Failed to open log file")
	}
	logger.Error("Test Log")
	if contains := CheckFileContains(testLogFile, "[ERROR]"); !contains {
		t.Error("Failed to stamp message with error")
	}
	CleanTestFile(testLogFile)
}

func TestLoggerFatalStamp(t *testing.T) {
	ClearTestLogFile(testLogFile)
	logger, err := New(testLogFile)
	if err != nil {
		t.Error("Failed to open log file")
	}
	logger.Fatal("Test Log")
	if contains := CheckFileContains(testLogFile, "[FATAL]"); !contains {
		t.Error("Failed to stamp message with fatal")
	}
	CleanTestFile(testLogFile)
}
