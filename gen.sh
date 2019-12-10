#!/usr/bin/env bash
protoc -I./proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,request_context=true:./generated/server proto/proto/v1/server.proto;
protoc -I./proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,request_context=true:./generated/storage proto/proto/v1/storage.proto;
protoc -I./proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,request_context=true:./generated/agent proto/proto/v1/agent.proto;
