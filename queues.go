package queues

import (
	"strings"

	"github.com/corpix/logger"
	"github.com/fatih/structs"

	"github.com/corpix/queues/handler"
	"github.com/corpix/queues/message"
	// "github.com/corpix/queues/queue/kafka"
	"github.com/corpix/queues/queue/nsq"
)

const (
	// KafkaQueueType = "kafka"
	NsqQueueType = "nsq"
)

type Config struct {
	Type string
	// Kafka kafka.Config
	Nsq nsq.Config
}

type Queue interface {
	Produce(message message.Message) error
	AddHandler(handler.Handler)
	Close() error
}

func NewFromConfig(l logger.Logger, c Config) (Queue, error) {
	for _, v := range structs.New(c).Fields() {
		if !strings.EqualFold(v.Name(), c.Type) {
			continue
		}

		switch {
		// case strings.EqualFold(c.Type, KafkaQueueType):
		// 	return kafka.NewFromConfig(
		// 		l,
		// 		v.Value().(kafka.Config),
		// 	)
		case strings.EqualFold(c.Type, NsqQueueType):
			return nsq.NewFromConfig(
				l,
				v.Value().(nsq.Config),
			)
		}
	}

	return nil, NewErrUnknownQueueType(c.Type)
}
