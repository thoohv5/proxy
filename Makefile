.PHONY: all
all: build run

BUILD_ENV       :=
GO				:= go
BUILD_TIME      := $(shell date +"%Y-%m-%d")
GIT_BRANCH      := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT_SHA1     := $(shell git rev-parse HEAD)

.PHONY: build
build:
	@rm -rf build && mkdir -p build && $(BUILD_ENV) $(GO) build -ldflags "-X main.BuildTime=${BUILD_TIME} -X main.GitBranch=${GIT_BRANCH} -X main.CommitSHA1=${COMMIT_SHA1}" -o build cmd/proxy/proxy.go

.PHONY: run
run:
	$(GO) run cmd/proxy/proxy.go
