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
docker run --rm --name jaeger -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest
```



#### 构建镜像

```shell
## 注意项目里面的 一些服务地址 如mysql,redis 要修改成 容器ip ...不然跑不起来的。 
docker build -t crgo:v1  --build-arg VERSION="$(git rev-parse --short HEAD)",BUILDTIME="$(date +%FT%T)" -f  ./build/crgo/Dockerfile ./ 
```








