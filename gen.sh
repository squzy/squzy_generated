#!/usr/bin/env bash
protoc -I./proto  --go_out=./generated proto/proto/v1/*.proto;
