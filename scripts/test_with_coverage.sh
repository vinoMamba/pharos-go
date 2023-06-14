#!/bin/bash

echo "testing..."

export GIN_MODE=release

go test ./test/**/*_test.go -v -coverprofile=coverage.out

echo "test done"
echo "generating coverage.html..."
go tool cover -html=coverage.out -o coverage.html
echo "the coverage.html is generated"
