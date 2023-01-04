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
```

#### 构建镜像

```
## 注意项目里面的 一些服务地址 如mysql,redis 要修改成 容器ip ...不然跑不起来的。 
docker build -t crgo:v1  --build-arg VERSION="$(git rev-parse --short HEAD)",BUILDTIME="$(date +%FT%T)" -f  ./build/crgo/Dockerfile ./ 
```

#### consul 本地启动
```shell script
# https://developer.hashicorp.com/consul/downloads 
# sudo mkdir -p /etc/consul.d
sudo consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul -node=n1 -bind=127.0.0.1 -ui -config-dir /etc/consul.d -rejoin -join 127.0.0.1 -client 0.0.0.0
```

#### prometheus 本地启动
```text
docker pull prom/prometheus:v2.20.1
docker run -d --name prometheus -p 9090:9090 -v /Users/mac/code/crgo/build/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus:v2.20.1
```






#### 命令行工具集
- 单词格式转换
- 便利的时间工具
- SQL语句到结构体的转换

#### 目录说明

- biz 客户端 grpc,http公共处理部分 如请求，会话，认证, http,grpc一些要共用pb.go文件的也房子啊这里
- build 构建相关
- cmd 命令行，初始化工作
- config 应用配置 模仿laravel .env 配置
- grpc 基于grpc框架的rpc服务
- http 基于gin的基础框架的http服务
- infra 基础架构
    - bizerror 响应处理
    - conf 配置处理
    - crtcp tcp长连接服务 (todo)
    - db 初始化mysql
    - discovery 基于consul服务注册、注销，发现
    - log log处理
    - job 定时任务、守护进程
    - rabbitmq (todo)
    - redis 初始化redis
    - sql2struct 表转结构体
    - timer 时间函数
    - util 工具包
    - word 字符转换处理
- model 表模型 基于gorm
- proto 所有的proto文件目录
- scripts 脚本运行
- Makefile make serve 启动http,grpc服务
- main.go 入口
- config.local.toml 本地配置文件
- sentinel.yaml 阿里sentinel配置文件
   






