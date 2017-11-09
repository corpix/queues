package result

import (
	"github.com/corpix/queues/message"
)

type Result struct {
	// Hey, go authors, you forgot to add this thing...
	// Called "generics" :D
	Value message.Message
	Err   error
}
