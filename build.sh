#!/usr/bin/env bash
# build.sh
set -e
mkdir -p build
go build -o build/interpreter ./cmd/interpreter
