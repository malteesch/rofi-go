package main

import (
	"context"
	"os"

	"github.com/malteesch/rofi-go/rofi"
)

func main() {
	app := &rofi.Application{
		Prompt: "Custom prompt",
		Message: "Custom message below ",
		Commands: []*rofi.Command{
			{
				Name: "Layer1Command",
				Application: &rofi.Application{
					Prompt: "Change prompt here",
					Commands: []*rofi.Command{
						{
							Name: "Layer2Command",
							Run:  func(ctx context.Context) {
								// YOUR COMMAND ACTION HERE
							},
						},
					},
				},
			},
		},
	}
	app.Launch(os.Stdout, os.Args[1:])
}
