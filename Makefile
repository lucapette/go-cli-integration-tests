test:
	go test ./...

build:
	go build cmd/echo-args.go

.DEFAULT_GOAL := build
