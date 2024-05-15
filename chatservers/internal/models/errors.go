package models

import "errors"

var (
	ErrAlreadyExists  = errors.New("already exists")
	ErrNotImplemented = errors.New("not implemented")
)
