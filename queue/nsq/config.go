package nsq

import (
	nsq "github.com/bitly/go-nsq"
)

type Config struct {
	Addr               string
	Topic              string
	Channel            string
	ConsumerBufferSize uint
	LogLevel           nsq.LogLevel
	Nsq                *nsq.Config
}
