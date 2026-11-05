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
		console.ConsoleStatus{Label: "Muted", Color: console.ConsoleColorNeutral, Class: ""},
		console.ConsoleStatus{Label: "Processing", Color: console.ConsoleColorInfo, Class: ""},
		console.ConsoleStatus{Label: "Delivered", Color: console.ConsoleColorSuccess, Class: ""},
		console.ConsoleStatus{Label: "No response", Color: console.ConsoleColorWarning, Class: ""},
		console.ConsoleStatus{Label: "Failed", Color: console.ConsoleColorError, Class: ""},
	)
	example.Check(t, statuses, "console/console-status/examples/status.html")

	detail := h.Div(h.Class("cui-console-page"),
		h.Div(
			console.ConsoleStatus{Label: "Delivered", Color: console.ConsoleColorSuccess, Class: ""},
			h.Span(h.Class("cui-console-status-detail"), g.Text("200 · 410 ms")),
		),
	)
	example.Check(t, detail, "console/console-status/examples/detail.html")

}

func TestStatusAnEmptyLabelPanics(t *testing.T) {
	mustPanic(t, "needs a Label", func() {
		example.Render(console.ConsoleStatus{Label: "", Color: console.ConsoleColorNeutral, Class: ""})
	})
}
