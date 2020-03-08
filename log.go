package goutils

import (
	"log"
	"os"
)

func NewLogger(logFilePath string, prefix string) *log.Logger {
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	// 创建一个日志对象
	logger := log.New(logFile, prefix, log.LstdFlags)
	return logger
}

func LogInfo(logFilePath string, msg string) {
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	// 创建一个日志对象
	debugLog := log.New(logFile, "", log.LstdFlags)
	debugLog.Println(msg)
}

func FileLogPrintln(logFilePath string, v ...interface{}) {
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	// 创建一个日志对象
	debugLog := log.New(logFile, "", log.LstdFlags)
	debugLog.Println(v...)
}

func FileLogPrintf(logFilePath string, format string, v ...interface{}) {
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	// 创建一个日志对象
	debugLog := log.New(logFile, "", log.LstdFlags)
	debugLog.Printf(format, v...)
}
