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
package main

import (
  "log"
  "net/http"

  "github.com/alecthomas/kingpin"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
  cliPrometheusPort = kingpin.Flag("prometheus-port", "Prometheus metrics port").Default(":9000").OverrideDefaultFromEnvar("METRICS_PORT").String()
  cliPrometheusPath = kingpin.Flag("prometheus-path", "Prometheus metrics path").Default("/metrics").OverrideDefaultFromEnvar("METRICS_PATH").String()
)

func main() {
  go metrics(*cliPrometheusPort, *cliPrometheusPath)
}

// Helper function for serving Prometheus metrics.
func metrics(port, path string) {
  http.Handle(path, promhttp.Handler())
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
