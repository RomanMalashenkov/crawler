package utils

import (
	"log"
	"os"
)

// SetupLogging создает файл для логов и настраивает вывод логов в этот файл
func SetupLogging() {
	logFile, err := os.OpenFile("visited.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Не удалось создать файл для логов:", err)
	}
	log.SetOutput(logFile)
}
