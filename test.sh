#!/bin/bash

# Run tests
for testingpkg in $(go list ./tests/.../); do
  go test -v $testingpkg
done












