package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestChipRendersTheExample(t *testing.T) {
	chips := h.P(h.Class("flex flex-wrap gap-2"),
		console.ConsoleChip{Key: "level", Value: "error", Href: "", RemoveHref: "#all", RemoveAriaLabel: "Remove level: error", Class: ""},
		console.ConsoleChip{Key: "release", Value: "2.4.1", Href: "", RemoveHref: "#all", RemoveAriaLabel: "Remove release: 2.4.1", Class: ""},
		console.ConsoleChip{Key: "os", Value: "iOS 18.1", Href: "#ios", RemoveHref: "", RemoveAriaLabel: "", Class: ""},
		console.ConsoleChip{Key: "device", Value: "iPhone 16", Href: "#device", RemoveHref: "", RemoveAriaLabel: "", Class: ""},
	)
	example.Check(t, chips, "console/console-chip/examples/chips.html")
}

func TestChipALinkWithARemoveControlPanics(t *testing.T) {
	mustPanic(t, "is a link or has a remove control, not both", func() {
		example.Render(console.ConsoleChip{Key: "env", Value: "prod", Href: "/a", RemoveHref: "/b", RemoveAriaLabel: "Remove", Class: ""})
	})
}

func TestChipARemoveControlWithoutALabelPanics(t *testing.T) {
	mustPanic(t, "needs a RemoveAriaLabel with a RemoveHref", func() {
		example.Render(console.ConsoleChip{Key: "env", Value: "prod", Href: "", RemoveHref: "/b", RemoveAriaLabel: "", Class: ""})
	})
}

func TestChipAnEmptyKeyPanics(t *testing.T) {
	mustPanic(t, "ConsoleChip needs a Key", func() {
		example.Render(console.ConsoleChip{Key: "", Value: "prod", Href: "", RemoveHref: "", RemoveAriaLabel: "", Class: ""})
	})
}
