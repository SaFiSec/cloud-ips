PNX Project: Go Edition
=======================

[![CircleCI](https://circleci.com/gh/previousnext/CHANGE_ME.svg?style=svg)](https://circleci.com/gh/previousnext/CHANGE_ME)

**Maintainer**: TBD

This is a brief description on what the project does............ maybe include some diagrams.

## Resources

* Resource 1
* Reouurce 2
* Resource 3

## Development

### Principles

* Code lives in the `workspace` directory

### Tools

* **Dependency management** - https://getgb.io
* **Build** - https://github.com/mitchellh/gox
* **Linting** - https://github.com/golang/lint

### Workflow

(While in the `workspace` directory)

**Installing a new dependency**

```bash
gb vendor fetch github.com/foo/bar
```

**Running quality checks**

```bash
make lint test
```

**Building binaries**

```bash
make build
```
