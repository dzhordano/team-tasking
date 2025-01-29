LOCAL_BIN:=$(CURDIR)/.bin


install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-task-api-v1

generate-task-api-v1:
	mkdir -p pkg/grpc/task/v1
	protoc --proto_path api/grpc/task/v1 \
	--go_out=pkg/grpc/task/v1/task_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=.bin/protoc-gen-go \
	--go-grpc_out=pkg/grpc/task/v1/task_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=.bin/protoc-gen-go-grpc \
	api/grpc/task/v1/task.proto
