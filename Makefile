test: build-with-coverage
	@rm -fr .coverdata
	@mkdir -p .coverdata
	@go test ./...
	@go tool covdata percent -i=.coverdata

check-coverage: test
	@go tool covdata textfmt -i=.coverdata -o profile.txt
	@go tool cover -html=profile.txt

build:
	@go build -o echo-args

build-with-coverage:
	@go build -cover -o echo-args-coverage

.DEFAULT_GOAL := build
