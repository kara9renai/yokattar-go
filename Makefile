BINARY := yokattar-go
MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

PATH := $(PATH):${MAKEFILE_DIR}bin
SHELL := env PATH="$(PATH)" /bin/bash
# for go
export CGO_ENABLED = 0
GOARCH = amd64

COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GIT_URL=local-git://

LDFLAGS := -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

build: build-linux

build-default:
	go build ${LDFLAGS} -o build/${BINARY}

build-linux:
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o build/${BINARY}-linux-${GOARCH} .

prepare: mod

mod:
	go mod download

test:
	go test $(shell go list ${MAKEFILE_DIR}/...)

lint:
	if ! [ -x $(GOPATH)/bin/golangci-lint ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.38.0 ; \
	fi
	golangci-lint run --concurrency 2

vet:
	go vet ./...

clean:
	git clean -f -X app bin build

.PHONY:	test clean

# for docker compose

stop:
	docker compose stop

up:
	docker compose up -d

reset-mysql:
	rm -rfd .data/mysql && make up

mysql:
	docker compose exec mysql bin/bash -c 'mysql -u$$MYSQL_USER -p$$MYSQL_PASSWORD'

NAME = Taro
DNAME = 太郎

create-account:
	curl -X 'POST' \
	'http://localhost:8080/v1/accounts' \
	-H 'accept: application/json' \
	-H 'Content-Type: application/json' \
	-d '{ \
	"username": "${NAME}", \
	"password": "P@ssw0rd", \
	"displayName": "${DNAME}", \
	"avatar": "avatar_pictrue.jpg", \
	"header": "header_picture.jpg", \
	"note": "University Student" \
	}'

FNAME = Saki

follow:
	curl -X 'POST' \
	'http://localhost:8080/v1/accounts/${FNAME}/follow' \
	-H 'accept: application/json' \
	-H 'Authentication: username ${NAME}' \
	-d ''

like:
	curl -X 'POST' \
	'http://localhost:8080/v1/like' \
	-H 'accept: application/json' \
	-H 'Authentication: username ${NAME}' \
	-d '{ \
	"like_id" : 1 \
	}'
