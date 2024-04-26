package errors

import (
	"errors"
	"fmt"
)

func New(message string) error {
	return errors.New(message)
}

func Newf(message string, a ...any) error {
	return fmt.Errorf(message, a...)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Wrap(err error, message string) error {
	return fmt.Errorf("%s: %w", message, err)
}

func Wrapf(err error, message string, a ...any) error {
	return fmt.Errorf("%s: %w", fmt.Sprintf(message, a...), err)
}

func Unwrap(err error) error {
	for errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}

	return err
}
