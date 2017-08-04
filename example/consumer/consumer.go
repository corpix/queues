package main

import (
	"time"

	"github.com/corpix/formats"
	logger "github.com/corpix/logger/target/logrus"
	"github.com/sirupsen/logrus"

	"github.com/corpix/queues"
	"github.com/corpix/queues/consumer"
	"github.com/corpix/queues/queue/channel"
)

type Message struct {
	Foo string
	Bar string
}

func main() {
	var (
		log = logger.New(logrus.New())
		err error
	)

	json, err := formats.New(formats.JSON)
	if err != nil {
		log.Fatal(err)
	}

	c, err := queues.NewFromConfig(
		queues.Config{
			Type: queues.ChannelQueueType,
			Channel: channel.Config{
				Capacity: 128,
			},
		},
		log,
	)
	if err != nil {
		log.Fatal(err)
	}

	consumer, err := consumer.New(
		Message{},
		json,
		log,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Consume(consumer.Handler)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for m := range consumer.GetFeed() {
			log.Printf("Consumed: %#v", m)
		}
	}()

	go func() {
		message := Message{"foo", "bar"}
		data, err := json.Marshal(message)
		if err != nil {
			log.Fatal(err)
		}

		for {
			log.Printf("Producing: %s", data)

			err = c.Produce(data)
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(5 * time.Second)
		}
	}()

	select {}
}
