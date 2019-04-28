#!/usr/bin/env bash


#permission
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. pb/form.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. pb/instance.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. pb/process.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. pb/history.proto