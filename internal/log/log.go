package log

import (
	"fmt"

	"go.uber.org/zap"
)

// Global non-exported sugar logger.
var sugar *zap.SugaredLogger

// Panic logging functions.
var Panic func(args ...interface{})
var Panicln func(args ...interface{})
var Panicf func(template string, args ...interface{})
var Panicw func(msg string, keysAndValues ...interface{})

// Fatal logging functions.
var Fatal func(args ...interface{})
var Fatalln func(args ...interface{})
var Fatalf func(template string, args ...interface{})
var Fatalw func(msg string, keysAndValues ...interface{})

// Error logging functions.
var Error func(args ...interface{})
var Errorln func(args ...interface{})
var Errorf func(template string, args ...interface{})
var Errorw func(msg string, keysAndValues ...interface{})

// Warn logging functions.
var Warn func(args ...interface{})
var Warnln func(args ...interface{})
var Warnf func(template string, args ...interface{})
var Warnw func(msg string, keysAndValues ...interface{})

// Info logging functions.
var Info func(args ...interface{})
var Infoln func(args ...interface{})
var Infof func(template string, args ...interface{})
var Infow func(msg string, keysAndValues ...interface{})

// Debug logging functions.
var Debug func(args ...interface{})
var Debugln func(args ...interface{})
var Debugf func(template string, args ...interface{})
var Debugw func(msg string, keysAndValues ...interface{})

func init() {
	// Initialize the sugared logger.
	conf := zap.NewProductionConfig()
	conf.Encoding = "console"

	if err := buildLogger(conf); err != nil {
		panic(err)
	}
}

// Function to configure the logger.
func ConfigureLogger(logFile string, logLevel string) error {
	conf := zap.NewProductionConfig()
	conf.Encoding = "console"

	// Configure file output.
	conf.OutputPaths = []string{logFile}

	// Configure log level.
	switch logLevel {
	case "panic":
		conf.Level.SetLevel(zap.PanicLevel)
	case "fatal":
		conf.Level.SetLevel(zap.FatalLevel)
	case "error":
		conf.Level.SetLevel(zap.ErrorLevel)
	case "warn":
		conf.Level.SetLevel(zap.WarnLevel)
	case "info":
		conf.Level.SetLevel(zap.InfoLevel)
	case "debug":
		conf.Level.SetLevel(zap.DebugLevel)
	default:
		return fmt.Errorf("invalid log level: %s", logLevel)
	}

	// Build a logger from the configuration.
	buildLogger(conf)

	return nil
}

// Function to build a logger from a logger configuration and set it as global logger.
func buildLogger(conf zap.Config) error {
	logger, _ := conf.Build()

	sugar = logger.Sugar()
	setExportedFunctions()

	return nil
}

// Function to set the package's exported log functions to a logger.
func setExportedFunctions() {
	Panic = sugar.Panic
	Panicf = sugar.Panicf
	Panicln = sugar.Panicln
	Panicw = sugar.Panicw
	Fatal = sugar.Fatal
	Fatalf = sugar.Fatalf
	Fatalln = sugar.Fatalln
	Fatalw = sugar.Fatalw
	Error = sugar.Error
	Errorf = sugar.Errorf
	Errorln = sugar.Errorln
	Errorw = sugar.Errorw
	Warn = sugar.Warn
	Warnf = sugar.Warnf
	Warnln = sugar.Warnln
	Warnw = sugar.Warnw
	Info = sugar.Info
	Infof = sugar.Infof
	Infoln = sugar.Infoln
	Infow = sugar.Infow
	Debug = sugar.Debug
	Debugf = sugar.Debugf
	Debugln = sugar.Debugln
	Debugw = sugar.Debugw
}
