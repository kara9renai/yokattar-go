# dev, builder
FROM golang:1.16 AS golang
WORKDIR /work/yokattar-go

# dev
FROM golang as dev
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# builder
FROM golang AS builder
COPY ./ ./
RUN make prepare build-linux

# release
FROM alpine AS app
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=builder /work/yokattar-go/build/yokattar-go-linux-amd64 /usr/local/bin/yokattar-go
EXPOSE 8080
ENTRYPOINT ["yokattar-go"]
