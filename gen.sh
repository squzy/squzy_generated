#!/usr/bin/env bash
protoc -I./proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,request_context=true:./generated/monitoring proto/proto/v1/squzy_monitoring.proto;
protoc -I./proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,request_context=true:./generated/storage proto/proto/v1/squzy_storage.proto;
protoc -I./proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,request_context=true:./generated/agent_server proto/proto/v1/squzy_agent_server.proto;
