clean:
	rm -rf app.bin

tests-unit:
	go test -race ./internal/...
	go test -coverprofile=coverage.out ./internal/pkg/...

tests-functional:
	go test -tags functional -race ./internal/cmd/...

lint:
	golangci-lint run

code-coverage:
	go tool cover -func=coverage.out

GONTAINER_BINARY ?= app.bin

BUILT_BY ?= make${MAKE_VERSION}

build: export DATETIME = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
build: export GITHASH = $(shell git rev-parse HEAD)
build: export VERSION = dev-$(shell git rev-parse --abbrev-ref HEAD)
build: export IS_GIT_DIRTY = $(shell git diff --quiet && echo 'false' || echo 'true')
build: clean
	go build -v -ldflags="-X 'main.date=${DATETIME}' -X 'main.commit=${GITHASH}' -X 'main.version=${VERSION}' -X 'main.isGitDirty=${IS_GIT_DIRTY}' -X 'main.builtBy=${BUILT_BY}'" -o ${GONTAINER_BINARY} main.go

global: export GONTAINER_BINARY = /usr/local/bin/gontainer
global: build

self-compile:
	gontainer build -i internal/gontainer/gontainer.yaml -i internal/gontainer/gontainer_\*.yaml -o internal/gontainer/gontainer.go

generate-stub:
	gontainer build -i internal/gontainer/gontainer.yaml -i internal/gontainer/gontainer_\*.yaml -o internal/gontainer/stub.go --stub

tests: tests-unit tests-functional lint

.DEFAULT_GOAL := build
