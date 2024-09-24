# Variables
PROTOC=protoc
PROTO_SRC=protos
OUT_DIR=protos/auth
PROTO_FILES=$(PROTO_SRC)/auth.proto

# Default target
all: generate

# Generate both Go and gRPC files from proto files
generate:
	$(PROTOC) --proto_path=$(PROTO_SRC) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

# Clean up generated files (optional)
clean:
	rm -f $(OUT_DIR)/*.pb.go $(OUT_DIR)/*.grpc.pb.go

.PHONY: all generate clean
