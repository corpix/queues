queues
-------

[![Build Status](https://travis-ci.org/corpix/queues.svg?branch=master)](https://travis-ci.org/corpix/queues)

Universal API on top of various message queue systems such as:

- nsq
- kafka

## Usage example

``` console
$ go run ./example/nsq/nsq.go
INFO[0000] INF    2 [ticker/example_consumer] (127.0.0.1:4150) connecting to nsqd
INFO[0000] Producing: hello
INFO[0000] INF    1 (127.0.0.1:4150) connecting to nsqd
INFO[0000] Consumed: hello
INFO[0005] Producing: hello
INFO[0005] Consumed: hello
INFO[0010] Producing: hello
INFO[0010] Consumed: hello
```
