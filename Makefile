# Makefile for aten/artr - Aten Remote Task Runner

APP_NAME := artr
BUILD_DIR := build
VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := -X 'github.com/atenteccompany/artr/cmd.Version=$(VERSION)'

GO := go
GOFILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all build clean run test version

all: build

build:
	@echo "🔧 Building $(APP_NAME) version $(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME) .

run:
	@echo "🚀 Running $(APP_NAME)..."
	$(BUILD_DIR)/$(APP_NAME)

version:
	@echo "📦 Version: $(VERSION)"

clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

test:
	@echo "🧪 Running tests..."
	$(GO) test ./...

release:
	@echo "🚀 Releasing $(APP_NAME) version $(VERSION)..."
	@mkdir -p dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -ldflags "$(LDFLAGS)" -o dist/$(APP_NAME)-linux-amd64 .
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build -ldflags "$(LDFLAGS)" -o dist/$(APP_NAME)-darwin-amd64 .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64  $(GO) build -ldflags "$(LDFLAGS)" -o dist/$(APP_NAME)-windows-amd64.exe .
	@echo "✅ Binaries built in ./dist"


