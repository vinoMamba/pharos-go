#!/bin/bash

export GIN_MODE=release
go test ./test/**/*_test.go -v -coverprofile=coverage.out
