package main

import (
	"io"
	"sync"
	"time"

	logger "github.com/corpix/loggers/logger/logrus"
	"github.com/sirupsen/logrus"

	"github.com/cryptounicorns/queues/message"
	"github.com/cryptounicorns/queues/queue/readwriter"
	"github.com/cryptounicorns/queues/result"
)

type Proxy struct {
	stream chan []byte
}

func (p *Proxy) Write(buf []byte) (int, error) {
	p.stream <- buf
	return len(buf), nil
}

func (p *Proxy) Read(buf []byte) (int, error) {
	return copy(
		buf,
		<-p.stream,
	), io.EOF
}

func main() {
	var (
		log = logger.New(logrus.New())
		err error
	)

	q := readwriter.New(
		&Proxy{stream: make(chan []byte)},
		readwriter.Config{},
		log,
	)
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
		var (
			stream <-chan result.Result
			err    error
		)

		stream, err = c.Consume()
		if err != nil {
			panic(err)
		}

		for r := range stream {
			switch {
			case r.Err != nil:
				panic(r.Err)
			default:
				log.Printf("Consumed: %s", r.Value)
			}
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
