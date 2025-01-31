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
	--go_out=pkg/grpc/task/v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=.bin/protoc-gen-go \
	--go-grpc_out=pkg/grpc/task/v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=.bin/protoc-gen-go-grpc \
	api/grpc/task/v1/task.proto

generate-rsa256-keypair:
	ssh-keygen -t rsa -b 4096 -m PEM -f jwt.key
	openssl rsa -in jwt.key -pubout -outform PEM -out jwt.key.pub

include .env
init.db:
	docker run --name=tasks-db -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -p ${POSTGRES_PORT}:5432 -d postgres

stop.db:
	docker rm -f tasks-db

exec.db:
	docker exec -it tasks-db bash -c "psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}"

migrate.up:
	migrate -source file://migrations -database ${POSTGRES_MIGRATIONS_URL} up

migrate.down:
	migrate -source file://migrations -database ${POSTGRES_MIGRATIONS_URL} down