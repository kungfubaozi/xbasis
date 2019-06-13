#!/usr/bin/env bash

#commons
protoc -I ${GOPATH}/src --proto_path=. --go_out=:. commons/dto/commons.proto

#permission
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/group.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/role.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/function.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/binding.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/dat.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/structure.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/verify.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/inner/verification.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. permission/pb/inner/accessible.proto

#user
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/update.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/login.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/safety.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/register.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/invite.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/active.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/grant.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/oauth.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/message.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/userinfo.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/authorization.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/inner/message.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. user/pb/inner/user.proto

#safety
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. safety/pb/blacklist.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. safety/pb/userlock.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. safety/pb/locking.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. safety/pb/inner/security.proto

#authentication
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. authentication/pb/inner/auth.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. authentication/pb/inner/token.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. authentication/pb/router.proto

#application
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. application/pb/application.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. application/pb/settings.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. application/pb/inner/status.proto
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. application/pb/inner/sync.proto

#analysis
protoc -I ${GOPATH}/src --proto_path=. --micro_out=. --go_out=:. analysis/pb/logger.proto

#micro
#protoc -I ${GOPATH}/src -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#        --proto_path=. --micro_out=. --go_out=:. permission/pb/test.proto

#process api
#protoc -I. \
#  -I${GOPATH}/src \
#  -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#   --proto_path=. --grpc-gateway_out=logtostderr=true:/Users/Richard/Desktop/Development/Golang/src/konekko.me/gosion/permission/pb/gateway \
#   permission/pb/test.proto