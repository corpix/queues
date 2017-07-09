package main

import (
	"time"

	log "github.com/corpix/logger/logrus"
	"github.com/sirupsen/logrus"

	"github.com/corpix/queues"
	"github.com/corpix/queues/message"
	"github.com/corpix/queues/queue/kafka"
)

func main() {
	logger := log.New(logrus.New())

	c, err := queues.NewFromConfig(
		logger,
		queues.Config{
			Type: queues.KafkaQueueType,
			Kafka: kafka.Config{
				Addrs: []string{"127.0.0.1:9092"},
				Topic: "ticker",
			},
		},
	)
	if err != nil {
		logger.Fatal(err)
	}

	err = c.Consume(
		func(m message.Message) {
			logger.Printf("Consumed: %s", m)
		},
	)
	if err != nil {
		logger.Fatal(err)
	}

	go func() {
		data := []byte("hello")
		for {
			logger.Printf("Producing: %s", data)
			c.Produce(data)
			time.Sleep(5 * time.Second)
		}
	}()

	select {}
}
