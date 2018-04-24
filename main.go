package main

import (
	"os"

	"github.com/previousnext/cloudfront-ip-sync-openvpn/cmd"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("cloudfront-ip-sync-openvpn", "Tool to render a openvpn config with cloudfront edge location IPs")

	cmd.Render(app)
	cmd.Version(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
