package handler

import (
	"context"

	"github.com/nick1729/resp-api-tmpl/internal/app/service/correction"
	"github.com/nick1729/resp-api-tmpl/internal/app/service/order"
	repo "github.com/nick1729/resp-api-tmpl/internal/pkg/repository"
)

//go:generate mockgen --build_flags test -package=handler -destination=./mock_test.go -source=${GOFILE}

type resources struct {
	correctionResource correctionResource
	orderResource      orderResource
}

type correctionService interface {
	Get(ctx context.Context, req *correction.GetReq) (*correction.GetResp, error)
}

type orderService interface {
	Get(ctx context.Context, req *order.GetReq) (*order.GetResp, error)
}

func newResources(
	repo *repo.Repository,
) resources {
	return resources{
		correctionResource{
			correction.NewService(
				repo,
			),
		},
		orderResource{
			order.NewService(
				repo,
			),
		},
	}
}
