package service

import "errors"

var (
	ErrNotFound     = errors.New("Entity not found")
	ErrEntityExists = errors.New("Entity already exists")
)
