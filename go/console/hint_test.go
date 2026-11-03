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

func TestHintAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "needs an ID", func() {
		example.Render(console.ConsoleHint{ID: "", AriaLabel: "x", Hint: "x"})
	})
}

func TestHintAnEmptyAriaLabelPanics(t *testing.T) {
	mustPanic(t, "ConsoleHint needs an AriaLabel", func() {
		example.Render(console.ConsoleHint{ID: "x", AriaLabel: "", Hint: "x"})
	})
}

func TestHintAnEmptyHintPanics(t *testing.T) {
	mustPanic(t, "ConsoleHint needs a Hint", func() {
		example.Render(console.ConsoleHint{ID: "x", AriaLabel: "x", Hint: ""})
	})
}
