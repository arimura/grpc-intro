.PHONY: proto server client run-server run-client run-server-dev run-client-dev clean

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		hoge.proto

server:
	go build -o bin/server ./server

client:
	go build -o bin/client ./client

run-server: server
	./bin/server

run-client: client
	./bin/client

# Dev targets that run directly without building binaries
run-server-dev:
	go run ./server/main.go

run-client-dev:
	go run ./client/main.go

clean:
	rm -rf proto bin