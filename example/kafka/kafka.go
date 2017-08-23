package main

import (
	"sync"
	"time"

	logger "github.com/corpix/loggers/logger/logrus"
	"github.com/sirupsen/logrus"

	"github.com/corpix/queues"
	"github.com/corpix/queues/message"
	"github.com/corpix/queues/queue/kafka"
)

func main() {
	originalLogger := logrus.New()
	log := logger.New(originalLogger)

	q, err := queues.NewFromConfig(
		queues.Config{
			Type: queues.KafkaQueueType,
			Kafka: kafka.Config{
				Addrs: []string{"127.0.0.1:9092"},
				Topic: "kafka-example",
			},
		},
		log,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer q.Close()

	c, err := q.Consumer()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	p, err := q.Producer()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	go func() {
		for m := range c.Consume() {
			log.Printf("Consumed: %s", m)
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		n := 0
		message := message.Message("hello")

		for {
			if n >= 5 {
				break
			}

			log.Printf("Producing: %s", message)

			err := p.Produce(message)
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(2 * time.Second)
			n++
		}
	}()

	wg.Wait()
	log.Print("Done")

}
