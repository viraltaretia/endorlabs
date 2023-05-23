package errors

import "errors"

var (
	ErrObjectNotFound    = errors.New("object not found")
	ErrInvalidObjectKind = errors.New("invalid object kind")
	ErrInvalidObject     = errors.New("invalid object")
	ErrInternal          = errors.New("internal error")
)
