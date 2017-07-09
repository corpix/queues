package queues

import (
	"fmt"
)

type ErrUnknownQueueType struct {
	t string
}

func (e *ErrUnknownQueueType) Error() string {
	return fmt.Sprintf(
		"Unknown queue type '%s'",
		e.t,
	)
}
func NewErrUnknownQueueType(t string) error {
	return &ErrUnknownQueueType{t}
}

//
