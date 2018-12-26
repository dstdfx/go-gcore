#!/bin/bash

echo "==> Running 'golangci-lint' ..."
golangci-lint run ./...
