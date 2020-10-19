FROM golang:1.13.3 AS base
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn/
WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
ENV TZ "Asia/Shanghai"
COPY . $WORKDIR
RUN CGO_ENABLED=0 GOOS=linux go build -o app  main.go



