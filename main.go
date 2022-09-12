package mylogger

import (
	"log"
)

/*
LogInfo print logingo
*/
func LogInfo(m string) {
	log.Printf("INFO - %v", m)
}
