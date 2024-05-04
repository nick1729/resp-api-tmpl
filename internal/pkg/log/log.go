package log

import (
	"os"
	"strings"

	"github.com/rs/zerolog"

	"github.com/nick1729/resp-api-tmpl/internal/pkg/config"
)

const defaultLoglevel = zerolog.InfoLevel

// Init - initializes logger.
func Init(cfg config.Log) *zerolog.Logger {
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

	resp := zerolog.New(os.Stderr).Level(logLevel).With().Timestamp().Logger()

	return &resp
}
