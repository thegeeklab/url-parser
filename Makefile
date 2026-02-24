# renovate: datasource=github-releases depName=mvdan/gofumpt
GOFUMPT_PACKAGE_VERSION := v0.9.2
# renovate: datasource=github-releases depName=golangci/golangci-lint
GOLANGCI_LINT_PACKAGE_VERSION := v2.10.1
# renovate: datasource=docker depName=docker.io/techknowlogick/xgo
XGO_PACKAGE_VERSION := go-1.26.0

EXECUTABLE := url-parser

DIST := dist
DIST_DIRS := $(DIST)
IMPORT := github.com/thegeeklab/$(EXECUTABLE)

GO ?= go
CWD ?= $(shell pwd)
PACKAGES ?= $(shell go list ./...)
SOURCES ?= $(shell find . -name "*.go" -type f)

GOFUMPT_PACKAGE ?= mvdan.cc/gofumpt@$(GOFUMPT_PACKAGE_VERSION)
XGO_PACKAGE ?= src.techknowlogick.com/xgo@latest
GOTESTSUM_PACKAGE ?= gotest.tools/gotestsum@latest

GENERATE ?=
XGO_TARGETS ?= linux/amd64,linux/arm-6,linux/arm-7,linux/arm64

TARGETOS ?= linux
TARGETARCH ?= amd64
ifneq ("$(TARGETVARIANT)","")
GOARM ?= $(subst v,,$(TARGETVARIANT))
endif
TAGS ?= netgo

ifndef VERSION
	ifneq ($(CI_COMMIT_TAG),)
		VERSION ?= $(subst v,,$(CI_COMMIT_TAG))
	else
		VERSION ?= $(shell git rev-parse --short HEAD)
	endif
endif

ifndef DATE
	DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%S%z")
endif

LDFLAGS += -s -w -X "main.BuildVersion=$(VERSION)" -X "main.BuildDate=$(DATE)"

.PHONY: all
all: clean build

.PHONY: clean
clean:
	$(GO) clean -i ./...
	rm -rf $(DIST_DIRS)

.PHONY: fmt
fmt:
	$(shell go env GOPATH)/bin/gofumpt -extra -w $(SOURCES)

.PHONY: golangci-lint
golangci-lint:
	$(shell go env GOPATH)/bin/golangci-lint run

.PHONY: lint
lint: golangci-lint

.PHONY: generate
generate:
	$(GO) generate $(GENERATE)

.PHONY: test
test:
	$(shell go env GOPATH)/bin/gotestsum --no-color=false -- -coverprofile=coverage.out $(PACKAGES)

.PHONY: build
build: $(DIST)/$(EXECUTABLE)

$(DIST)/$(EXECUTABLE): $(SOURCES)
	GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) GOARM=$(GOARM) $(GO) build -v -tags '$(TAGS)' -ldflags '-extldflags "-static" $(LDFLAGS)' -o $@ ./cmd/$(EXECUTABLE)

$(DIST_DIRS):
	mkdir -p $(DIST_DIRS)

.PHONY: xgo
xgo: | $(DIST_DIRS)
	$(shell go env GOPATH)/bin/xgo -go $(XGO_PACKAGE_VERSION) -v -ldflags '-extldflags "-static" $(LDFLAGS)' -tags '$(TAGS)' -targets '$(XGO_TARGETS)' -out $(EXECUTABLE) --pkg cmd/$(EXECUTABLE) .
	cp /build/* $(CWD)/$(DIST)
	ls -l $(CWD)/$(DIST)

.PHONY: checksum
checksum:
	cd $(DIST); $(foreach file,$(wildcard $(DIST)/$(EXECUTABLE)-*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)
	ls -l $(CWD)/$(DIST)

.PHONY: release
release: xgo checksum

.PHONY: deps
deps:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_LINT_PACKAGE_VERSION)
	$(GO) mod download
	$(GO) install $(GOFUMPT_PACKAGE)
	$(GO) install $(XGO_PACKAGE)
	$(GO) install $(GOTESTSUM_PACKAGE)
