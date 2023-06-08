#!/bin/bash
go build -o pharos && ./pharos migrate  create $1
