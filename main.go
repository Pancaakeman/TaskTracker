package main

import (
	"context"
	"os"
	"taskTracker/cmds"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:        "help",
		Usage:       "help",
		Description: "TaskTracker is a simple task keeping application written in Go\nOffers simple features to Add, Delete, Edit and Query Tasks",
		Commands:    cmds.Commands(),
	}
	cmd.Run(context.Background(), os.Args)

}
