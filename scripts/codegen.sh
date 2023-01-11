find grpc/biz -type f -name "*.pb.go" -delete

absoulte_path=`pwd`






# https://developers.google.com/protocol-buffers/docs/reference/go-generated#package
# 默认就是import 模式
protoc --go_out=.  --go_opt=paths=import  \
    --go-grpc_out=.  --go-grpc_opt=paths=import \
    proto/helloworld.proto


# 利用 proto文件种的 go_package 来生成到指定位置。
# 绝对路径,它是共用的，被其他proto 倒入， 暂时没找到好的办法， 不然其他proto 生成的 import 倒入路径有点问题。
protoc --proto_path=proto --go_out=$absoulte_path/biz/common --go_opt=paths=source_relative \
    --go-grpc_out=$absoulte_path/biz/common --go-grpc_opt=paths=source_relative \
    common.proto

# source_relative  模式， 和go_out， go-grpc_out 路径一致
protoc --proto_path=proto --go_out=$absoulte_path/biz/session --go_opt=paths=source_relative \
    --go-grpc_out=$absoulte_path/biz/session --go-grpc_opt=paths=source_relative \
    session.proto


protoc --proto_path=proto  --go_out=. --go-grpc_out=.    bootstrap.proto


# 用户服务
protoc --proto_path=proto  --go_out=. --go-grpc_out=.  user.proto

# 库存服务
protoc --proto_path=proto --go_out=.  --go-grpc_out=.  inventory.proto

# 商品服务
protoc --proto_path=proto --go_out=. --go-grpc_out=.  goods.proto
