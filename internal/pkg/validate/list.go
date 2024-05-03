package validate

import "github.com/nick1729/resp-api-tmpl/internal/pkg/errors"

// IsEmptyList - returns true if the collection is empty.
func IsEmptyList[T comparable](v []T) bool {
	return len(v) == 0
}

// RequiredList - returns an error if the collection is empty.
func RequiredList[T comparable](list []T) error {
	if len(list) == 0 {
		return errors.Newf("is required")
	}

	return nil
}

// LessOrEqualListLen - returns an error if the collection size is larger than max.
func LessOrEqualListLen[T comparable](v []T, max int) error {
	err := LessOrEqualInt(len(v), max)
	if err != nil {
		return errors.Wrap(err, "invalid list length")
	}

	return nil
}
