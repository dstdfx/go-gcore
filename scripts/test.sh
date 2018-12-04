#!/bin/bash

echo "==> Running 'gotest' ..."
for testingpkg in $(go list ./tests/.../); do
  go test -count=1 -v $testingpkg
done
