# rofi-go
A library to create [Rofi](https://github.com/davatorium/rofi) plugins/scripts using [Go](https://go.dev/).

> [!WARNING]
> At the moment only the basic functionality is provided.

## Concept
One Application is comprised of Commands.
Commands can have a runnable action and/or a sub-application holding other commands.
With this you can have arbitrarily nested command flows.

## TODO 
- [ ] allow all metadata to be attached
- [ ] modifying the state inside command actions
