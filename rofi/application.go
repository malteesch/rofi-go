package rofi

import (
	"fmt"
	"os"
)

type Application struct {
	Prompt   string
	Message  string
	Commands []*Command
}

func (a *Application) Launch() {
	w := os.Stdout
	if len(os.Args) == 1 {
		_, err := a.WriteTo(w)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
	if len(os.Args) == 2 {
		command := findCommand(os.Args[1], a)
		if command != nil {
			command.run()
		}
	}
}

func findCommand(n string, c *Application) *Command {
	var foundCommand *Command
	for _, cmd := range c.Commands {
		if cmd.Name == n {
			return cmd
		}
		if cmd.Application != nil {
			foundCommand = findCommand(n, cmd.Application)
			if foundCommand == nil {
				continue
			} else {
				return foundCommand
			}
		}
	}
	return nil
}
