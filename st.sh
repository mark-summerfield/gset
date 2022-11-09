#!/bin/bash
clc -s -e gset_1_test.go  gset_2_test.go
go mod tidy
go fmt .
staticcheck .
go vet .
golangci-lint run
git st
