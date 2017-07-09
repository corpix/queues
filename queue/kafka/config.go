package kafka

import (
	"github.com/Shopify/sarama"
)

type Config struct {
	Addrs  []string
	Topic  string
	Sarama *sarama.Config
}
