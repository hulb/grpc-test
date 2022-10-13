.PHONY: default
default: 
	./protoc/bin/protoc --proto_path=./protoc/include/google/protobuf/ --proto_path=./proto --go_out . --go_opt=module=github.com/hulb/grpc-test --go-grpc_out . --go-grpc_opt=module=github.com/hulb/grpc-test hello.proto