package result

import (
	"github.com/cryptounicorns/queues/message"
)

type Result struct {
	// Hey, go authors, you forgot to add this thing...
	// Called "generics" :D
	Value message.Message
	Err   error
}

type Generic struct {
	// This drives me to the point of a murder.
	// Hey, gophers, you still don't need generics?
	// Eat this.
	Value interface{}
	Err   error
}
