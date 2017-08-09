package errors

import (
	"fmt"
)

// ErrNilArgument represents a nil argument error.
type ErrNilArgument struct {
	Value interface{}
}

func (e *ErrNilArgument) Error() string {
	return fmt.Sprintf(
		"Received nil argument of type '%T'",
		e.Value,
	)
}

// NewErrNilArgument creates new ErrNilArgument error.
func NewErrNilArgument(value interface{}) error {
	return &ErrNilArgument{value}
}
