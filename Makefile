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

.PHONY: all

all: upx

linux: linux-amd64 linux-386 linux-arm64 linux-arm

windows: windows-amd64 windows-386 windows-arm64

darwin: darwin-amd64 darwin-arm64

frontend:
	yarn install
	yarn build

linux-amd64: frontend
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-$@-original $(SRC_DIR)

linux-arm64: frontend
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-$@-original $(SRC_DIR)

linux-386: frontend
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-$@-original $(SRC_DIR)

linux-arm: frontend
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-$@-original $(SRC_DIR)

windows-amd64: frontend
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-$@-original.exe $(SRC_DIR)

windows-386: frontend
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-$@-original.exe $(SRC_DIR)

windows-arm64: frontend
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-$@-original.exe $(SRC_DIR)

darwin-amd64: frontend
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-$@-original $(SRC_DIR)

darwin-arm64: frontend
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-$@-original $(SRC_DIR)

upx: linux windows darwin
	@mkdir -p $(OUTPUT_DIR)
	upx -$(UPX_LEVEL) $(OUTPUT_DIR)/$(APP_NAME)-linux-amd64-original -o $(OUTPUT_DIR)/$(APP_NAME)-linux-amd64
	upx -$(UPX_LEVEL) $(OUTPUT_DIR)/$(APP_NAME)-linux-386-original -o $(OUTPUT_DIR)/$(APP_NAME)-linux-386
	upx -$(UPX_LEVEL) $(OUTPUT_DIR)/$(APP_NAME)-linux-arm64-original -o $(OUTPUT_DIR)/$(APP_NAME)-linux-arm64
	upx -$(UPX_LEVEL) $(OUTPUT_DIR)/$(APP_NAME)-linux-arm-original -o $(OUTPUT_DIR)/$(APP_NAME)-linux-arm
	upx -$(UPX_LEVEL) $(OUTPUT_DIR)/$(APP_NAME)-windows-amd64-original.exe -o $(OUTPUT_DIR)/$(APP_NAME)-windows-amd64.exe
	upx -$(UPX_LEVEL) $(OUTPUT_DIR)/$(APP_NAME)-windows-386-original.exe -o $(OUTPUT_DIR)/$(APP_NAME)-windows-386.exe
	upx -$(UPX_LEVEL) $(OUTPUT_DIR)/$(APP_NAME)-windows-arm64-original.exe -o $(OUTPUT_DIR)/$(APP_NAME)-windows-arm64.exe
	upx -$(UPX_LEVEL) $(OUTPUT_DIR)/$(APP_NAME)-darwin-amd64-original -o $(OUTPUT_DIR)/$(APP_NAME)-darwin-amd64 --force-macos
	upx -$(UPX_LEVEL) $(OUTPUT_DIR)/$(APP_NAME)-darwin-arm64-original -o $(OUTPUT_DIR)/$(APP_NAME)-darwin-arm64 --force-macos
	rm -v $(OUTPUT_DIR)/*-original*


clean:
	rm -vrf $(OUTPUT_DIR)
	rm -vrf $(SRC_DIR)/dist