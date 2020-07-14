# build stage
FROM golang:1.14.4 AS build-env
ADD . /src
RUN \
    apt-get update \
    && apt-get install -y librdkafka-dev \
    && cd /src \
    # && go build -mod vendor -o /mlass_pod_watcher
    && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -mod vendor -o /mlass_pod_watcher

# final stage
FROM alpine:3.7
COPY --from=build-env /mlass_pod_watcher /usr/local/bin

RUN \
    apk update \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/*

# EXPOSE 80 443

ENTRYPOINT ["mlass_pod_watcher"]
