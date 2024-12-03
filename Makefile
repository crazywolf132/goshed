# Build variables
BINARY_NAME=goshed
VERSION=$(shell git describe --tags --always --dirty)
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

# Environment
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

# Docker
DOCKER_IMAGE=goshed
DOCKER_TAG=latest

# Colors
COLOR_RESET=\033[0m
COLOR_BOLD=\033[1m
COLOR_BLUE=\033[34m
COLOR_GREEN=\033[32m
COLOR_CYAN=\033[36m

# Default target
.PHONY: all
all: clean build test

# Build the application
.PHONY: build
build: 
	@printf "$(COLOR_BLUE)$(COLOR_BOLD)Building $(BINARY_NAME)...$(COLOR_RESET)\n"
	@go build ${LDFLAGS} -o $(GOBIN)/$(BINARY_NAME) ./cmd/goshed

# Cross compile for multiple platforms
.PHONY: build-all
build-all: build-linux build-windows build-darwin

.PHONY: build-linux
build-linux:
	@printf "$(COLOR_BLUE)Building for Linux...$(COLOR_RESET)\n"
	@GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o $(GOBIN)/$(BINARY_NAME)-linux-amd64 ./cmd/goshed

.PHONY: build-windows
build-windows:
	@printf "$(COLOR_BLUE)Building for Windows...$(COLOR_RESET)\n"
	@GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o $(GOBIN)/$(BINARY_NAME)-windows-amd64.exe ./cmd/goshed

.PHONY: build-darwin
build-darwin:
	@printf "$(COLOR_BLUE)Building for macOS...$(COLOR_RESET)\n"
	@GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o $(GOBIN)/$(BINARY_NAME)-darwin-amd64 ./cmd/goshed

# Run the application
.PHONY: run
run: build
	@printf "$(COLOR_GREEN)Running $(BINARY_NAME)...$(COLOR_RESET)\n"
	@$(GOBIN)/$(BINARY_NAME)

# Clean build files
.PHONY: clean
clean:
	@printf "$(COLOR_BLUE)Cleaning build files...$(COLOR_RESET)\n"
	@rm -rf $(GOBIN)
	@go clean
	@rm -f coverage.out

# Run tests
.PHONY: test
test:
	@printf "$(COLOR_CYAN)Running tests...$(COLOR_RESET)\n"
	@go test -v -race ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@printf "$(COLOR_CYAN)Running tests with coverage...$(COLOR_RESET)\n"
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

# Lint the code
.PHONY: lint
lint:
	@printf "$(COLOR_CYAN)Linting code...$(COLOR_RESET)\n"
	@golangci-lint run ./...

# Format the code
.PHONY: fmt
fmt:
	@printf "$(COLOR_CYAN)Formatting code...$(COLOR_RESET)\n"
	@go fmt ./...

# Check for security vulnerabilities
.PHONY: security
security:
	@printf "$(COLOR_CYAN)Checking for security vulnerabilities...$(COLOR_RESET)\n"
	@gosec ./...

# Install development dependencies
.PHONY: deps
deps:
	@printf "$(COLOR_BLUE)Installing development dependencies...$(COLOR_RESET)\n"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/securego/gosec/v2/cmd/gosec@latest

# Update dependencies
.PHONY: update-deps
update-deps:
	@printf "$(COLOR_BLUE)Updating dependencies...$(COLOR_RESET)\n"
	@go get -u ./...
	@go mod tidy

# Docker targets
.PHONY: docker-build
docker-build:
	@printf "$(COLOR_BLUE)Building Docker image...$(COLOR_RESET)\n"
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

.PHONY: docker-run
docker-run: docker-build
	@printf "$(COLOR_GREEN)Running Docker container...$(COLOR_RESET)\n"
	@docker run --rm -it $(DOCKER_IMAGE):$(DOCKER_TAG)

# Install the binary
.PHONY: install
install: build
	@printf "$(COLOR_GREEN)Installing $(BINARY_NAME)...$(COLOR_RESET)\n"
	@go install ./cmd/goshed

# Generate documentation
.PHONY: docs
docs:
	@printf "$(COLOR_CYAN)Generating documentation...$(COLOR_RESET)\n"
	@go doc -all > DOCUMENTATION.md

# Show help
.PHONY: help
help:
	@printf "$(COLOR_BOLD)Available targets:$(COLOR_RESET)\n"
	@printf "$(COLOR_BLUE)  build$(COLOR_RESET)          - Build the application\n"
	@printf "$(COLOR_BLUE)  build-all$(COLOR_RESET)      - Build for all platforms\n"
	@printf "$(COLOR_BLUE)  run$(COLOR_RESET)            - Run the application\n"
	@printf "$(COLOR_BLUE)  clean$(COLOR_RESET)          - Clean build files\n"
	@printf "$(COLOR_BLUE)  test$(COLOR_RESET)           - Run tests\n"
	@printf "$(COLOR_BLUE)  test-coverage$(COLOR_RESET)  - Run tests with coverage\n"
	@printf "$(COLOR_BLUE)  lint$(COLOR_RESET)           - Lint the code\n"
	@printf "$(COLOR_BLUE)  fmt$(COLOR_RESET)            - Format the code\n"
	@printf "$(COLOR_BLUE)  security$(COLOR_RESET)       - Check for security vulnerabilities\n"
	@printf "$(COLOR_BLUE)  deps$(COLOR_RESET)           - Install development dependencies\n"
	@printf "$(COLOR_BLUE)  update-deps$(COLOR_RESET)    - Update dependencies\n"
	@printf "$(COLOR_BLUE)  docker-build$(COLOR_RESET)   - Build Docker image\n"
	@printf "$(COLOR_BLUE)  docker-run$(COLOR_RESET)     - Run in Docker container\n"
	@printf "$(COLOR_BLUE)  install$(COLOR_RESET)        - Install the binary\n"
	@printf "$(COLOR_BLUE)  docs$(COLOR_RESET)           - Generate documentation\n"

.DEFAULT_GOAL := help 