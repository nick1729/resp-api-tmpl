package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type dbLogger struct {
	zlog *zerolog.Logger

	IgnoreRecordNotFoundError bool
	SlowThreshold             time.Duration
}

var _ gormlogger.Interface = (*dbLogger)(nil)

var logLevelMap = map[gormlogger.LogLevel]zerolog.Level{
	gormlogger.Silent: zerolog.Disabled,
	gormlogger.Error:  zerolog.ErrorLevel,
	gormlogger.Warn:   zerolog.WarnLevel,
	gormlogger.Info:   zerolog.InfoLevel,
}

func gormToZlogLevel(logLevel gormlogger.LogLevel) zerolog.Level {
	zlogLevel, ok := logLevelMap[logLevel]
	if !ok {
		return zerolog.DebugLevel
	}

	return zlogLevel
}

func NewDBLogger(logger zerolog.Logger) *dbLogger {
	return &dbLogger{
		zlog: &logger,
	}
}

func (l *dbLogger) Info(ctx context.Context, msg string, args ...any) {
	l.logEvent(zerolog.InfoLevel, msg, args)
}

func (l *dbLogger) Warn(ctx context.Context, msg string, args ...any) {
	l.logEvent(zerolog.WarnLevel, msg, args)
}

func (l *dbLogger) Error(ctx context.Context, msg string, args ...any) {
	l.logEvent(zerolog.ErrorLevel, msg, args)
}

func (l *dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	zlogLevel := l.zlog.GetLevel()
	if zlogLevel == zerolog.Disabled {
		return
	}

	elapsed := time.Since(begin)
	logCtx := l.zlog.With().CallerWithSkipFrameCount(5)

	buildLogger := func() (loggerPtr *zerolog.Logger, sql string) {
		sql, rows := fc()
		logCtx = logCtx.Err(err)
		if rows >= 0 {
			logCtx = logCtx.Int64("rowsAffected", rows)
		}
		logger := logCtx.Logger()
		return &logger, sql
	}

	switch {
	case err != nil && zlogLevel <= zerolog.ErrorLevel && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		logger, sql := buildLogger()
		logger.Error().Msg(sql)

	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && zlogLevel <= zerolog.WarnLevel:
		logger, sql := buildLogger()
		logger.Warn().
			Str("sql", sql).
			Dur("elapsed", elapsed).
			Dur("slowThreshold", l.SlowThreshold).
			Msg("slow database query")

	case zlogLevel <= zerolog.TraceLevel:
		logger, sql := buildLogger()
		logger.Trace().Msg(sql)
	}
}

func (l dbLogger) logEvent(level zerolog.Level, msg string, args ...any) {
	if l.zlog.GetLevel() > level {
		return
	}

	logger := l.logger(args)
	logger.WithLevel(level).Msg(msg)
}

func (l *dbLogger) LogMode(logLevel gormlogger.LogLevel) gormlogger.Interface {
	childLogger := l.zlog.Level(gormToZlogLevel(logLevel))
	newlogger := *l
	newlogger.zlog = &childLogger

	return &newlogger
}

func (l dbLogger) logger(args ...any) zerolog.Logger {
	logCtx := l.zlog.With()
	for idx, arg := range args {
		logCtx.Interface(fmt.Sprintf("arg%d", idx), arg)
	}

	return logCtx.Logger()
}
