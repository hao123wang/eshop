.PHONY: proto

INCLUDE_DIR = /home/ubuntu/software/protoc-35.0/include
PROTO_FILES = $(wildcard files/*.proto)
proto:
	@echo "proto_files = $(PROTO_FILES)"
	@echo "include_dir = $(INCLUDE_DIR)"
	@protoc --go_out=. --go-grpc_out=. \
	-I $(INCLUDE_DIR) \
	-I . \
	$(PROTO_FILES)