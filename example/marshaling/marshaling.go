package main

import (
	"sync"
	"time"

	"github.com/corpix/formats"
	logger "github.com/corpix/loggers/logger/logrus"
	"github.com/sirupsen/logrus"

	"github.com/cryptounicorns/queues"
	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/producer"
	"github.com/cryptounicorns/queues/queue/channel"
	"github.com/cryptounicorns/queues/result"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	log := logger.New(logrus.New())

	format, err := formats.New(formats.JSON)
	if err != nil {
		log.Fatal(err)
	}

	q, err := queues.New(
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
	defer q.Close()

	c, err := q.Consumer()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	mc := consumer.NewUnmarshal(c, Message{}, format)
	defer mc.Close()

	p, err := q.Producer()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	mp := producer.NewMarshal(p, format)

	go func() {
		var (
			stream <-chan result.Generic
			err    error
		)

		stream, err = mc.Consume()
		if err != nil {
			panic(err)
		}

		for r := range stream {
			switch {
			case r.Err != nil:
				panic(r.Err)
			default:
				log.Printf("Consumed: %+v of type %T", r.Value, r.Value)
			}
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		n := 0
		message := Message{"hello"}

		for {
			if n >= 5 {
				break
			}

			log.Printf("Producing: %+v of type %T", message, message)

			err := mp.Produce(message)
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
