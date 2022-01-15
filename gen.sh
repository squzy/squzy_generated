#!/usr/bin/env bash
protoc -I./proto  --go_out=plugins=grpc:./generated proto/proto/v1/*.proto;
