package errors

import (
	"fmt"
)

// ErrUnknownQueueType represents unsupported type error.
type ErrUnknownQueueType struct {
	t string
}

func (e *ErrUnknownQueueType) Error() string {
	return fmt.Sprintf(
		"Unknown queue type '%s'",
		e.t,
	)
}

//

// NewErrUnknownQueueType creates new ErrUnknownQueueType error.
func NewErrUnknownQueueType(t string) error {
	return &ErrUnknownQueueType{t}
}
