package validate

import "github.com/nick1729/resp-api-tmpl/internal/pkg/errors"

type integers interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// IsEmptyInts - returns true if all values ​​are equal 0.
func IsEmptyInts[T integers](in ...T) bool {
	for _, v := range in {
		if v != 0 {
			return false
		}
	}

	return true
}

// BetweenOrEqual - returns an error if the value is less than min or greater than max.
func BetweenOrEqual[T integers](in, min, max T) error {
	if in < min || in > max {
		return errors.Newf("the %d value must be between or equal %d and %d", in, min, max)
	}

	return nil
}

// LessOrEqualInt - returns an error if the value is greater than max.
func LessOrEqualInt[T integers](v, max T) error {
	if v > max {
		return errors.Newf("the %d value must be less or equal %d", v, max)
	}

	return nil
}
