package main

import (
	"os"

	"github.com/previousnext/gopher/pkg/go-version"
)

const (
	versionFile = "version/version.go"
)

func main() {
	out, err := os.Create(versionFile)
	if err != nil {
		panic(err)
	}

	err = goversion.GenerateVersionFile(out)
	if err != nil {
		panic(err)
	}
}
