package kafka

import (
	"github.com/Shopify/sarama"
)

type Config struct {
	Addrs              []string
	Topic              string
	ConsumerBufferSize uint
	Kafka              *sarama.Config
}
