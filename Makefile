SHELL := /bin/bash
BASEDIR = $(shell pwd)

all: mod fmt imports lint test
mod:
	go mod tidy
fmt:
	gofmt -w .
imports:	
	goimports -w .	
lint:
	golangci-lint run
.PHONY: test
test:
	sh scripts/test.sh
help:
	@echo "mod - go mod tidy"
	@echo "fmt - go format"	
	@echo "imports - go imports"
	@echo "lint - run golangci-lint"
	@echo "test - unit test"