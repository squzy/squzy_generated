squzy: .gen_squzy

logger: .logger

.gen_squzy:
	protoc -I./proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,request_context=true:./generated proto/squzy/service.proto;

.logger:
	protoc -I./proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,request_context=true:./generated proto/logger/service.proto;
