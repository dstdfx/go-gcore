#!/bin/bash

echo "==> Running 'goimports' ..."
goimports -l -d `find . -name '*.go' | grep -v vendor`
