#!/bin/bash
clc -s -e gset_test.go
go mod tidy
go fmt .
staticcheck .
go vet .
golangci-lint run
git st
