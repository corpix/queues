
.DEFAULT_GOAL = all

name     := queues
package  := github.com/corpix/$(name)
build    := ./build
numcpus  := $(shell cat /proc/cpuinfo | grep '^processor\s*:' | wc -l)
version  := $(shell git rev-list --count HEAD).$(shell git rev-parse --short HEAD)
build_id := 0x$(shell dd if=/dev/urandom bs=40 count=1 2> /dev/null | sha1sum | awk '{print $$1}')
ldflags  := -X $(package)/cli.version=$(version) \
            -B $(build_id)

.PHONY: all
all: $(name)

.PHONY: $(name)
$(name): dependencies
	mkdir -p $(build)
	@echo "Build id: $(build_id)"
	go build -a -ldflags "$(ldflags)" -v \
                 -o build/$(name)            \
                 $(package)/$(name)

.PHONY: build
build: $(name)

.PHONY: test
test: dependencies
	go test -v \
           $(shell glide novendor)

.PHONY: bench
bench: dependencies
	go test        \
           -bench=. -v \
           $(shell glide novendor)

.PHONY: lint
lint: dependencies
	go vet $(shell glide novendor)
	gometalinter                     \
		--deadline=5m            \
		--concurrency=$(numcpus) \
		$(shell glide novendor)

.PHONY: check
check: lint test

.PHONY: tools
tools:
	@if [ ! -e "$(GOPATH)"/bin/glide ]; then go get github.com/Masterminds/glide; fi
	@if [ ! -e "$(GOPATH)"/bin/godef ]; then go get github.com/rogpeppe/godef; fi
	@if [ ! -e "$(GOPATH)"/bin/gocode ]; then go get github.com/nsf/gocode; fi
	@if [ ! -e "$(GOPATH)"/bin/gometalinter ]; then go get github.com/alecthomas/gometalinter && gometalinter --install; fi
	@if [ ! -e "$(GOPATH)"/src/github.com/stretchr/testify/assert ]; then go get github.com/stretchr/testify/assert; fi


.PHONY: dependencies
dependencies: tools
	glide install

.PHONY: clean
clean: tools
	rm -rf $(build)
	glide cache-clear
