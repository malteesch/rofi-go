package rofi

import (
	"context"
	"fmt"
	"io"
	"os"
)

const writerContextKey = "writer"

type Application struct {
	Prompt      string
	subCommands []*Command
}

func (a *Application) AddCommand(c *Command) {
	a.subCommands = append(a.subCommands, c)
}

func (a *Application) Launch(w io.Writer) {
	if len(os.Args) == 1 {
		_, err := a.WriteTo(w)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
	ctx := context.Background()
	if len(os.Args) == 2 {
		command := findCommand(os.Args[1], a)
		if command != nil {
			command.Run(context.WithValue(ctx, writerContextKey, w))
		}
	}
}

func (a *Application) commands() []*Command {
	return a.subCommands
}

func findCommand(n string, c commandContainer) *Command {
	var foundCommand *Command
	for _, cmd := range c.commands() {
		if cmd.Name == n {
			return cmd
		}
		if len(cmd.subCommands) > 0 {
			foundCommand = findCommand(n, cmd)
			if foundCommand == nil {
				continue
			} else {
				return foundCommand
			}
		}
	}
	return nil
}
