Gopher
======

[![CircleCI](https://circleci.com/gh/previousnext/CHANGE_ME.svg?style=svg)](https://circleci.com/gh/previousnext/CHANGE_ME)

![Logo](/logo/small.png "Logo")

**Maintainer**: Gopher

This is a brief description on what the project does.

## Development

### Getting started

To work on this project you will first need Go installed on your machine.

#### Manual Setup

First make sure Go is properly installed and that a GOPATH has been set. You will also need to add $GOPATH/bin to your $PATH. For steps on getting started with Go: https://golang.org/doc/install

Next, using Git, clone this repository into $GOPATH/src/github.com/previousnext/CHANGE_ME. All the necessary dependencies are either vendored or automatically installed, so you just need to type `make test`. This will run the tests and compile the binary. If this exits with exit status 0, then everything is working!

```bash
$ cd "$GOPATH/src/github.com/previousnext/CHANGE_ME"
$ make test
```

To compile a development version of CHANGE_ME, run `make build`. This will build everything using gox and put binaries in the bin and $GOPATH/bin folders:

```bash
$ make build
...

# Linux:
$ bin/CHANGE_ME_linux_amd64 --help

# OSX:
$ bin/CHANGE_ME_darwin_amd64 --help
```

#### Easy Setup

Alternatively, you can use the [Docker Compose](docker-compose.yml) stack in the root of this repo to stand up a container with the appropriate dev tooling already set up for you.

Using Git, clone this repo on your local machine. Run the test suite to ensure the tooling works.

```bash
$ docker-compose run --rm dev make test
```

To compile a development version of CHANGE_ME, run `make build`. This will build everything using gox and put binaries in the bin and $GOPATH/bin folders:

```bash
$ docker-compose run --rm dev make build

...

$ docker-compose run --rm dev bin/CHANGE_ME_linux_amd64 --help
```


### Documentation

See `/docs`

### Resources

* [Dave Cheney - Reproducible Builds](https://www.youtube.com/watch?v=c3dW80eO88I)
* [Bryan Cantril - Debugging under fire](https://www.youtube.com/watch?v=30jNsCVLpAE&t=2675s)
* [Sam Boyer - The New Era of Go Package Management](https://www.youtube.com/watch?v=5LtMb090AZI)
* [Kelsey Hightower - From development to production](https://www.youtube.com/watch?v=XL9CQobFB8I&t=787s)

### Tools

```bash
# Dependency management
go get -u github.com/golang/dep/cmd/dep

# Testing
go get -u github.com/golang/lint/golint

# Release management.
go get -u github.com/tcnksm/ghr

# Build
go get -u github.com/mitchellh/gox
```

### Workflow

**Testing**

```bash
make lint test
```

**Building**

```bash
make build
```

**Releasing**

```bash
make release
```