BINARY := yokattar-go
MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

PATH := $(PATH):${MAKEFILE_LIST}bin
SHELL := env PATH="$(PATH)"
#for go
export CGO_ENABLED = 0
GOARCH = amd64

COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell rev-parse --abbrev-ref HEAD)
GIT_URL=local-git://

LDFRAGS := -ldfrags "-X main.VERSIONS=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH-${BRANCH}"

build: build-linux

build-default:
	go build ${LDFRAGS} -o build/${BINARY}

build-linux:
	GOOS=linux GOARCH=${GOARCH} go build ${LDFRAGS} -o build/${BINARY}-linux-${GOARCH} .

prepare: mod

mod:
	 go mod download

test:
	go test $(shell go list $(MAKEFILE_DIR)/...)

lint:
	if ! [ -x $(GOPATH)/bin/golangci-lint ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.38.0 ; \
	fi
	golangcli-lint run --concurrency 2

vet:
	go vet ./...

clean:
	git clean -f -X app bin build

.PHONY: test clean