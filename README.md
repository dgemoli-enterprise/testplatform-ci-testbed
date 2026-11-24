# Hello world

A Go CLI that prints `hello world` and exits.

## Build and test

With Go 1.26 or newer and Make installed:

```sh
make test
make build
./bin/hello
```

## Run in a container

`make bundle` builds a static Linux binary and packages it in a minimal image
using `Containerfile`.

```sh
make bundle
podman run --rm localhost/hello-world:local
```

To use Docker, run `make bundle CONTAINER_ENGINE=docker`, then use `docker run`
with the same image name.
