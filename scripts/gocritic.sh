#!/bin/bash

echo "==> Running 'gocritic' ..."
gocritic check-project --enable=all -withExperimental .
