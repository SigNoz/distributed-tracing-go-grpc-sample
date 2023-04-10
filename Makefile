# Makefile for building the golang package
REPONAME ?= signoz
IMAGE_NAME ?= golang-grpc-distributed-tracing
SERVER_DOCKER_TAG ?= server
CLIENT_DOCKER_TAG ?= client

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

LD_FLAGS ?=

.PHONY: dependencies
dependencies:
	go mod tidy
	go mod download

.PHONY: build-binaries
build-binaries: dependencies
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$(GOARCH)/client ./client
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$(GOARCH)/server ./server

.PHONY: docker-server
docker-server:
	docker build -t $(IMAGE_NAME):$(SERVER_DOCKER_TAG) -f ./server/Dockerfile .

.PHONY: docker-client
docker-client:
	docker build -t $(IMAGE_NAME):$(CLIENT_DOCKER_TAG) -f ./client/Dockerfile .

.PHONY: docker-server-push
docker-server-push:
	docker buildx build --platform linux/arm64,linux/amd64 -f ./server/Dockerfile --push  -t $(REPONAME)/$(IMAGE_NAME):$(SERVER_DOCKER_TAG) .

.PHONY: docker-client-push
docker-client-push:
	docker buildx build --platform linux/arm64,linux/amd64 -f ./client/Dockerfile --push  -t $(REPONAME)/$(IMAGE_NAME):$(CLIENT_DOCKER_TAG) .
