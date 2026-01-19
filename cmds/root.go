package cmds

import "github.com/urfave/cli/v3"

func Commands() []*cli.Command {
	commands := []*cli.Command{
		AddCmd(),
	}
	return commands
}
