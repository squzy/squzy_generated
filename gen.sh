#!/usr/bin/env bash
protoc -I./proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,request_context=true:./generated proto/proto/v1/*.proto;
