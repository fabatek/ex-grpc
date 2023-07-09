#!/usr/bin/env bash

set -e

cd "$(dirname "$0")"
cd ..

docker build . \
  -f ./scripts/docker/Dockerfile.protoc \
  --build-arg=PROTOC_VERSION=3.14.0 \
  --build-arg=PROTOC_GEN_GO_VERSION=v1.25.0 \
  --build-arg=PROTOC_GEN_GO_GRPC_VERSION=v1.1 \
  --tag=syns-datalake/protoc-gen-go-1.24.0

# for golang
docker run \
  -v "$(pwd)"/proto:/proto \
  -it syns-datalake/protoc-gen-go-1.24.0 \
  protoc \
    --go_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:/proto \
    --go-grpc_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:/proto \
    --go_opt=paths=source_relative \
    --go-grpc_opt=paths=source_relative \
    --proto_path=/proto \
    /proto/faba.proto     


 