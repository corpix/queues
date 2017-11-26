package result

import (
	"github.com/cryptounicorns/queues/message"
)

type Result struct {
	Value message.Message
	Err   error
}

type Generic struct {
	Value interface{}
	Err   error
}
