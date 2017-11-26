package consumer

import (
	"reflect"

	"github.com/corpix/formats"

	"github.com/cryptounicorns/queues/result"
)

type Unmarshal struct {
	consumer    Consumer
	reflectType reflect.Type
	format      formats.Format
	done        chan struct{}
}

func (c *Unmarshal) pump(consumerStream <-chan result.Result, stream chan result.Generic) {
	var (
		r result.Generic
	)

	for cr := range consumerStream {
		select {
		case <-c.done:
			return
		default:
			if cr.Err == nil {
				r.Value = reflect.New(c.reflectType).Interface()
				r.Err = c.format.Unmarshal(
					cr.Value,
					r.Value,
				)
			} else {
				r.Err = cr.Err
			}

			stream <- r
		}
	}
}

func (c *Unmarshal) Consume() (<-chan result.Generic, error) {
	var (
		stream         = make(chan result.Generic)
		consumerStream <-chan result.Result
		err            error
	)

	consumerStream, err = c.consumer.Consume()
	if err != nil {
		return nil, err
	}

	go c.pump(consumerStream, stream)

	return stream, nil
}

func (c *Unmarshal) Close() error {
	close(c.done)
	return nil
}

func NewUnmarshal(cr Consumer, t interface{}, f formats.Format) *Unmarshal {
	return &Unmarshal{
		consumer:    cr,
		reflectType: reflect.TypeOf(t),
		format:      f,
		done:        make(chan struct{}),
	}
}
