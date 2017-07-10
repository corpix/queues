package main

import (
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"

	"github.com/corpix/logger/encoder"
	logrusLogger "github.com/corpix/logger/target/logrus"
	nsqLogger "github.com/corpix/logger/target/nsq"
)

func main() {
	p, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	c := nsqLogger.Config{
		Level: nsqLogger.InfoLevel,
		Topic: "log",
	}

	l := nsqLogger.New(
		c,
		p,
		logrusLogger.New(logrus.New()),
		encoder.NewJSON(),
	)

	l.Debug("Hidden")
	l.Print("Info")
	l.Error("Error")
	l.Fatal("Fatal")
}
