package logging

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger
	once   sync.Once
)

// Level represents log levels.
type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

// Init initializes the global logger with the specified level.
// This should be called once at application startup.
func Init(level Level) {
	once.Do(func() {
		logger = newLogger(level)
	})
}

// Get returns the global logger. If Init hasn't been called,
// it returns a no-op logger.
func Get() *zap.SugaredLogger {
	if logger == nil {
		// Return a no-op logger if not initialized
		return zap.NewNop().Sugar()
	}
	return logger
}

func newLogger(level Level) *zap.SugaredLogger {
	var zapLevel zapcore.Level
	switch level {
	case LevelDebug:
		zapLevel = zapcore.DebugLevel
	case LevelInfo:
		zapLevel = zapcore.InfoLevel
	case LevelWarn:
		zapLevel = zapcore.WarnLevel
	case LevelError:
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapLevel),
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// Simplify the output format
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	zapLogger, err := config.Build()
	if err != nil {
		// Fall back to nop logger on error
		return zap.NewNop().Sugar()
	}

	return zapLogger.Sugar()
}

// Debug logs a debug message.
func Debug(args ...interface{}) {
	Get().Debug(args...)
}

// Debugf logs a formatted debug message.
func Debugf(template string, args ...interface{}) {
	Get().Debugf(template, args...)
}

// Info logs an info message.
func Info(args ...interface{}) {
	Get().Info(args...)
}

// Infof logs a formatted info message.
func Infof(template string, args ...interface{}) {
	Get().Infof(template, args...)
}

// Warn logs a warning message.
func Warn(args ...interface{}) {
	Get().Warn(args...)
}

// Warnf logs a formatted warning message.
func Warnf(template string, args ...interface{}) {
	Get().Warnf(template, args...)
}

// Error logs an error message.
func Error(args ...interface{}) {
	Get().Error(args...)
}

// Errorf logs a formatted error message.
func Errorf(template string, args ...interface{}) {
	Get().Errorf(template, args...)
}
