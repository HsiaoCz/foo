gen:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/pv1/foo.proto

run:
	@go run cmd/main.go

tidy:
	@go mod tidy
