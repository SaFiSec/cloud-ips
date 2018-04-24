package main

import (
	"os"

	"github.com/previousnext/cloudfront-ip-sync-openvpn/cmd"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("cloudfront-ip-sync-openvpn", "Bootstrap a go utility")

	cmd.Version(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
