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

func TestChipALinkWithARemoveControlPanics(t *testing.T) {
	mustPanic(t, "is a link or has a remove control, not both", func() {
		example.Render(console.ConsoleChip{Key: "env", Value: "prod", Href: "/a", Remove: &console.ConsoleChipRemove{Href: "/b", AriaLabel: "Remove"}, Class: ""})
	})
}

func TestChipARemoveControlWithoutALabelPanics(t *testing.T) {
	mustPanic(t, "ConsoleChipRemove needs an Href and an AriaLabel", func() {
		example.Render(console.ConsoleChip{Key: "env", Value: "prod", Href: "", Remove: &console.ConsoleChipRemove{Href: "/b", AriaLabel: ""}, Class: ""})
	})
	mustPanic(t, "ConsoleChipRemove needs an Href and an AriaLabel", func() {
		example.Render(console.ConsoleChip{Key: "env", Value: "prod", Href: "", Remove: &console.ConsoleChipRemove{Href: "", AriaLabel: "Remove"}, Class: ""})
	})
}

func TestChipAnEmptyKeyPanics(t *testing.T) {
	mustPanic(t, "ConsoleChip needs a Key", func() {
		example.Render(console.ConsoleChip{Key: "", Value: "prod", Href: "", Remove: nil, Class: ""})
	})
}
