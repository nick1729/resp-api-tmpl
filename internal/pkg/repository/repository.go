package repository

import "github.com/nick1729/resp-api-tmpl/internal/pkg/storage/postgres"

type Repository struct {
	Pool postgres.Service
	// logger
	// metrics
}

func New(pool postgres.Service) *Repository {
	return &Repository{
		Pool: pool,
	}
}
