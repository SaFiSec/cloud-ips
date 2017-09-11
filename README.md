PNX Project: Go Edition
=======================

[![CircleCI](https://circleci.com/gh/previousnext/CHANGE_ME.svg?style=svg)](https://circleci.com/gh/previousnext/CHANGE_ME)

**Maintainer**: CHANGE ME

This is a brief description on what the project does.

## Resources

* [Dave Cheney - Reproducible Builds](https://www.youtube.com/watch?v=c3dW80eO88I)

## Development

### Principles

#### Structure

* Code lives in the `workspace` directory

#### Logging

Logging should use the package `github.com/prometheus/common/log`

This results in a log like:

```
INFO[0000] Serving Prometheus metrics endpoint           source="main.go:23"
```

Notice the `source="main.go:23"`, this allows us to track down the line of code.

#### Metrics

Metrics should be exposed to Prometheus on:

* Port `9000`
* Path `/metrics`

We should use the package `github.com/prometheus/client_golang/prometheus/promhttp` and implemented with the following line:

```
var cliPrometheus = kingpin.Flag("prometheus", "Prometheus metrics endpoint").Default(":9000").OverrideDefaultFromEnvar("PROMETHEUS").String()

func main() {
        go metrics(*cliPrometheus)
}

// Helper function for serving Prometheus metrics.
func metrics(port string) {
        http.Handle("/metrics", promhttp.Handler())
        log.Fatal(http.ListenAndServe(port, nil))
}
```

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
