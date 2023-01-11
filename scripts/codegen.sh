find grpc/biz -type f -name "*.pb.go" -delete

absoulte_path=`pwd`


# 利用 proto文件种的 go_package 来生成到指定位置。
protoc --go_out=. --go_opt=paths=import \
    --go-grpc_out=. --go-grpc_opt=paths=import \
    proto/helloworld.proto


# 直接指定绝对路径。
cd proto
protoc --go_out=$absoulte_path/biz/session --go_opt=paths=source_relative \
    --go-grpc_out=$absoulte_path/biz/session --go-grpc_opt=paths=source_relative \
    session.proto


protoc --go_out=$absoulte_path/grpc/biz/bootstrap --go_opt=paths=source_relative \
    --go-grpc_out=$absoulte_path/grpc/biz/bootstrap --go-grpc_opt=paths=source_relative \
    bootstrap.proto




protoc --go_out=$absoulte_path/biz/user --go_opt=paths=source_relative \
    --go-grpc_out=$absoulte_path/biz/user --go-grpc_opt=paths=source_relative \
    user.proto



protoc --go_out=$absoulte_path/biz/goods --go_opt=paths=source_relative \
    --go-grpc_out=$absoulte_path/biz/goods --go-grpc_opt=paths=source_relative \
    goods.proto


protoc --go_out=$absoulte_path/biz/inventory --go_opt=paths=source_relative \
    --go-grpc_out=$absoulte_path/biz/inventory --go-grpc_opt=paths=source_relative \
    inventory.proto



