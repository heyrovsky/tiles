BINARY_NAME        := tiles
CONFIG_PACKAGE     := github.com/heyrovsky/tiles/config

# Auto-detect values unless provided
VERSION            ?= $(shell git describe --tags --always --dirty | sed 's/^v//')
GIT_HASH           ?= $(shell git rev-parse HEAD)
GO_VERSION         ?= $(shell go version | sed -E 's/.*go([0-9.]+).*/\1/')


BUILD_FLAGS        := -trimpath -buildvcs=false -ldflags "\
	-w -s \
	-X '$(CONFIG_PACKAGE).APP_VERSION=$(VERSION)' \
	-X '$(CONFIG_PACKAGE).APP_NAME=$(BINARY_NAME)' \
	-X '$(CONFIG_PACKAGE).APP_GIT_HASH=$(GIT_HASH)' \
	-X '$(CONFIG_PACKAGE).APP_GO_VERSION=$(GO_VERSION)'"

PLATFORMS          := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64

.PHONY: all build clean release deps info

all: build

## Print debug build information
info:
	@echo "VERSION     = $(VERSION)"
	@echo "GIT_HASH    = $(GIT_HASH)"
	@echo "GO_VERSION  = $(GO_VERSION)"
	@echo "BUILD_DEPS  = $(BUILD_DEPS)"
	@echo "BUILD_FLAGS = $(BUILD_FLAGS)"

build:
	@echo "Building static binary $(BINARY_NAME)..."
	CGO_ENABLED=0 GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH) \
	go build $(BUILD_FLAGS) -o $(BINARY_NAME) cmd/app/main.go
	@echo "Static build complete: ./$(BINARY_NAME)"

clean:
	@echo "Cleaning..."
	rm -rf $(BINARY_NAME) dist
	@echo "Clean complete."

release:
	@echo "Building static release binaries..."
	mkdir -p dist
	@for PLATFORM in $(PLATFORMS); do \
		OS=$$(echo $$PLATFORM | cut -d/ -f1); \
		ARCH=$$(echo $$PLATFORM | cut -d/ -f2); \
		EXT=$$( [ "$$OS" = "windows" ] && echo ".exe" || echo "" ); \
		OUTPUT=dist/$(BINARY_NAME)-$$OS-$$ARCH$$EXT; \
		echo "Building $$OUTPUT..."; \
		CGO_ENABLED=0 GOOS=$$OS GOARCH=$$ARCH go build $(BUILD_FLAGS) -o "$$OUTPUT" cmd/app/main.go || exit 1; \
	done
	@echo "All static binaries are in ./dist"

deps:
	@echo "Listing go module dependencies..."
	go list -m all
