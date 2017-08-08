package nsq

import (
	"github.com/bitly/go-nsq"

	"github.com/corpix/queues/message"
)

type Handler func(m message.Message)

func (h Handler) HandleMessage(m *nsq.Message) error {
	h(message.Message(m.Body))

	return nil
}

func NewHandler(h Handler) nsq.Handler {
	return Handler(h)
}
