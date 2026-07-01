APP_NAME := dnotes
CMD_PATH := .
BIN_DIR := bin
BIN_PATH := $(BIN_DIR)/$(APP_NAME)

.PHONY: help build run test fmt vet tidy clean

help:
	@echo "Available targets:"
	@echo "  make build  - Build the CLI"
	@echo "  make run    - Run the CLI"
	@echo "  make test   - Run tests"
	@echo "  make fmt    - Format code"
	@echo "  make vet    - Run go vet"
	@echo "  make tidy   - Tidy modules"
	@echo "  make clean  - Remove build artifacts"

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_PATH) $(CMD_PATH)

run:
	go run $(CMD_PATH)

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -rf $(BIN_DIR)
