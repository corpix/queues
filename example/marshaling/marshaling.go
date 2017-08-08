package main

import (
	"sync"
	"time"

	"github.com/corpix/formats"
	logger "github.com/corpix/logger/target/logrus"
	"github.com/sirupsen/logrus"

	"github.com/corpix/queues"
	"github.com/corpix/queues/consumer"
	"github.com/corpix/queues/producer"
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

	q, err := queues.NewFromConfig(
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

	c, err := q.Consumer()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	uc, err := consumer.NewUnmarshalConsumer(
		Message{},
		c,
		json,
		log,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer uc.Close()

	p, err := q.Producer()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	mp, err := producer.NewMarshalProducer(
		p,
		json,
		log,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer mp.Close()

	go func() {
		for m := range uc.Consume() {
			log.Printf("Consumed: %#v", m)
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		n := 0
		message := Message{"foo", "bar"}

		for {
			if n >= 5 {
				break
			}

			log.Printf("Producing: %#v", message)

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
