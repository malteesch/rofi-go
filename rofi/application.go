package rofi

import (
	"context"
	"fmt"
	"io"
	"os"
)

const Data = "rofiData"

type Application struct {
	// Value of the prompt (text before the user input)
	Prompt string
	// A message to be displayed below the input bar
	Message string
	// The commands (entries) that are available in this application
	Commands []*Command
	// Pass any data between invocations
	//
	// Primitive types are converted to strings.
	// Structs get marshalled as json.
	// Unmarshalling is up to the user.
	Data any
}

func (a *Application) Launch(w io.Writer, args []string) {
	if len(args) == 0 {
		_, err := a.WriteTo(w)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
	ctx := context.Background()
	rofiData := os.Getenv("ROFI_DATA")
	if rofiData != "" {
		ctx = context.WithValue(ctx, Data, rofiData)
	}
	if len(args) == 1 {
		command := findCommand(args[0], a)
		if command != nil {
			command.run(ctx, w)
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
	return foundCommand
}
