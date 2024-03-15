.PHONY: gen_go gen_python
all: gen_go gen_python

gen_go:
	protoc --go_out=. --go-grpc_out=. avatar_service.proto

gen_python:
	python -m grpc_tools.protoc -I. --python_out=./avatar --grpc_python_out=./avatar avatar_service.proto