.: generate

dep:
	@echo "install protoc"
	@sudo ./ops/scripts/install-protobuf.sh
	@echo "install protoc addons"
	@go get -u github.com/golang/protobuf/proto
	@go get -u github.com/golang/protobuf/protoc-gen-go
	@go get -u moul.io/protoc-gen-gotemplate
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

generate:
	@echo "proto generation link entity"
	@protoc -I/usr/local/include -I. \
	--gotemplate_out=all=true,template_dir=pkg/api/graphql/template:pkg/api/graphql \
	--go_out=plugins=grpc:. \
	pkg/link/link.proto

	@echo "proto generation gRPC-web"
	@protoc -I/usr/local/include -I. \
	-I=pkg/api/grpc-web \
	-I=third_party/googleapis \
	--plugin=protoc-gen-grpc-gateway=${GOPATH}/bin/protoc-gen-grpc-gateway \
	--go_out=plugins=grpc:. \
	--swagger_out=logtostderr=true,allow_delete_body=true:. \
	--grpc-gateway_out=logtostderr=true,allow_delete_body=true:. \
	pkg/api/grpc-web/api.proto
	@mv pkg/api/grpc-web/api.swagger.json docs/api.swagger.json

	@echo "Generate go static"
	@go generate pkg/api/graphql/schema/schema.go

golint:
	@golangci-lint run

test:
	@go test -race -coverprofile=coverage.txt -covermode=atomic

run:
	@docker-compose \
         -f docker-compose.yaml \
         -f ops/docker-compose/database/redis.yaml \
         -f ops/docker-compose/gataway/traefik.yaml \
         -f ops/docker-compose/tooling/opentracing.yaml \
         up -d

down:
	@docker-compose down --remove-orphans
