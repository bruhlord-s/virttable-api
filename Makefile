.PHONY: init
init:
	cp configs/apiserver.toml.sample configs/apiserver.toml && cp docker-compose.yml.sample docker-compose.yml

.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL: build