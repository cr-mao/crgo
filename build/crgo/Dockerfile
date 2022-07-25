FROM golang:1.15.4-alpine3.12 AS base
ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on GOPROXY=https://goproxy.cn/ TZ="Asia/Shanghai"
WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY . $WORKDIR
ARG VERSION
ARG BUILDTIME
RUN  go build -ldflags="-s -w -X infra.Version=${VERSION} -X infra.BuildTime=${BUILDTIME}" -o /usr/local/bin/app  main.go && go clean -cache

FROM alpine:3.12
ENV TZ "Asia/Shanghai"
COPY --from=base /usr/local/bin/app /usr/local/bin/