.PHONY: proto

INCLUDE_DIR = /home/ubuntu/software/protoc-35.0/include
PROTO_FILES = $(wildcard proto/*.proto)

proto:
	@echo "proto_file = $(PROTO_FILES)"
	@echo "include_dir = $(INCLUDE_DIR)"
	@protoc --go_out=. --go-grpc_out=. \
	-I $(INCLUDE_DIR) \
	-I . \
	$(PROTO_FILES)						
# 指定输入路径为当前目录 .
# 指定导入的 protobuf 系统库 include 路径
# 指定导入的其他 proto package 路径
# 选中 ./proto 目录下的所有 proto 文件