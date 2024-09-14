package logger

import (
	"fmt"
	"log"
	"os"
)

func shouldLog(level string) bool {
	currentLevel := os.Getenv("LOG_LEVEL")
	switch currentLevel {
	case "DEBUG":
		return true
	case "INFO":
		return level == "INFO" || level == "ERROR"
	case "ERROR":
		return level == "ERROR"
	default:
		return false
	}
}

func Msg(ErrType, Msg string) {
	if !shouldLog(ErrType) {
		return
	}

	// set file name do you want to log
	filename := "activity.log"

	// create log file with read and write permission to owner and read only to others
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("ERROR: problem to open or create file log")
	}

	// log config format for log file
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	// log registry with UTC time, you need convert to your timezone when look some information
	log.Printf("UTC %s: %s\n", ErrType, Msg)
}
