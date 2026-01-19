package cmds

import (
	"context"
	"fmt"
	"os"
	"slices"
	"taskTracker/internal"

	"github.com/urfave/cli/v3"
)

func AddCmd() *cli.Command {
	cmdAdd := &cli.Command{
		Name:        "add",
		Usage:       "Command to add a task",
		UsageText:   "add [task name] [task Desc] [task Urgency]",
		Description: "Simple Command to add a Task to the Task List",
		Action: func(_ context.Context, command *cli.Command) error {
			os.Args = slices.Delete(os.Args, 0, 2)
			fmt.Printf("%+v\n", os.Args)
			// TODO: Use Flags instead for safety
			newTask := internal.Task{
				Name:    os.Args[0],
				Desc:    os.Args[1],
				Urgency: os.Args[2],
			}
			internal.AddTask(newTask)
			internal.ListTasks()

			return nil
		},
	}
	_ = cmdAdd.Run(context.Background(), os.Args)
	return cmdAdd
}
