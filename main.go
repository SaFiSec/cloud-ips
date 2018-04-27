package main

import (
	"os"

	"github.com/previousnext/cloud-ips/cmd"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("cloud-ips", "Tool to fetch IP addresses for various cloud providers, regions and services")

	cmd.Get(app)
	cmd.Version(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
