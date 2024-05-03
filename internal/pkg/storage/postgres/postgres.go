package postgres

import (
	"context"
	"time"

	"github.com/nick1729/resp-api-tmpl/internal/pkg/config"
	"github.com/nick1729/resp-api-tmpl/internal/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

const pingTimeout = 2 * time.Second

func New(ctx context.Context, cfg config.Postgres, logger *zerolog.Logger) (Service, error) {
	logger.Info().Msg("opening database")

	dsn := cfg.ConnString()

	globalLogLevel := log.Logger.GetLevel()
	dbLogger := NewDBLogger(log.Level(globalLogLevel))

	dialector := postgres.Open(dsn)

	gormCfg := gorm.Config{
		Logger: dbLogger,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}

	gormDB, err := gorm.Open(dialector, &gormCfg)
	if err != nil {
		return Service{}, errors.Wrap(err, "opening database connection")
	}

	s := Service{
		DB: gormDB,
	}

	closeConnOnReturn := true

	defer func() {
		if !closeConnOnReturn {
			return
		}

		err := s.Close()
		if err != nil {
			log.Debug().AnErr("cause", err).Msg("cannot close database connection")
		}
	}()

	sqlDB, err := gormDB.DB()
	if err != nil {
		return Service{}, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(ctx, pingTimeout)
	defer cancel()

	err = sqlDB.PingContext(ctx)
	if err != nil {
		return Service{}, errors.Wrap(err, "checking database connection")
	}

	log.Trace().Msg("enabling write-ahead-log journal mode")

	closeConnOnReturn = false

	return s, nil
}

func (s *Service) Close() error {
	database, err := s.DB.DB()
	if err != nil {
		return err
	}

	return database.Close()
}
