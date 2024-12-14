package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

var environment = os.Getenv("ENVIRONMENT")

func Error(message string) {
	printLog(message, "Error")
}

func Info(message string) {
	printLog(message, "Info")
}

func Debug(message string) {
	printLog(message, "Debug")
}

func Warning(message string) {
	printLog(message, "Warning")
}

func printLog(message string, logType string) {
	fmt.Printf("%v - [%s] %s: %s\n", time.Now().Format(time.RFC3339), environment, logType, message)
}
