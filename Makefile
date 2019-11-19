squzy: .gen_squzy

.gen_squzy:
	protoc -I./proto --go_out=plugins=grpc,request_context=true:./generated proto/squzy/service.proto;
