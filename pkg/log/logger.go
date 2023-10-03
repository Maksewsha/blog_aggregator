package log

import (
	"log"
	"os"
)

type MyLogger interface {
	Info(string)
	Warn(string)
	Error(string)
}

type myLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

const (
	warnPrefix  = "WARN: "
	infoPrefix  = "INFO: "
	errorPrefix = "ERROR: "
)

func NewAppLogger() MyLogger {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0x0666)
	if err != nil {
		log.Fatal(err)
	}

	infoLogger := log.New(file, infoPrefix, log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger := log.New(file, warnPrefix, log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(file, errorPrefix, log.Ldate|log.Ltime|log.Lshortfile)

	return &myLogger{
		infoLogger:  infoLogger,
		warnLogger:  warnLogger,
		errorLogger: errorLogger,
	}
}

func (ml *myLogger) Info(message string) {
	ml.infoLogger.Println(message)
}

func (ml *myLogger) Warn(message string) {
	ml.warnLogger.Println(message)
}

func (ml *myLogger) Error(message string) {
	ml.errorLogger.Println(message)
}
