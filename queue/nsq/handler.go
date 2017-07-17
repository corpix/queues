package nsq

import (
	"github.com/bitly/go-nsq"

	"github.com/corpix/queues/handler"
	"github.com/corpix/queues/message"
)

//

type Handler handler.Handler

func (h Handler) HandleMessage(m *nsq.Message) error {
	h(message.Message(m.Body))

	return nil
}

//

func NewHandler(h handler.Handler) nsq.Handler {
	return Handler(h)
}
