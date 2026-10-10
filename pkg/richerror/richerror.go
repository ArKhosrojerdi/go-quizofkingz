package richerror

import (
	"errors"
	"fmt"
)

type Kind int

const (
	KindInvalid Kind = iota + 1
	KindUnauthorized
	KindNotFound
	KindConflict
	KindUnexpected
)

type Op string

type RichError struct {
	operation    Op
	wrappedError error
	message      string
	kind         Kind
}

func New(op Op) RichError {
	return RichError{operation: op}
}

func (r RichError) WithErr(err error) RichError {
	r.wrappedError = err
	return r
}

func (r RichError) WithMessage(message string) RichError {
	r.message = message
	return r
}

func (r RichError) WithKind(kind Kind) RichError {
	r.kind = kind
	return r
}

func (r RichError) Error() string {
	if r.wrappedError == nil {
		return fmt.Sprintf("%s: %s", r.operation, r.message)
	}

	return fmt.Sprintf("%s: %s: %v", r.operation, r.message, r.wrappedError)
}

func (r RichError) Unwrap() error {
	return r.wrappedError
}

func (r RichError) Kind() Kind {
	return r.kind
}

func (r RichError) Message() string {
	return r.message
}

func KindOf(err error) Kind {
	var richErr RichError
	if errors.As(err, &richErr) && richErr.kind != 0 {
		return richErr.kind
	}

	return KindUnexpected
}

func MessageOf(err error) string {
	var richErr RichError
	if errors.As(err, &richErr) {
		return richErr.message
	}

	return ""
}
