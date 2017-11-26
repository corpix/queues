package main

import (
	"bytes"
	"sync"
	"time"

	"github.com/corpix/formats"
	logger "github.com/corpix/loggers/logger/logrus"
	"github.com/sirupsen/logrus"

	"github.com/cryptounicorns/queues"
	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/message"
	"github.com/cryptounicorns/queues/producer"
	"github.com/cryptounicorns/queues/queue/channel"
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
			err error
		)

		err = consumer.PipeToWriterWith(
			mc,
			func(v interface{}) (message.Message, error) {
				var (
					buf []byte
					err error
				)

				buf, err = format.Marshal(v)
				if err != nil {
					return nil, err
				}

				return buf, nil
			},
			log,
		)
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		var (
			n   = 0
			err error
		)

		for {
			if n >= 5 {
				break
			}

			err = producer.PipeFromReaderWith(
				bytes.NewBuffer([]byte(`{"text":"hello"}`)),
				func(buf []byte) (interface{}, error) {
					var (
						m   = Message{}
						err error
					)

					err = format.Unmarshal(buf, &m)
					if err != nil {
						return nil, err
					}

					return m, nil
				},
				mp,
			)
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
