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

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build-c1api
build-c1api:
	echo "Building c1api openapi client... (source: /specs/c1-openapi.yaml)"
	rm -rf build/c1api
	mkdir -p build/c1api
	podman run --rm -v "${PWD}/build/c1api:/output" \
		-v "${PWD}/specs:/specs" \
		docker.io/openapitools/openapi-generator-cli generate \
		-i /specs/c1-openapi.yaml \
		-g go \
	    -o /output \
		--additional-properties=enumClassPrefix=true,hideGenerationTimestamp=true,structPrefix=true,disallowAdditionalPropertiesIfNotPresent=false,packageName=c1api,isGoSubmodule=true
	rm -rf build/c1api/go.mod build/c1api/go.sum
	rm -rf internal/c1api
	mkdir -p internal/c1api
	mv build/c1api internal/
	find internal/c1api \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i'' -e 's/GIT_USER_ID\/GIT_REPO_ID/conductorone\/cone\/internal/g'
	go mod tidy -v
	go mod vendor