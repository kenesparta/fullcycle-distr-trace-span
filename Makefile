.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go generate ./internal/temperature/infra/grpc
	go mod tidy

.PHONY: service-a/build
service-a/build:
	GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/service_a

.PHONY: service-b/build
service-b/build:
	GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/service_b
