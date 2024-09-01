SHELL = /usr/bin/env bash
.SHELLFLAGS = -ecuo pipefail

COVERPROFILE ?= coverage.out

.PHONY: gif/generate
gif/generate:
	./scripts/gif.sh

.PHONY: test
test:
	go test -v -cover -coverprofile=$(COVERPROFILE) ./...

.PHONY: test/coverage
test/coverage:
	go tool cover -html=$(COVERPROFILE) -o coverage.html

.PHONY: deps/update
deps/update:
	go mod tidy
	go get -u ./...