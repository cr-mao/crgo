find grpcproject/biz -type f -name "*.pb.go" -delete
protoc --proto_path=grpcproject/proto --go_out=plugins=grpc:grpcproject ./grpcproject/proto/helloworld.proto
