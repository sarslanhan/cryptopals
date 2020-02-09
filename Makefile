ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
SHELL := /bin/sh

VERSION ?= $(shell git describe --tags --always --dirty --abbrev=0)

SOURCES = $(shell find $(ROOT_DIR) -name "*.go" -print | grep -v /vendor/ )
COVERAGE_DIR ?= $(PWD)/coverage

export GO111MODULE = on

.PHONY: test check

default: all

all: build check

check: checkfmt test lint

build:
	go build ./...

.PHONY: test
test:
	GIN_MODE=test go test -race -v ./...

checkfmt:
	@[ -z $$(gofmt -l $(SOURCES)) ] || (echo "Sources not formatted correctly. Fix by running: make fmt" && false)

fmt: $(SOURCES)
	gofmt -s -w $(SOURCES)

lint:
	golangci-lint run
