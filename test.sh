#!/bin/bash

go test -v -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out
