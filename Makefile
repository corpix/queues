.DEFAULT_GOAL = all

numcpus  := $(shell cat /proc/cpuinfo | grep '^processor\s*:' | wc -l)
version  := $(shell git rev-list --count HEAD).$(shell git rev-parse --short HEAD)

name     := queues
package  := github.com/cryptounicorns/$(name)

.PHONY: all
all:: dependencies

.PHONY: dependencies
dependencies::
	glide install

.PHONY: test
test:: dependencies
	go test -v $(shell glide novendor)

.PHONY: bench
bench:: dependencies
	go test -bench=. -v $(shell glide novendor)

.PHONY: lint
lint:: dependencies
	go vet -v $(shell glide novendor)

.PHONY: check
check:: lint test

.PHONY: clean
clean::
	git clean -xddff
