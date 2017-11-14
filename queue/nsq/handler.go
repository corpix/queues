package nsq

import (
	"github.com/bitly/go-nsq"

	"github.com/cryptounicorns/queues/message"
)

type Handler func(m message.Message)

func (h Handler) HandleMessage(m *nsq.Message) error {
	h(m.Body)

	return nil
}

func NewHandler(h Handler) nsq.Handler {
	return Handler(h)
}
