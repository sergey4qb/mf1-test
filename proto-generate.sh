#!/bin/bash

mkdir -p ./proto/pb

CreateUserProto() {
  echo "-> Processing: user.proto"
  protoc -I="./proto" \
    --go_out="./proto/pb" \
    --go_opt=Muser.proto="." \
    --go-grpc_out=require_unimplemented_servers=false:"./proto/pb" \
    --go-grpc_opt=Muser.proto="." \
    --experimental_allow_proto3_optional \
    proto/user.proto
}

CreateUserProto
