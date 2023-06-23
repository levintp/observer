package logging

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func Init() {
	Logger = log.New()

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Logger.SetOutput(os.Stdout)

	Logger.SetLevel(log.DebugLevel)
}
