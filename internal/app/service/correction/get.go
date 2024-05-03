package correction

import (
	"context"
	"errors"

	repo "github.com/nick1729/resp-api-tmpl/internal/pkg/repository"
)

type GetReq struct {
	ID string `json:"id"`
}

type GetResp struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Payload   string `json:"payload"`
	IsError   bool   `json:"is_error"`
}

func (s *Service) Get(ctx context.Context, req *GetReq) (*GetResp, error) {
	// metrics

	order, err := s.repo.GetCorrectionByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, repo.ErrObjectNotFound) {
			return nil, err // handling.NewNotFound(handling.XXXNotFoundErrorCode, "order not found")
		}

		//ctx.Log().Error().Err(err).Msg("getting order from the repo")

		return nil, err
	}

	resp := GetResp{
		ID: order.ID,
		//CreatedAt: json.ConvertTime(order.CreatedAt),
		//UpdatedAt: json.ConvertTime(order.UpdatedAt),
		Payload: order.Payload,
		IsError: order.IsError,
	}

	return &resp, nil
}
