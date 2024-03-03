package service

import "errors"

var (
	ErrNotFound     = errors.New("entity not found")
	ErrEntityExists = errors.New("entity already exists")
)
