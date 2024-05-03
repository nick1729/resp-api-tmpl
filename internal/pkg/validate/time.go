package validate

import (
	"time"

	"github.com/nick1729/resp-api-tmpl/internal/pkg/errors"
)

// RequiredTime - returns an error if the time is not set (January 1, year 1, 00:00:00 UTC).
func RequiredTime(v time.Time) error {
	if v.IsZero() {
		return errors.Newf("is required")
	}

	return nil
}
