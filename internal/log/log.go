package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	if err := buildLogger(generateLoggerConfiguration()); err != nil {
		panic(err)
	}
}

// Function to configure the logger.
func ConfigureLogger(logFile string, logLevel string) error {
	Infow("Configuring logging facility", "log_file", logFile, "level", logLevel)

	cfg := generateLoggerConfiguration()

	// Configure file output.
	cfg.OutputPaths = []string{logFile}

	// Configure log level.
	switch logLevel {
	case "panic":
		cfg.Level.SetLevel(zap.PanicLevel)
	case "fatal":
		cfg.Level.SetLevel(zap.FatalLevel)
	case "error":
		cfg.Level.SetLevel(zap.ErrorLevel)
	case "warn":
		cfg.Level.SetLevel(zap.WarnLevel)
	case "info":
		cfg.Level.SetLevel(zap.InfoLevel)
	case "debug":
		cfg.Level.SetLevel(zap.DebugLevel)
	default:
		return fmt.Errorf("invalid log level: %s", logLevel)
	}

	// Build a logger from the configuration.
	buildLogger(cfg)
	defer sugar.Sync()

	return nil
}

func generateLoggerConfiguration() zap.Config {
	cfg := zap.NewProductionConfig()

	cfg.DisableStacktrace = true
	cfg.Encoding = "console"
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return cfg
}

// Function to build a logger from a logger configuration and set it as global logger.
func buildLogger(cfg zap.Config) error {
	logger, _ := cfg.Build()

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
