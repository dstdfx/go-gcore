#!/bin/bash

echo "==> Running 'gofmt' ..."
gofmt -l `find . -name '*.go' | grep -v vendor`
