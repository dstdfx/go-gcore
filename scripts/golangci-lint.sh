#!/bin/bash

echo "==> Running 'golangci-lint' ..."
golangci-lint run \
    --enable golint \
    --enable interfacer \
    --enable unconvert \
    --enable goconst \
    --enable gocyclo \
    --enable gofmt \
    --enable goimports \
    --enable depguard \
    --enable unparam \
    --enable nakedret \
    --enable prealloc \
    --enable gocritic \
    ./...
