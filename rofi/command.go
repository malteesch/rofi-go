package rofi

import (
	"os"
)

type Command struct {
	Name        string
	Info        string
	Run         func()
	Application *Application
}

func (c *Command) run() {
	if c.Run != nil {
		c.Run()
	}
	if c.Application != nil {
		c.Application.WriteTo(os.Stdout)
	}
}
