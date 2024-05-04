package order

import (
	"context"
	"errors"
	"time"

	repo "github.com/nick1729/resp-api-tmpl/internal/pkg/repository"
)

type GetReq struct {
	ID string `json:"id"`
}

type GetResp struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	UserID    string     `json:"user_id"`
	Payload   string     `json:"payload"`
}

func (s *Service) Get(ctx context.Context, req *GetReq) (*GetResp, error) {
	// metrics

	order, err := s.repo.GetOrderByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, repo.ErrObjectNotFound) {
			return nil, err // handling.NewNotFound(handling.XXXNotFoundErrorCode, "order not found")
		}

		// ctx.Log().Error().Err(err).Msg("getting order from the repo")

		return nil, err
	}

	resp := GetResp{
		ID:        order.ID,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		UserID:    order.UserID,
		Payload:   order.Payload,
	}

	return &resp, nil
}
