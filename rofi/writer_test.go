package rofi_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/malteesch/rofi-go/rofi"
)

func TestMetadataIsWritten(t *testing.T) {
	app := &rofi.Application{
		Prompt: "TestPrompt",
		Data:   "data",
	}
	sb := strings.Builder{}
	app.Launch(&sb, []string{})

	expected := fmt.Sprintf("%c%s%c%s\n%c%s%c%s\n", 0x00, "prompt", 0x1f, "TestPrompt", 0x00, "data", 0x1f, "data")

	actual := sb.String()
	if expected != actual {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
}

func TestCommandIsWritten(t *testing.T) {
	app := &rofi.Application{
		Commands: []*rofi.Command{
			{
				Name: "Command",
			},
		},
	}
	sb := strings.Builder{}
	app.Launch(&sb, []string{})

	expected := "Command\n"

	actual := sb.String()
	if expected != actual {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
}

func TestSubApplicationIsWritten(t *testing.T) {
	app := &rofi.Application{
		Commands: []*rofi.Command{
			{
				Name: "Command",
				Application: &rofi.Application{
					Commands: []*rofi.Command{
						{
							Name: "SubCommand",
						},
					},
				},
			},
		},
	}
	sb := strings.Builder{}
	app.Launch(&sb, []string{"Command"})

	expected := "SubCommand\n"

	actual := sb.String()
	if expected != actual {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
}
