package log

import (
	"os"
	"strings"

	"github.com/nick1729/resp-api-tmpl/internal/pkg/config"
	"github.com/rs/zerolog"
)

const defaultLoglevel = zerolog.InfoLevel

func Create(cfg config.Log) *zerolog.Logger {
	levels := map[string]zerolog.Level{
		"debug": zerolog.DebugLevel,
		"info":  zerolog.InfoLevel,
		"warn":  zerolog.WarnLevel,
		"error": zerolog.ErrorLevel,
	}

	logLevel, ok := levels[strings.ToLower(cfg.Level)]
	if !ok {
		logLevel = defaultLoglevel
	}

	logger := zerolog.New(os.Stderr).Level(logLevel).With().Timestamp().Logger()

	return &logger
}
