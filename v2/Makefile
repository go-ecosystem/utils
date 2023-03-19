SHELL := /bin/bash
BASEDIR = $(shell pwd)

all: fmt imports mod lint test
install-pre-commit:
	brew install pre-commit
install-git-hooks:
	pre-commit install --hook-type commit-msg
	pre-commit install
run-pre-commit:
	pre-commit run --all-files
fmt:
	gofmt -w .
imports:
ifeq (, $(shell which goimports))
	go install golang.org/x/tools/cmd/goimports@latest
endif
	goimports -w .
mod:
	go mod tidy
lint: mod
ifeq (, $(shell which golangci-lint))
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
endif
	golangci-lint run -c .golangci.yml
.PHONY: test
test: mod
	go test -gcflags=-l -coverpkg=./... -coverprofile=coverage.data ./...
help:
	@echo "fmt - format the source code"
	@echo "imports - goimports"
	@echo "mod - go mod tidy"
	@echo "lint - run golangci-lint"
	@echo "test - unit test"
