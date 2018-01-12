package cmd

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

type cmdVersion struct {
	BuildVersion string
	BuildDate    string
}

func (cmd *cmdVersion) run(c *kingpin.ParseContext) error {
	fmt.Println(fmt.Sprintf("Version: %s", cmd.BuildVersion))
	fmt.Println(fmt.Sprintf("Build Date: %s", cmd.BuildDate))
	return nil
}

// Version declares the "version" sub command.
func Version(app *kingpin.Application, buildVersion, buildDate string) {
	c := new(cmdVersion)
	c.BuildVersion = buildVersion
	c.BuildDate = buildDate

	app.Command("version", fmt.Sprintf("Prints %s version", app.Name)).Action(c.run)
}
