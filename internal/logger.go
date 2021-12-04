package internal

import (
	"log"
	"os"
)

func NewLogger(path string) *log.Logger {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return log.New(file, "LOG: ", log.Ldate|log.Llongfile)
}
