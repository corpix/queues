queues
-------

[![Build Status](https://travis-ci.org/corpix/queues.svg?branch=master)](https://travis-ci.org/corpix/queues)

Universal API on top of various message queue systems such as:

- nsq
- kafka

## NSQ example

Prepare NSQ:

> All commands should be run in separate terminal windows.

``` console
$ sudo rkt run --interactive corpix.github.io/nsq:1.0.0 --net=host -- nsqd --broadcast-address=127.0.0.1 --lookupd-tcp-address=127.0.0.1:4160 --tcp-address=127.0.0.1:4150
$ sudo rkt run --interactive corpix.github.io/nsq:1.0.0 --net=host -- nsqlookupd --tcp-address=127.0.0.1:4160

# Optionally you could run nsqadmin which will provide you a WEBUI for nsq topics etc.
$ sudo rkt run --interactive corpix.github.io/nsq:1.0.0 --net=host -- nsqadmin --lookupd-http-address=127.0.0.1:4161 --http-address=127.0.0.1:4171
```

Now you could run an NSQ example:

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
