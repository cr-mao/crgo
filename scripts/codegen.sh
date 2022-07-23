find grpc/biz -type f -name "*.pb.go" -delete
protoc --proto_path=proto --go_out=plugins=grpc:grpc ./proto/helloworld.proto
protoc --proto_path=proto --go_out=plugins=grpc:grpc ./proto/session.proto
protoc --proto_path=proto --go_out=plugins=grpc:grpc ./proto/bootstrap.proto
