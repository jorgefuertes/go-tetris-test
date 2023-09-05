SHELL = /usr/bin/env bash

deps:
	go mod tidy

run: deps
	go run cmd/tetris.go
