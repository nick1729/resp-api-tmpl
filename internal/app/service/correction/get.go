package correction

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
	Payload   string     `json:"payload"`
	IsError   bool       `json:"is_error"`
}

func (s *Service) Get(ctx context.Context, req *GetReq) (*GetResp, error) {
	// metrics

	order, err := s.repo.GetCorrectionByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, repo.ErrObjectNotFound) {
			return nil, err
		}

		return nil, err
	}

	resp := GetResp{
		ID:        order.ID,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		Payload:   order.Payload,
		IsError:   order.IsError,
	}

	return &resp, nil
}
