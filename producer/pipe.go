package producer

import (
	"io"
	"io/ioutil"

	"github.com/cryptounicorns/queues/message"
)

type PrepareForProducerFn = func(buf []byte) (interface{}, error)

func PipeFromReader(r io.Reader, p Producer) error {
	var (
		buf message.Message
		err error
	)

	buf, err = ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return p.Produce(buf)
}

func PipeFromReaderWith(r io.Reader, fn PrepareForProducerFn, p Generic) error {
	var (
		m   interface{}
		buf []byte
		err error
	)

	buf, err = ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	m, err = fn(buf)
	if err != nil {
		return err
	}

	return p.Produce(m)
}
