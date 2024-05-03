package order

import (
	"context"

	"github.com/nick1729/resp-api-tmpl/internal/pkg/model"
)

//go:generate mockgen --build_flags test -package=order -destination=./mock_test.go -source=${GOFILE}

type Service struct {
	repo Repo
}

type Repo interface {
	GetOrderByID(ctx context.Context, id string) (*model.Order, error)
}

func NewService(
	repo Repo,
) *Service {
	return &Service{
		repo,
	}
}
