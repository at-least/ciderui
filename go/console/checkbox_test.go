package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestCheckboxRendersTheExample(t *testing.T) {
	checkboxes := h.Div(h.Class("grid gap-3"),
		console.ConsoleCheckbox{Label: "Mute this issue", AriaLabel: "", Control: input(h.Name("mute"))},
		console.ConsoleCheckbox{Label: "Notify me when it comes back", AriaLabel: "", Control: input(h.Name("notify"), h.Checked())},
		console.ConsoleCheckbox{Label: "Share with every project", AriaLabel: "", Control: input(h.Name("share"), h.Disabled())},
		console.ConsoleCheckbox{Label: "", AriaLabel: "Select all", Control: input(h.Name("all"))},
	)
	example.Check(t, checkboxes, "console/console-checkbox/examples/checkbox.html")
}

func TestCheckboxPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleCheckbox {
		return console.ConsoleCheckbox{Label: "x", AriaLabel: "", Control: input()}
	}, []example.Break[console.ConsoleCheckbox]{
		{Want: "console.ConsoleCheckbox needs a Label or an AriaLabel, not both", Change: func(p *console.ConsoleCheckbox) { p.Label = "" }},
		{Want: "console.ConsoleCheckbox needs a Label or an AriaLabel, not both", Change: func(p *console.ConsoleCheckbox) { p.AriaLabel = "y" }},
		{Want: "console.ConsoleCheckbox needs a Control", Change: func(p *console.ConsoleCheckbox) { p.Control = nil }},
	})
}
