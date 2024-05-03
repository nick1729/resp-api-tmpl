package validate

import (
	"unicode/utf8"

	"github.com/nick1729/resp-api-tmpl/internal/pkg/errors"
)

// isEmptyStrings - returns true if all lines are empty.
func isEmptyStrings(list ...string) bool {
	for _, v := range list {
		if v != "" {
			return false
		}
	}

	return true
}

// isNotEmptyStrings - returns true if all strings are non-empty.
func isNotEmptyStrings(list ...string) bool {
	for _, v := range list {
		if v == "" {
			return false
		}
	}

	return true
}

// IsEmptyStrings - returns true if all lines are empty.
func IsEmptyStrings(list ...string) bool {
	return isEmptyStrings(list...)
}

// IsNotEmptyStrings - returns true if all strings are non-empty.
func IsNotEmptyStrings(list ...string) bool {
	return isNotEmptyStrings(list...)
}

// RequiredStrings - returns an error if all lines are empty.
func RequiredStrings(list ...string) error {
	if isEmptyStrings(list...) {
		return errors.Newf("is required")
	}

	return nil
}

// LessOrEqualStringLen - returns an error if the string length is greater than max.
func LessOrEqualStringLen(v string, max int) error {
	runeLen := utf8.RuneCountInString(v)
	if runeLen > max {
		return errors.Newf("the %s value is of length %d must be less or equal %d", v, runeLen, max)
	}

	return nil
}
