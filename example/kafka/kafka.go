package main

import (
	"time"

	logger "github.com/corpix/logger/target/logrus"
	"github.com/sirupsen/logrus"

	"github.com/corpix/queues"
	"github.com/corpix/queues/message"
	"github.com/corpix/queues/queue/kafka"
)

func main() {
	log := logger.New(logrus.New())

	c, err := queues.NewFromConfig(
		queues.Config{
			Type: queues.KafkaQueueType,
			Kafka: kafka.Config{
				Addrs: []string{"127.0.0.1:9092"},
				Topic: "ticker",
			},
		},
		log,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Consume(
		func(m message.Message) {
			log.Printf("Consumed: %s", m)
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		data := []byte("hello")
		for {
			log.Printf("Producing: %s", data)
			err := c.Produce(data)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	select {}
}
