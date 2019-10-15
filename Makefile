.: generate

generate:
	@echo "proto generation link entity"
	@protoc -I=pkg/internal/link \
	--go_out=pkg/internal/link \
	pkg/internal/link/link.proto
