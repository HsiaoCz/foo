run: build
	@./bin/foo

build:
	@go build -o bin/foo main.go

test:
	@go test -v ./...

gen:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/pv1/foo.proto

tidy:
	@go mod tidy

.PHONY: run build gen test tidy