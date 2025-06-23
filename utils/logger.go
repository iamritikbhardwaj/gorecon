package utils

import (
	"log"
)

func Error(message string) {
	log.Println("ERROR" + " " + message)
}

func Info(message string) {
	log.Println("INFO" + " " + message)
}

func Warn(message string) {
	log.Println("Warn" + " " + message)
}
