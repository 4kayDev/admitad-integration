package storage

import "errors"

var (
	ErrEntityNotFound = errors.New("Entity not found")
	ErrEntityExists   = errors.New("Entity already exists")
	ErrInvalidValue   = errors.New("Invalid value")
	ErrForeignKey     = errors.New("Foreign key constraint failed")
)
