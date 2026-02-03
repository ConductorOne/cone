GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
BUILD_DIR = dist/${GOOS}_${GOARCH}

ifeq ($(GOOS),windows)
OUTPUT_PATH = ${BUILD_DIR}/cone.exe
else
OUTPUT_PATH = ${BUILD_DIR}/cone
endif

.PHONY: build
build:
	go build -o ${OUTPUT_PATH} ./cmd/cone

.PHONY: update-deps
update-deps:
	go get -d -u ./...
	go mod tidy -v
	go mod vendor

.PHONY: add-deps
add-dep:
	go mod tidy -v
	go mod vendor

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: install-hooks
install-hooks:
	@echo '#!/bin/sh' > .git/hooks/pre-push
	@echo 'make lint' >> .git/hooks/pre-push
	@chmod +x .git/hooks/pre-push
	@echo "Installed pre-push hook to run 'make lint'"

