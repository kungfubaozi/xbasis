#!/usr/bin/env bash

#commons
protoc -I ${GOPATH}/src --proto_path=. --go_out=:. commons/dto/commons.proto

#permission
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/group.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/role.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/function.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/binding.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/verification.proto

#user
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/update.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/login.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/safety.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/register.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/user.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/invite.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/active.proto

#safety
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. safety/pb/frozen.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. safety/pb/blacklist.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. safety/pb/locking.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. safety/pb/status.proto

#authentication
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. authentication/pb/auth.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. authentication/pb/token.proto

#application
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. application/pb/application.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. application/pb/settings.proto


#micro
protoc -I ${GOPATH}/src -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --proto_path=. --micro_out=. --go_out=:. permission/pb/test.proto

#process api
#protoc -I. \
#  -I${GOPATH}/src \
#  -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#   --proto_path=. --grpc-gateway_out=logtostderr=true:/Users/Richard/Desktop/Development/Golang/src/konekko.me/gosion/permission/pb/gateway \
#   permission/pb/test.proto