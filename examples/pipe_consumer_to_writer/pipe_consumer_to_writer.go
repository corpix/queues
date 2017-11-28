package main

import (
	"bytes"
	"sync"

	"github.com/corpix/formats"
	logger "github.com/corpix/loggers/logger/logrus"
	"github.com/sirupsen/logrus"

	"github.com/cryptounicorns/queues"
	"github.com/cryptounicorns/queues/producer"
	"github.com/cryptounicorns/queues/queue/nsq"
)

const (
	format = formats.JSON
)

var (
	queue = queues.Config{
		Type: queues.NsqQueueType,
		Nsq: nsq.Config{
			Addr:  "127.0.0.1:4150",
			Topic: "pipe",
		},
	}
)

func main() {
	var (
		log = logger.New(logrus.New())
		wg  = &sync.WaitGroup{}
		w   = bytes.NewBuffer(nil)
		f   formats.Format
		q   queues.Queue
		p   producer.Producer
		err error
	)

	f, err = formats.New(format)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		var (
			err error
		)
		err = queues.PipeConsumerToWriterWith(
			queues.GenericConfig{
				Format: format,
				Queue:  queue,
			},
			func(v interface{}) ([]byte, error) {
				var (
					buf []byte
					err error
				)
				// XXX: It is not adviced to do side-effects here
				// but we need this wg.Done() to show you a buffer contents :)
				defer wg.Done()

				buf, err = f.Marshal(v)
				if err != nil {
					return nil, err
				}

				return buf, nil
			},
			w,
			log,
		)
		if err != nil {
			log.Fatal(err)
		}
	}()

	q, err = queues.New(queue, log)
	if err != nil {
		log.Fatal(err)
	}
	defer q.Close()

	p, err = q.Producer()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	wg.Add(1)
	err = p.Produce([]byte(`{"text": "hello"}`))
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	log.Printf("Buffer contents: %s", w.Bytes())
}
