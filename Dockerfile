# build app
FROM golang:1.22-alpine AS builder

ENV GOPROXY https://goproxy.cn,direct
ENV CGO_ENABLED 0

WORKDIR /build
COPY . .

RUN go mod download
RUN go build -ldflags="-s -w" -o app

# build image
FROM alpine

WORKDIR /srv
COPY --from=builder /build/app /srv

ENTRYPOINT ["/srv/app"]