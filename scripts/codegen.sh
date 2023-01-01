find grpc/biz -type f -name "*.pb.go" -delete

absoulte_path=`pwd`
cd proto

protoc --go_out=$absoulte_path/grpc/biz/helloworld --go_opt=paths=source_relative \
    --go-grpc_out=$absoulte_path/grpc/biz/helloworld --go-grpc_opt=paths=source_relative \
    helloworld.proto


protoc --go_out=$absoulte_path/biz/session --go_opt=paths=source_relative \
    --go-grpc_out=$absoulte_path/biz/session --go-grpc_opt=paths=source_relative \
    session.proto


protoc --go_out=$absoulte_path/grpc/biz/bootstrap --go_opt=paths=source_relative \
    --go-grpc_out=$absoulte_path/grpc/biz/bootstrap --go-grpc_opt=paths=source_relative \
    bootstrap.proto




protoc --go_out=$absoulte_path/biz/user --go_opt=paths=source_relative \
    --go-grpc_out=$absoulte_path/biz/user --go-grpc_opt=paths=source_relative \
    user.proto





