package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestChipRendersTheExample(t *testing.T) {
	chips := h.P(h.Class("flex flex-wrap gap-2"),
		console.ConsoleChip{Key: "level", Value: "error", Href: "", Remove: &console.ConsoleChipRemove{Href: "#all", AriaLabel: "Remove level: error"}, Class: ""},
		console.ConsoleChip{Key: "release", Value: "2.4.1", Href: "", Remove: &console.ConsoleChipRemove{Href: "#all", AriaLabel: "Remove release: 2.4.1"}, Class: ""},
		console.ConsoleChip{Key: "os", Value: "iOS 18.1", Href: "#ios", Remove: nil, Class: ""},
		console.ConsoleChip{Key: "device", Value: "iPhone 16", Href: "#device", Remove: nil, Class: ""},
	)
	example.Check(t, chips, "console/console-chip/examples/chips.html")
}

func TestChipPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleChip {
		return console.ConsoleChip{Key: "env", Value: "prod", Href: "", Remove: nil, Class: ""}
	}, []example.Break[console.ConsoleChip]{
		{Want: "console.ConsoleChip needs a Key", Change: func(p *console.ConsoleChip) { p.Key = "" }},
		{Want: "console.ConsoleChip is a link or has a remove control, not both", Change: func(p *console.ConsoleChip) {
			p.Href, p.Remove = "/a", &console.ConsoleChipRemove{Href: "/b", AriaLabel: "Remove"}
		}},
		{Want: "console.ConsoleChipRemove needs an Href and an AriaLabel", Change: func(p *console.ConsoleChip) { p.Remove = &console.ConsoleChipRemove{Href: "/b", AriaLabel: ""} }},
		{Want: "console.ConsoleChipRemove needs an Href and an AriaLabel", Change: func(p *console.ConsoleChip) { p.Remove = &console.ConsoleChipRemove{Href: "", AriaLabel: "Remove"} }},
	})
}
