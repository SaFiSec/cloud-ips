#!/usr/bin/make -f

export CGO_ENABLED=0

PROJECT=github.com/previousnext/CHANGE_ME

# Builds the project.
build:
	gox -os='linux darwin' -arch='amd64' -output='bin/CHANGE_ME_{{.OS}}_{{.Arch}}' -ldflags='-extldflags "-static"' $(PROJECT)

# Run all lint checking with exit codes for CI.
lint:
	golint -set_exit_status `go list ./... | grep -v /vendor/`

# Run tests with coverage reporting.
test:
	go test -cover ./...

IMAGE=previousnext/CHANGE_ME
VERSION=$(shell git describe --tags --always)

# Releases the project Docker Hub.
release-docker:
	docker build -t ${IMAGE}:${VERSION} -t ${IMAGE}:latest .
	docker push ${IMAGE}:${VERSION}
	docker push ${IMAGE}:latest

release-github: build
	ghr -u previousnext "${VERSION}" ./bin/

release: release-docker release-github

.PHONY: build lint test release-docker release-github release
