package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestStatusRendersTheExamples(t *testing.T) {
	statuses := h.Div(h.Class("cui-console-page grid gap-3"),
		console.ConsoleStatus{Label: "Muted", Color: console.ConsoleColorNeutral, Series: "", Class: ""},
		console.ConsoleStatus{Label: "New", Color: console.ConsoleColorBrand, Series: "", Class: ""},
		console.ConsoleStatus{Label: "Processing", Color: console.ConsoleColorInfo, Series: "", Class: ""},
		console.ConsoleStatus{Label: "Delivered", Color: console.ConsoleColorSuccess, Series: "", Class: ""},
		console.ConsoleStatus{Label: "No response", Color: console.ConsoleColorWarning, Series: "", Class: ""},
		console.ConsoleStatus{Label: "Failed", Color: console.ConsoleColorError, Series: "", Class: ""},
	)
	example.Check(t, statuses, "console/console-status/examples/status.html")

	detail := h.Div(h.Class("cui-console-page"),
		h.Div(
			console.ConsoleStatus{Label: "Delivered", Color: console.ConsoleColorSuccess, Series: "", Class: ""},
			h.Span(h.Class("cui-console-status-detail"), g.Text("200 · 410 ms")),
		),
	)
	example.Check(t, detail, "console/console-status/examples/detail.html")

	key := h.Div(h.Class("cui-console-page flex flex-wrap gap-4"),
		console.ConsoleStatus{Label: "iOS", Color: console.ConsoleColorNeutral, Series: "oklch(54.2% 0.162 251.9)", Class: ""},
		console.ConsoleStatus{Label: "Android", Color: console.ConsoleColorNeutral, Series: "oklch(66.6% 0.179 58.3)", Class: ""},
	)
	example.Check(t, key, "console/console-status/examples/key.html")
}

func TestStatusAnEmptyLabelPanics(t *testing.T) {
	mustPanic(t, "needs a Label", func() {
		example.Render(console.ConsoleStatus{Label: "", Color: console.ConsoleColorNeutral, Series: "", Class: ""})
	})
}

func TestStatusAColorAndASeriesPanic(t *testing.T) {
	mustPanic(t, "a color or a series", func() {
		example.Render(console.ConsoleStatus{Label: "iOS", Color: console.ConsoleColorBrand, Series: "red", Class: ""})
	})
}
