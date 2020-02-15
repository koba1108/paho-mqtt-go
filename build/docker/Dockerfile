# バイナリビルド
FROM golang:1.13.4 AS build-env

ENV GO111MODULE=on

ADD . /build
WORKDIR /build

RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static" "-w" "-s"' -o main cmd/subscriber/main.go

# バイナリ圧縮
FROM alpine:3.10.3 AS upx
RUN apk update && apk add --no-cache upx
WORKDIR /build
COPY --from=build-env /build/main /build/main
RUN upx --best /build/main

# 実行イメージ
FROM alpine:3.10.3
RUN apk update && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=upx /build/main main
COPY --from=build-env /build/secrets /app/secrets

EXPOSE 8080
ENTRYPOINT ./main
