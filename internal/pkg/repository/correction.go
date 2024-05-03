package repository

import (
	"context"

	"github.com/nick1729/resp-api-tmpl/internal/pkg/errors"
	"github.com/nick1729/resp-api-tmpl/internal/pkg/model"
	"gorm.io/gorm"
)

func (r *Repository) GetCorrectionByID(ctx context.Context, id string) (*model.Correction, error) {
	// metrics

	var resp model.Correction

	if id == "" {
		return nil, errors.Newf("id is empty")
	}

	tx := r.Pool.DB.WithContext(ctx).
		Where(&model.Correction{ID: id}).
		Take(&resp)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, ErrObjectNotFound
		}

		return nil, errors.Wrap(tx.Error, "executing select correction record by id query")
	}

	return &resp, nil
}
