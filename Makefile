.DEFAULT_GOAL := run

generate_protoc:
	protoc --go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=require_unimplemented_servers=false,paths=source_relative \
	ecommerce/product_info.proto

run:
	go run ./main.go

build: generate_protoc
	go build -o ./bin/server ./*.go