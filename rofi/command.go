package rofi

import (
	"context"
	"io"
)

type Command struct {
	Name        string
	Info        string
	Run         func(context.Context)
	Application *Application
}

func (c *Command) run(ctx context.Context, w io.Writer) {
	if c.Run != nil {
		c.Run(ctx)
	}
	if c.Application != nil {
		c.Application.WriteTo(w)
	}
}
