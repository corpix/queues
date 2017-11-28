package main

import (
	"bytes"
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
	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/producer"
	"github.com/cryptounicorns/queues/queue/channel"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	var (
		log = logger.New(logrus.New())
		wg  = &sync.WaitGroup{}
		f   formats.Format
		s   stores.Store
		q   queues.Queue
		c   consumer.Consumer
		mc  consumer.Generic
		p   producer.Producer
		mp  producer.Generic
		sm  map[string]interface{}
		err error
	)

	f, err = formats.New(formats.JSON)
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
	defer s.Close()

	q, err = queues.New(
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

	c, err = q.Consumer()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	mc = consumer.NewUnmarshal(c, Message{}, f)
	defer mc.Close()

	p, err = q.Producer()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	mp = producer.NewMarshal(p, f)
	defer mp.Close()

	go func() {
		var (
			err error
		)

		err = consumer.PipeToStoreWith(
			mc,
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
		)
		if err != nil {
			log.Fatal(err)
		}
	}()

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
			wg.Add(1)

			err = producer.PipeFromReaderWith(
				bytes.NewBuffer([]byte(`{"text":"hello"}`)),
				func(buf []byte) (interface{}, error) {
					var (
						m   = Message{}
						err error
					)

					err = f.Unmarshal(buf, &m)
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

			n++
		}
	}()

	wg.Wait()
	log.Print("Done")

	sm, err = s.Map()
	if err != nil {
		log.Fatal(err)
	}

	// Here will be only 1 record because
	// we adding data with key sha1(data)
	// and data always same.
	log.Printf("Store contents: %s", spew.Sdump(sm))
}
