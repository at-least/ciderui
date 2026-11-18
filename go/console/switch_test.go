package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestSwitchRendersTheExample(t *testing.T) {
	switches := h.Div(h.Class("grid gap-3"),
		console.ConsoleSwitch{Label: "Only what needs you", Control: input(h.Name("todo"))},
		console.ConsoleSwitch{Label: "Show live events", Control: input(h.Name("live"), h.Checked())},
		console.ConsoleSwitch{Label: "Share with every project", Control: input(h.Name("share"), h.Disabled())},
	)
	example.Check(t, switches, "console/console-switch/examples/switch.html")
}

func TestSwitchPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleSwitch {
		return console.ConsoleSwitch{Label: "x", Control: input()}
	}, []example.Break[console.ConsoleSwitch]{
		{Want: "console.ConsoleSwitch needs a Label", Change: func(p *console.ConsoleSwitch) { p.Label = "" }},
		{Want: "console.ConsoleSwitch needs a Control", Change: func(p *console.ConsoleSwitch) { p.Control = nil }},
	})
}
