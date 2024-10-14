# Project name
PROJECT_NAME := ProjectD

# Directories
PROTO_DIR := pkg/proto
SERVER_DIR := cmd/server
CLIENT_DIR := cmd/client

# Protobuf files
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

# Go commands
GOCMD := go
GOBUILD := $(GOCMD) build
GORUN := $(GOCMD) run
GOTEST := $(GOCMD) test
GOCLEAN := $(GOCMD) clean
GOMOD := $(GOCMD) mod

# Protobuf commands
PROTOC := protoc
PROTOC_GEN_GO := protoc-gen-go
PROTOC_GEN_GO_GRPC := protoc-gen-go-grpc

# Targets
.PHONY: all build clean test run-server run-client generate

all: build

build: generate
	@echo "Building server and client..."
	$(GOBUILD) -o bin/server $(SERVER_DIR)/main.go
	$(GOBUILD) -o bin/client $(CLIENT_DIR)/main.go

clean:
	@echo "Cleaning up..."
	$(GOCLEAN)
	rm -rf bin

test:
	@echo "Running tests..."
	$(GOTEST) ./...

run-server:
	@echo "Running server..."
	$(GORUN) $(SERVER_DIR)/main.go

run-client:
	@echo "Running client..."
	$(GORUN) $(CLIENT_DIR)/main.go

generate: $(PROTO_FILES)
	@echo "Generating Go code from proto files..."
	$(PROTOC) --go_out=$(PROTO_DIR) --go-grpc_out=$(PROTO_DIR) $(PROTO_FILES)

$(PROTO_DIR)/%.pb.go: $(PROTO_DIR)/%.proto
	$(PROTOC) --go_out=$(PROTO_DIR) --go-grpc_out=$(PROTO_DIR) $<

$(PROTO_DIR)/%_grpc.pb.go: $(PROTO_DIR)/%.proto
	$(PROTOC) --go_out=$(PROTO_DIR) --go-grpc_out=$(PROTO_DIR) $<
