package models

import "errors"

var (
	ErrAlreadyExists  = errors.New("already exists")
	ErrDoesNotExist   = errors.New("does not exist")
	ErrNotImplemented = errors.New("not implemented")
)
