package logger

import (
"fmt"
"log"
"os"
)

// Msg Message Log module for logs.
// to use this log.Msg("CRITICAL", "cannot create db")
func Msg(ErrType, Msg string) {

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