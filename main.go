package main

import (
	"github.com/previousnext/gopher/cmd"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

//go:generate go run scripts/generate-version.go

func main() {
	app := kingpin.New("Gopher", "Bootstrap a go utility")

	cmd.Version(app, BuildVersion, BuildDate)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
