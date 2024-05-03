package correction

import (
	"context"

	"github.com/nick1729/resp-api-tmpl/internal/pkg/model"
)

//go:generate easyjson -all -output_filename correction_easyjson.go -pkg

//go:generate mockgen --build_flags test -package=correction -destination=./mock_test.go -source=${GOFILE}

type Service struct {
	repo Repo
}

type Repo interface {
	GetCorrectionByID(ctx context.Context, id string) (*model.Correction, error)
}

func NewService(
	repo Repo,
) *Service {
	return &Service{
		repo,
	}
}
