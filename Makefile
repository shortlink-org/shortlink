.: generate

generate:
	@echo "proto generation link entity"
	@protoc \
	-I=. \
	--go_out=. \
	pkg/link/link.proto

	@echo "proto generation gRPC-web"
	@protoc \
	-I $(shell pwd) \
	-I=pkg/api/grpc-web \
	-I=third_party/googleapis \
	--go_out=plugins=grpc:pkg/api/grpc-web \
	--grpc-gateway_out=logtostderr=true:pkg/api/grpc-web \
	pkg/api/grpc-web/api.proto
