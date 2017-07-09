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

## Kafka example

Prepare Kafka:

> All commands should be run in separate terminal windows.

``` console
$ sudo rkt run --interactive coreos.com/etcd:v3.1.8 --net=host -- --log-output=stderr --debug
$ sudo rkt run --interactive corpix.github.io/zetcd:0.0.2 --net=host -- --zkaddr=127.0.0.1:2181 --endpoints=127.0.0.1:2379 --logtostderr

# Init zetcd with data(or kafka will fail to start)
$ sudo rkt run                                 \
    --interactive corpix.github.io/zetcd:0.0.2 \
    --net=host --exec=/bin/bash                \
    -- -c "
        zkctl create '/' ''
        zkctl create '/brokers' ''
        zkctl create '/brokers/ids' ''
        zkctl create '/brokers/topics' ''
    "

$ sudo rkt run --interactive corpix.github.io/kafka:2.12-0.10.2.1-1496226351 --net=host
```

Now you could run a Kafka example:

``` console
$ go run ./example/kafka/kafka.go
INFO[0001] Producing: hello
INFO[0001] Consumed: hello
INFO[0006] Producing: hello
INFO[0006] Consumed: hello
```
