gen:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative example.proto

run-server:
	go run server/server.go

run-client:
	go run client/client.go