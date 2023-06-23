package common

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// Function to initialize the logger.
func InitLogger() {
	// Set initial logger configuration.
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

// Function to configure the logger.
func ConfigureLogger(logFile string, logLevel string) error {
	// Configure file output.
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.FileMode(0o0664))
	if err != nil {
		return fmt.Errorf("failed to open log file %s: %v", logFile, err)
	}
	log.SetOutput(file)

	// Configure log level.
	switch logLevel {
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	default:
		return fmt.Errorf("invalid log level: %s", logLevel)
	}

	return nil
}
