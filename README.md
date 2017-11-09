queues
-------

[![Build Status](https://travis-ci.org/corpix/queues.svg?branch=master)](https://travis-ci.org/corpix/queues)

Universal API on top of various message queue systems such as:

- channel
- nsq
- kafka

## Channel example

This package supports a basic wrapper around Go channel.

> It is intended to show the basic principles and for debugging.

``` console
$ go run ./example/channel/channel.go
INFO[0000] Producing: hello
INFO[0000] Consumed: hello
INFO[0002] Producing: hello
INFO[0002] Consumed: hello
INFO[0004] Producing: hello
INFO[0004] Consumed: hello
INFO[0006] Producing: hello
INFO[0006] Consumed: hello
INFO[0008] Producing: hello
INFO[0008] Consumed: hello
INFO[0010] Done
```

## NSQ example

Now you could run an NSQ example:

``` console
$ go run ./example/nsq/nsq.go
INFO[0000] INF    1 [nsq-example/queues-nsq-example] (127.0.0.1:4150) connecting to nsqd
INFO[0000] Producing: hello
INFO[0000] INF    2 (127.0.0.1:4150) connecting to nsqd
INFO[0000] Consumed: hello
INFO[0000] Consumed: hello
INFO[0002] Producing: hello
INFO[0002] Consumed: hello
INFO[0004] Producing: hello
INFO[0004] Consumed: hello
INFO[0006] Producing: hello
INFO[0006] Consumed: hello
INFO[0008] Producing: hello
INFO[0008] Consumed: hello
INFO[0010] Done
INFO[0010] INF    2 stopping
INFO[0010] INF    2 (127.0.0.1:4150) beginning close
INFO[0010] INF    1 [nsq-example/queues-nsq-example] stopping...
INFO[0010] INF    2 exiting router
INFO[0010] INF    2 (127.0.0.1:4150) readLoop exiting
INFO[0010] INF    2 (127.0.0.1:4150) breaking out of writeLoop
INFO[0010] INF    2 (127.0.0.1:4150) writeLoop exiting
INFO[0010] INF    1 [nsq-example/queues-nsq-example] (127.0.0.1:4150) received CLOSE_WAIT from nsqd
INFO[0010] INF    1 [nsq-example/queues-nsq-example] (127.0.0.1:4150) beginning close
INFO[0010] INF    1 [nsq-example/queues-nsq-example] (127.0.0.1:4150) readLoop exiting
INFO[0010] INF    1 [nsq-example/queues-nsq-example] (127.0.0.1:4150) breaking out of writeLoop
INFO[0010] INF    1 [nsq-example/queues-nsq-example] (127.0.0.1:4150) writeLoop exiting
INFO[0010] INF    2 (127.0.0.1:4150) finished draining, cleanup exiting
INFO[0010] INF    2 (127.0.0.1:4150) clean close complete
INFO[0010] INF    1 [nsq-example/queues-nsq-example] (127.0.0.1:4150) finished draining, cleanup exiting
INFO[0010] INF    1 [nsq-example/queues-nsq-example] (127.0.0.1:4150) clean close complete
INFO[0010] WRN    1 [nsq-example/queues-nsq-example] there are 0 connections left alive
INFO[0010] INF    1 [nsq-example/queues-nsq-example] stopping handlers
INFO[0010] INF    1 [nsq-example/queues-nsq-example] rdyLoop exiting
```

## Kafka example

Now you could run a Kafka example:

``` console
$ go run ./example/kafka/kafka.go
INFO[0001] Producing: hello
INFO[0001] Consumed: hello
INFO[0003] Producing: hello
INFO[0003] Consumed: hello
INFO[0005] Producing: hello
INFO[0005] Consumed: hello
INFO[0007] Producing: hello
INFO[0007] Consumed: hello
INFO[0009] Producing: hello
INFO[0009] Consumed: hello
INFO[0011] Done
```
