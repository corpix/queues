package nsq

import (
	"encoding/hex"
	"math/rand"
	"time"

	nsq "github.com/bitly/go-nsq"
	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/producer"
)

const (
	Name = "nsq"
)

var (
	random *rand.Rand
)

func init() {
	random = rand.New(
		rand.NewSource(
			time.Now().UnixNano(),
		),
	)
}

type Nsq struct {
	config Config
	log    loggers.Logger
}

func (q *Nsq) Producer() (producer.Producer, error) {
	return NewProducer(q.config, q.log)
}

func (q *Nsq) Consumer() (consumer.Consumer, error) {
	return NewConsumer(q.config, q.log)
}

func (q *Nsq) Close() error { return nil }

func New(c Config, l loggers.Logger) *Nsq {
	if c.Nsq == nil {
		c.Nsq = nsq.NewConfig()
	}
	if c.Channel == "" {
		var (
			buf = make([]byte, 4)
		)

		// XXX: We know that this sort of random source
		// does return nil error every time.
		random.Read(buf)
		c.Channel = "queue-" + hex.EncodeToString(buf)
	}

	return &Nsq{
		config: c,
		log:    l,
	}
}
