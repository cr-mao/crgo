# my golang 大杂烩
   随心写....
  - 项目框架的搭建
  - 集成一些常用工具集
  - 面试题解题 
  
  
 
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


 
#### 命令行工具集
- 单词格式转换
- 便利的时间工具
- SQL语句到结构体的转换



#### 目录说明

- http 基于gin的基础框架的http服务
- grpc 基于grpc框架的rpc服务
- infra  基础架构
   - conf 配置处理
   - log log处理
   - db  初始化mysql
   - redis 初始化redis
- scripts  脚本运行
- docs     文档说明   
- build    构建相关
- module 一些模块实现案例以及面试题 (**和项目没任何关系**)
   






