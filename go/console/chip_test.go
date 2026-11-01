package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestChipRendersTheExample(t *testing.T) {
	chips := h.P(h.Class("flex flex-wrap gap-2"),
		console.ConsoleChip{Key: "level", Value: "error", Href: "", RemoveHref: "#all", RemoveLabel: "Remove level: error", Class: ""},
		console.ConsoleChip{Key: "release", Value: "2.4.1", Href: "", RemoveHref: "#all", RemoveLabel: "Remove release: 2.4.1", Class: ""},
		console.ConsoleChip{Key: "os", Value: "iOS 18.1", Href: "#ios", RemoveHref: "", RemoveLabel: "", Class: ""},
		console.ConsoleChip{Key: "device", Value: "iPhone 16", Href: "#device", RemoveHref: "", RemoveLabel: "", Class: ""},
	)
	example.Check(t, chips, "console/console-chip/examples/chips.html")
}

func TestChipALinkWithARemoveControlPanics(t *testing.T) {
	mustPanic(t, "is a link or has a remove control, not both", func() {
		example.Render(console.ConsoleChip{Key: "env", Value: "prod", Href: "/a", RemoveHref: "/b", RemoveLabel: "Remove", Class: ""})
	})
}

func TestChipARemoveControlWithoutALabelPanics(t *testing.T) {
	mustPanic(t, "needs a RemoveLabel with a RemoveHref", func() {
		example.Render(console.ConsoleChip{Key: "env", Value: "prod", Href: "", RemoveHref: "/b", RemoveLabel: "", Class: ""})
	})
}
