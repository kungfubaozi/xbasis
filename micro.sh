#!/usr/bin/env bash

#MICRO_REGISTRY=consul MICRO_API_HANDLER=rpc micro api
docker run -p 8080:8080 \
	-e MICRO_REGISTRY=consul \
	-e MICRO_REGISTRY_ADDRESS=192.168.80.67:8500 \
	-e MICRO_API_HANDLER=rpc \
	-e MICRO_REGISTER_TTL=30 \
	-e MICRO_CLIENT=grpc \
	-e MICRO_SERVER=grpc \
	-e MICRO_REGISTER_INTERVAL=15 \
	-e MICRO_API_NAMESPACE=gs.svc \
    microhq/micro api \
    --address=:8080