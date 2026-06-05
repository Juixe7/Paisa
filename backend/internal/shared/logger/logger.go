package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func Init(level string) {
	// Configure global time format
	zerolog.TimeFieldFormat = time.RFC3339

	// Choose console writer for local dev, raw JSON for production
	var logWriter zerolog.LevelWriter
	if os.Getenv("APP_ENV") == "local" {
		logWriter = zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"})
	} else {
		logWriter = zerolog.MultiLevelWriter(os.Stdout)
	}

	// Set global log level
	zerologLevel := zerolog.InfoLevel
	switch level {
	case "debug":
		zerologLevel = zerolog.DebugLevel
	case "info":
		zerologLevel = zerolog.InfoLevel
	case "warn":
		zerologLevel = zerolog.WarnLevel
	case "error":
		zerologLevel = zerolog.ErrorLevel
	}
	zerolog.SetGlobalLevel(zerologLevel)

	Log = zerolog.New(logWriter).With().Timestamp().Logger()
	Log.Info().Msgf("Logger initialized with level: %s", level)
}
