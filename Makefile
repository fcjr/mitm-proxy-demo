.PHONY: default install build release clean

NAME=mitm-proxy-demo
VERSION := $(shell cat .VERSION)
BUILD_DIR=bin
LDFLAGS=""

.DEFAULT_GOAL :=
default:
	# default does nothing

install:
	go install -ldflags=$(LDFLAGS) cmd/$(NAME)/main.go

build:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(NAME)-$(VERSION)_darwin_amd64 -ldflags=$(LDFLAGS) cmd/$(NAME)/main.go
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(NAME)-$(VERSION)_windows_amd64.exe -ldflags=$(LDFLAGS) cmd/$(NAME)/main.go

release: build
	git tag v$(VERSION) && git push origin v$(VERSION)
	hub release create -m $(VERSION) -a $(BUILD_DIR)/$(NAME)-$(VERSION)_darwin_amd64 -a $(BUILD_DIR)/$(NAME)-$(VERSION)_windows_amd64.exe v$(VERSION)

clean:
	rm -rf $(BUILD_DIR)
