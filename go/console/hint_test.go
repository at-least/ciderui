package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestHintRendersTheExample(t *testing.T) {
	hint := h.P(h.Class("cui-console-card-title"), g.Text("Trend (7d)"),
		console.ConsoleHint{ID: "hint-trend", AriaLabel: "What Trend (7d) means", Hint: "Events per day over the last 7 days, whatever the window above."},
	)
	example.Check(t, hint, "console/console-hint/examples/hint.html")
}

func TestHintPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleHint {
		return console.ConsoleHint{ID: "x", AriaLabel: "x", Hint: "x"}
	}, []example.Break[console.ConsoleHint]{
		{Want: "console.ConsoleHint needs an ID", Change: func(p *console.ConsoleHint) { p.ID = "" }},
		{Want: "console.ConsoleHint needs an AriaLabel", Change: func(p *console.ConsoleHint) { p.AriaLabel = "" }},
		{Want: "console.ConsoleHint needs a Hint", Change: func(p *console.ConsoleHint) { p.Hint = "" }},
	})
}
