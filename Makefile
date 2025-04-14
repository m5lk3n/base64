## help: print this help message
.PHONY: help
help:
	@echo 'usage: make <target>'
	@echo
	@echo '  where <target> is one of the following:'
	@echo
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## clean: remove all generated files
.PHONY: clean
clean:
	rm -f go.mod
	rm -f base64

## init: initialize the module
.PHONY: init
init:
	go mod init lttl.dev/base64
	go mod tidy

## get: get all dependencies
.PHONY: get
get: 
	go get

## build: build the binary for the current OS/ARCH
.PHONY: build
build:
	go fmt
	go vet
	go build

## all: make everything from scratch
.PHONY: all
all: clean init get build

## docker-build: build the Docker image
.PHONY: docker-build
docker-build:
	go fmt
	go vet
	env GOOS=linux GOARCH=amd64 go build
	docker build -t lttl.dev/base64:0.1.0 .