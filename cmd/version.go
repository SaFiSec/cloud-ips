package cmd

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/ryanuber/columnize"
)

type cmdVersion struct {
	APICompatibility int
	BuildVersion string
	BuildDate    string
}

func (cmd *cmdVersion) run(c *kingpin.ParseContext) error {
	fmt.Println(renderVersionOutput(cmd))
	return nil
}

// Version declares the "version" sub command.
func Version(app *kingpin.Application, buildVersion, buildDate string, apiCompatibility int) {
	c := new(cmdVersion)
	c.BuildVersion = buildVersion
	c.BuildDate = buildDate
	c.APICompatibility = apiCompatibility

	app.Command("version", fmt.Sprintf("Prints %s version", app.Name)).Action(c.run)
}

// RenderVersionOutput is responsible for producing the rendered version info string.
func renderVersionOutput(cmd *cmdVersion) string {
	output := []string{
		fmt.Sprintf("Version | %s", cmd.BuildVersion),
		fmt.Sprintf("Date | %s", cmd.BuildDate),
		fmt.Sprintf("API | v%d", cmd.APICompatibility),
	}
	return columnize.SimpleFormat(output)
}
