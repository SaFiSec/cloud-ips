package go_version

import (
	"io"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"github.com/pkg/errors"
)

type VersionConstants struct {
	BuildVersion string
	BuildDate    string
}

// GenerateVersionFile renders a version.go file.
func GenerateVersionFile(out io.Writer) error {
	// Execute git describe to insert the version identifier.
	version, err := getVersionFromGit()
	if err != nil {
		return errors.Wrap(err, "Encountered error running git describe")
	}

	// Initialise template from version.go.tmpl in this package.
	tmpl, err := template.New("versiongo").ParseFiles("version.go.tmpl")
	if err != nil {
		return errors.Wrap(err, "Error parsing template")
	}

	// Pass writer, template and variables to renderer.
	params := VersionConstants{
		BuildDate:    time.Now().Format(time.RFC3339),
		BuildVersion: version,
	}
	err = RenderVersionFile(out, tmpl, params)
	if err != nil {
		return errors.Wrap(err, "Error rendering template")
	}

	return nil
}

// RenderVersionFile renders a template with passed parameters into a writer.
func RenderVersionFile(out io.Writer, tmpl *template.Template, params VersionConstants) error {
	err := tmpl.Execute(out, params)
	return err
}

// getVersionFromGit executes git describe and returns the value.
func getVersionFromGit() (string, error) {
	out, err := exec.Command("git", "describe", "--tags", "--always").Output()
	if err != nil {
		return "", errors.Wrap(err, "Unable to determine version from git")
	}

	return strings.TrimSpace(string(out)), nil
}
