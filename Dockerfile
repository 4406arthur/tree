# build stage
FROM golang:1.15.2-alpine AS build-env
WORKDIR /src
ADD . /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -o /mlaas_pod_watcher

# final stage
FROM alpine:3.12.1
COPY --from=build-env /mlaas_pod_watcher /usr/local/bin

RUN \
    apk update \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/*

# EXPOSE 80 443

ENTRYPOINT ["mlaas_pod_watcher"]
