#  my golang framework
  - 项目框架的搭建
  
  
 
#### Usage
```shell 
$ export GO111MODULE=on
$ go mod download
$ go get github.com/golang/protobuf/protoc-gen-go@v1.3.5
$ brew install protobuf

## Or Download https://github.com/protocolbuffers/protobuf/releases/download/v3.9.2/protoc-3.9.2-osx-x86_64.zip
$ unzip protoc-3.9.2-osx-x86_64.zip

$ go run main.go --help 
```


### consul 本地启动
```shell script
sudo consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul -node=n1 -bind=127.0.0.1 -ui -config-dir /etc/consul.d -rejoin -join 127.0.0.1 -client 0.0.0.0
```


 
#### 命令行工具集
- 单词格式转换
- 便利的时间工具
- SQL语句到结构体的转换



#### 目录说明
- config 应用配置 模仿laravel .env 配置
- http 基于gin的基础框架的http服务
- grpc 基于grpc框架的rpc服务
- infra  基础架构
   - conf 配置处理
   - log log处理
   - db  初始化mysql
   - redis 初始化redis
   - rabbitmq 初始化rabbitmq
   - sql2struct  表转结构体
   - timer 时间函数
   - util 工具包
   - word 字符转换处理
   - crtcp tcp长连接 (todo)
   
- scripts  脚本运行
- docs     文档说明   
- build    构建相关
   






