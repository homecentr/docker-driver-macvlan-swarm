FROM golang:1.10-alpine as builder

COPY . /go/src/github.com/homecentr/docker-driver-macvlan-swarm

WORKDIR /go/src/github.com/homecentr/docker-driver-macvlan-swarm

RUN set -ex && \
    apk add --no-cache --virtual .build-deps gcc libc-dev git && \
    go get ./... && \
    go install --ldflags '-extldflags "-static"' && \
    apk del .build-deps

CMD ["/go/bin/docker-driver-macvlan-swarm"]

FROM alpine:3.11.3
RUN apk update && \
    mkdir -p /run/docker/plugins

COPY --from=builder /go/bin/docker-driver-macvlan-swarm .

CMD ["/docker-driver-macvlan-swarm"]
