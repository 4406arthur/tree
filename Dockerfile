# build stage
FROM golang:1.14.3-alpine AS build-env
ADD . /src
RUN \
    cd /src && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -o /mlass_pod_watcher

# final stage
FROM alpine:3.7
COPY --from=build-env /mlass_pod_watcher /usr/local/bin

RUN \
    apk update \
    && apk add ca-certificates \
    && apk add librdkafka-dev pkgconf \
    && rm -rf /var/cache/apk/*

# EXPOSE 80 443

ENTRYPOINT ["mlass_pod_watcher"]
