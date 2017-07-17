package main

import (
	"time"

	log "github.com/corpix/logger/target/logrus"
	"github.com/sirupsen/logrus"

	"github.com/corpix/queues"
	"github.com/corpix/queues/message"
	"github.com/corpix/queues/queue/nsq"
)

func main() {
	logrusLogger := logrus.New()
	log := log.New(logrusLogger)

	c, err := queues.NewFromConfig(
		queues.Config{
			Type: queues.NsqQueueType,
			Nsq: nsq.Config{
				Addr:    "127.0.0.1:4150",
				Topic:   "ticker",
				Channel: "example_consumer",
				LogLevel: nsq.NewLogLevelFromLogrus(
					logrusLogger.Level,
				),
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
				log.Error(err)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	select {}
}
