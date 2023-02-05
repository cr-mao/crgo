# my golang framework

- 项目框架的搭建

#### Usage

```shell 
$ export GO111MODULE=on
$ go mod download

#$ go get github.com/golang/protobuf/protoc-gen-go@v1.3.5
#$ brew install protobuf

## Or Download https://github.com/protocolbuffers/protobuf/releases/download/v3.9.2/protoc-3.9.2-osx-x86_64.zip
#$ unzip protoc-3.9.2-osx-x86_64.zip

# www.crblog.cc/posts/go/grpc-protobuf-action  参见这篇文章 安装新的proto相关工具

$ go run main.go --help 
$ make serve
```


#### consul 本地启动
```shell
# https://developer.hashicorp.com/consul/downloads 
# sudo mkdir -p /etc/consul.d
sudo consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul -node=n1 -bind=127.0.0.1 -ui -config-dir /etc/consul.d -rejoin -join 127.0.0.1 -client 0.0.0.0
```

#### prometheus 本地启动
```shell
docker pull prom/prometheus:v2.20.1
docker run -d --name prometheus -p 9090:9090 -v /Users/mac/code/crgo/build/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus:v2.20.1
```


#### jaeger 本地启动

https://www.jaegertracing.io/docs/1.41/getting-started/

```shell
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -e COLLECTOR_OTLP_ENABLED=true \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one
```



#### 构建镜像

```shell
## 注意项目里面的 一些服务地址 如mysql,redis 要修改成 容器ip ...不然跑不起来的。 
docker build -t crgo:v1  --build-arg VERSION="$(git rev-parse --short HEAD)",BUILDTIME="$(date +%FT%T)" -f  ./build/crgo/Dockerfile ./ 
```








