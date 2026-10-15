APP_NAME	?= data-center
SRC_DIR		?= .
OUTPUT_DIR	?= build

VERSION		?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT      ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
BUILD_TIME  ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')

UPX_LEVEL	?= 9

LDFLAGS     := -s -w \
               -X 'main.Version=$(VERSION)' \
               -X 'main.GitCommit=$(COMMIT)' \
               -X 'main.BuildTime=$(BUILD_TIME)'

.PHONY: all run-backend run-frontend

all: clean prepare linux windows darwin

linux: linux-amd64 linux-386 linux-arm64 linux-arm

windows: windows-amd64 windows-386 windows-arm64

darwin: darwin-amd64 darwin-arm64

prepare: make-dir build-frontend install-swaggo build-docs

run-frontend:
	@yarn dev --host

run-backend:
	@go run . run

make-dir:
	@mkdir -p $(OUTPUT_DIR)

install-swaggo:
	go install github.com/swaggo/swag/cmd/swag@latest

build-docs:
	swag init -g main.go -o docs

build-frontend:
	@yarn install
	@yarn build

define build-bin
	@echo "Building $(APP_NAME)-$(1)-$(2)";\
	if [ "$(1)" = "windows" ]; then\
		EXT=.exe;\
	else\
		EXT=;\
	fi;\
	if [ "$(1)" = "darwin" ]; then\
		UPX_FLAG=--force-macos;\
	else\
		UPX_FLAG=;\
	fi;\
	echo "    Building original binary";\
	CGO_ENABLED=0 GOOS=$(1) GOARCH=$(2) go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-$(1)-$(2)-original$$EXT $(SRC_DIR);\
	echo "    Building upx binary";\
	upx -q -$(UPX_LEVEL) $(OUTPUT_DIR)/$(APP_NAME)-$(1)-$(2)-original$$EXT -o $(OUTPUT_DIR)/$(APP_NAME)-$(1)-$(2)$$EXT $$UPX_FLAG;\
	echo "    Cleaning up original binary";\
	rm -vf $(OUTPUT_DIR)/$(APP_NAME)-$(1)-$(2)-original$$EXT;
endef

linux-amd64:
	$(call build-bin,linux,amd64)

linux-arm64:
	$(call build-bin,linux,arm64)

linux-386:
	$(call build-bin,linux,386)

linux-arm:
	$(call build-bin,linux,arm)

windows-amd64:
	$(call build-bin,windows,amd64)

windows-386:
	$(call build-bin,windows,386)

windows-arm64:
	$(call build-bin,windows,arm64)

darwin-amd64:
	$(call build-bin,darwin,amd64)

darwin-arm64:
	$(call build-bin,darwin,arm64)

clean:
	@echo "Cleaning up"
	@rm -rf $(OUTPUT_DIR)
	@rm -rf $(SRC_DIR)/dist