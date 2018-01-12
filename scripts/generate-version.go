package main

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	versionFile = "version/version.go"
)

func main() {
	out, _ := os.Create(versionFile)
	out.Write([]byte("package version \n\nconst (\n"))

	// Execute git describe to insert the version identifier.
	version, err := execGitDescribe()
	if err != nil {
		panic(err)
	}
	writeConst(out, "BuildVersion", version)

	// Insert date too.
	writeConst(out, "BuildDate", time.Now().Format(time.RFC3339))

	out.Write([]byte(")\n"))
}

// writeConst is a helper method to render a constant.
func writeConst(FileHandler *os.File, Name, Value string) {
	FileHandler.Write([]byte("	" + Name + " = `" + Value + "`\n"))
}

// execGitDescribe executes git describe and returns the value.
func execGitDescribe() (string, error) {
	out, err := exec.Command("git", "describe", "--tags", "--always").Output()
	if err != nil {
		return "", errors.Wrap(err, "Unable to determine version from git")
	}

	return strings.TrimSpace(string(out)), nil
}
