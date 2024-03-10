gen_go:
	protoc --go_out=. --go-grpc_out=. avatar_service.proto