CC=go
PROJECT_PATH=$(shell pwd)
PROJECT_NAME=go-channel-signal-app
MODULE_NAME=programming
TARGET_DIR=bin
VERSION_FILE=version.txt
VERSION=$$(cat version.txt)
BUILD_NUM_FILE=build_num.txt
OUTPUT=$(PROJECT_PATH)/$(TARGET_DIR)/$(MODULE_NAME)_$(VERSION).$$(cat $(BUILD_NUM_FILE))
MAIN_DIR=/main
LDFLAGS=-X main.BUILD_TIME=`date -u '+%Y-%m-%d_%H:%M:%S'`
LDFLAGS+=-X main.GIT_HASH=`git rev-parse HEAD`
LDFLAGS+=-X main.BUILD_NUMBER=$$(cat $(BUILD_NUM_FILE))
LDFLAGS+=-X main.VERSION_NUMBER=$$(cat $(VERSION_FILE))
LDFLAGS+=-s -w

all: clean build

build:
	@echo $$(($$(cat $(BUILD_NUM_FILE)) + 1 )) > $(BUILD_NUM_FILE)
	CGO_ENABLED=0 GOOS=linux go build -ldflags "$(LDFLAGS)" -o $(OUTPUT) $(PROJECT_PATH)$(MAIN_DIR)
	cp $(OUTPUT) ./ex_$(MODULE_NAME)

clean:
	rm -f $(PROJECT_PATH)/ex_*

