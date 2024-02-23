# Service A
.PHONY: service-a/init
service-a/init:
	go install github.com/google/wire/cmd/wire@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go generate ./internal/infra/grpc
	go mod tidy

# Service B
.PHONY: service-b/init
service-b/init:
	go install github.com/google/wire/cmd/wire@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go generate ./internal/infra/grpc
	go mod tidy
