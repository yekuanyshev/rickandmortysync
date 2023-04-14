package logging

import "log"

const (
	FATAL = "FATAL"
	ERROR = "ERROR"
	DEBUG = "DEBUG"
	WARN  = "WARN"
	INFO  = "INFO"
)

func Fatal(message string) {
	log.Fatalf("[FATAL] %s\n", message)
}

func Error(message string) {
	log.Println("[ERROR]", message)
}

func Debug(message string) {
	log.Println("[DEBUG]", message)
}

func Warn(message string) {
	log.Println("[WARN]", message)
}

func Info(message string) {
	log.Println("[INFO]", message)
}
