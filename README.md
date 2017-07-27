PNX Project: Go Edition
=======================

*TODO: CircleCI status badge goes here*

**Maintainer**: TBD

*TODO: Write a brief description about what this project does*

## Resources

* Resource 1
* Reouurce 2
* Resource 3

## Development

### Principles

* Code lives in the "workspace" directory

### Tools

* **Dependency management** - https://getgb.io
* **Build** - https://github.com/mitchellh/gox
* **Linting** - https://github.com/golang/lint

### Workflow

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
