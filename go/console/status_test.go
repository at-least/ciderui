package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestStatusRendersTheExamples(t *testing.T) {
	statuses := h.Div(h.Class("cui-console-page grid gap-3"),
		console.ConsoleStatus{Label: "Muted", Color: console.ConsoleColorNeutral, Detail: ""},
		console.ConsoleStatus{Label: "Processing", Color: console.ConsoleColorInfo, Detail: ""},
		console.ConsoleStatus{Label: "Delivered", Color: console.ConsoleColorSuccess, Detail: ""},
		console.ConsoleStatus{Label: "No response", Color: console.ConsoleColorWarning, Detail: ""},
		console.ConsoleStatus{Label: "Failed", Color: console.ConsoleColorError, Detail: ""},
	)
	example.Check(t, statuses, "console/console-status/examples/status.html")

	detail := h.Div(h.Class("cui-console-page"),
		console.ConsoleStatus{Label: "Delivered", Color: console.ConsoleColorSuccess, Detail: "200 · 410 ms"},
	)
	example.Check(t, detail, "console/console-status/examples/detail.html")

}

func TestStatusPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleStatus {
		return console.ConsoleStatus{Label: "x", Color: console.ConsoleColorNeutral, Detail: ""}
	}, []example.Break[console.ConsoleStatus]{
		{Want: "console.ConsoleStatus needs a Label", Change: func(p *console.ConsoleStatus) { p.Label = "" }},
	})
}
