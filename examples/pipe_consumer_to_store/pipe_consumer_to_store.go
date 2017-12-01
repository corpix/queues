package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"sync"

	"github.com/corpix/formats"
	logger "github.com/corpix/loggers/logger/logrus"
	"github.com/corpix/stores"
	"github.com/corpix/stores/store/memory"
	"github.com/davecgh/go-spew/spew"
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
		Type: nsq.Name,
		Nsq: nsq.Config{
			Addr:  "127.0.0.1:4150",
			Topic: "pipe",
		},
	}
)

func main() {
	var (
		log    = logger.New(logrus.New())
		wg     = &sync.WaitGroup{}
		ctx    context.Context
		cancel context.CancelFunc
		f      formats.Format
		s      stores.Store
		q      queues.Queue
		p      producer.Producer
		sm     map[string]interface{}
		err    error
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	f, err = formats.New(format)
	if err != nil {
		log.Fatal(err)
	}

	s, err = stores.New(
		stores.Config{
			Type:   memory.Name,
			Memory: memory.Config{},
		},
		log,
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		var (
			err error
		)
		err = queues.PipeConsumerToStoreWith(
			queues.GenericConfig{
				Format: format,
				Queue:  queue,
			},
			ctx,
			func(v interface{}) (string, interface{}, error) {
				// XXX: It is not adviced to do side-effects here
				// but we need this wg.Done() to show you a store contents :)
				defer wg.Done()
				var (
					hash = sha1.New()
					buf  []byte
					err  error
				)

				buf, err = f.Marshal(v)
				if err != nil {
					return "", nil, err
				}

				_, err = hash.Write(buf)
				if err != nil {
					return "", nil, err
				}

				return hex.EncodeToString(hash.Sum(nil)), v, nil
			},
			s,
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

	sm, err = s.Map()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Store contents: %s", spew.Sdump(sm))
}
