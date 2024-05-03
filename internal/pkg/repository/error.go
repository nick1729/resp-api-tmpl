package repository

import "errors"

var (
	ErrObjectNotFound  = errors.New("object not found")
	ErrDuplicateRowKey = errors.New("duplicate row key")
)
