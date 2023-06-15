#!/bin/bash

export GIN_MODE=release

go build -o pharos; ./pharos cover
