package storage

import "errors"

var (
	ErrEntityNotFound   = errors.New("Entity not found")
	ErrEntityExists     = errors.New("Entity already exists")
	ErrInvalidValueType = errors.New("Invalid value data type")
	ErrForeignKey       = errors.New("Foreign key constraint failed")
)
