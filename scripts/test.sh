#!/bin/bash

echo "==> Running 'gotest' ..."
go test -covermode=count -coverprofile=coverage.out -v ./gcore/.
