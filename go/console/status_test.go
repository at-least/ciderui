package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestStatusRendersTheExamples(t *testing.T) {
	statuses := h.Div(h.Class("cui-console-page grid gap-3"),
		console.ConsoleStatus{Label: "Muted", Color: console.ConsoleColorNeutral, Detail: "", Class: ""},
		console.ConsoleStatus{Label: "Processing", Color: console.ConsoleColorInfo, Detail: "", Class: ""},
		console.ConsoleStatus{Label: "Delivered", Color: console.ConsoleColorSuccess, Detail: "", Class: ""},
		console.ConsoleStatus{Label: "No response", Color: console.ConsoleColorWarning, Detail: "", Class: ""},
		console.ConsoleStatus{Label: "Failed", Color: console.ConsoleColorError, Detail: "", Class: ""},
	)
	example.Check(t, statuses, "console/console-status/examples/status.html")

	detail := h.Div(h.Class("cui-console-page"),
		console.ConsoleStatus{Label: "Delivered", Color: console.ConsoleColorSuccess, Detail: "200 · 410 ms", Class: ""},
	)
	example.Check(t, detail, "console/console-status/examples/detail.html")

}

func TestStatusPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleStatus {
		return console.ConsoleStatus{Label: "x", Color: console.ConsoleColorNeutral, Detail: "", Class: ""}
	}, []example.Break[console.ConsoleStatus]{
		{Want: "console.ConsoleStatus needs a Label", Change: func(p *console.ConsoleStatus) { p.Label = "" }},
	})
}
