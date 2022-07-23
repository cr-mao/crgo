find grpc/biz -type f -name "*.pb.go" -delete
protoc --proto_path=grpc/proto --go_out=plugins=grpc:grpc ./grpc/proto/helloworld.proto
protoc --proto_path=grpc/proto --go_out=plugins=grpc:grpc ./grpc/proto/session.proto
