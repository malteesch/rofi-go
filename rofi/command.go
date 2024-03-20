package rofi

import (
	"context"
	"io"
)

type Command struct {
	Name        string
	Info        *string
	Action      func(context.Context)
	subCommands []*Command
}

type commandContainer interface {
	commands() []*Command
}

func (c *Command) AddSubCommand(sc *Command) {
	c.subCommands = append(c.subCommands, sc)
}

func (c *Command) Run(ctx context.Context) {
	if c.Action != nil {
		c.Action(ctx)
	}
	if len(c.subCommands) > 0 {
		w := ctx.Value(writerContextKey).(io.Writer)
		for _, cmd := range c.subCommands {
			cmd.WriteTo(w)
		}
	}
}

func EmptyAction(ctx context.Context) {}

func (c *Command) commands() []*Command {
	return c.subCommands
}
