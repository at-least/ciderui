package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestBadgeRendersTheExample(t *testing.T) {
	badges := g.Group{
		console.ConsoleBadge{Label: "Unhandled", Color: console.ConsoleColorNeutral, Class: ""},
		console.ConsoleBadge{Label: "Info", Color: console.ConsoleColorInfo, Class: ""},
		console.ConsoleBadge{Label: "Resolved", Color: console.ConsoleColorSuccess, Class: ""},
		console.ConsoleBadge{Label: "Warning", Color: console.ConsoleColorWarning, Class: ""},
		console.ConsoleBadge{Label: "Fatal", Color: console.ConsoleColorError, Class: ""},
	}
	example.Check(t, h.P(h.Class("flex flex-wrap gap-2"), badges), "console/console-badge/examples/badges.html")
}

func TestBadgeAnEmptyLabelPanics(t *testing.T) {
	mustPanic(t, "needs a Label", func() {
		example.Render(console.ConsoleBadge{Label: "", Color: console.ConsoleColorNeutral, Class: ""})
	})
}
