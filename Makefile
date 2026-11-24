CONTAINER_ENGINE ?= podman
IMAGE ?= localhost/hello-world:local

.PHONY: build test bundle

build:
	go build -o bin/hello .

test:
	go test ./...

bundle:
	CGO_ENABLED=0 GOOS=linux go build -trimpath -o bin/linux/hello .
	$(CONTAINER_ENGINE) build -f Containerfile -t $(IMAGE) bin/linux
