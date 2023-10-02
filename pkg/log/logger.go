package log

import (
	"log"
	"os"
)

type appLogger struct {
	outFile *os.File
}

func NewAppLogger() *appLogger {
	logger := &appLogger{}
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0x0666)
	if err != nil {
		log.Fatal(err)
	}
	logger.outFile = file

	return logger
}

func (*appLogger) Info(message string) {

}
