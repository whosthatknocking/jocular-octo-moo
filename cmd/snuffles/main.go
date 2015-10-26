package main

import (
	"os"
	"path"

	"github.com/jawher/mow.cli"
)

var (
	// Build version
	Version string

	// Build date
	BuildDate string

	// Git build commit
	BuildCommit string
)

func main() {
	progName := path.Base(os.Args[0])

	if Version == "" {
		Version = "(development version)"
	}

	app := cli.App(progName, "Export docker info")
	defineCommands(app.Cmd)
	app.Run(os.Args)
}
