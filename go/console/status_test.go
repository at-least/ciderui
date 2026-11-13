package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestStatusRendersTheExamples(t *testing.T) {
	status := h.Div(h.Class("cui-console-page"),
		console.ConsoleStatus{Label: "No response", Detail: ""},
	)
	example.Check(t, status, "console/console-status/examples/status.html")

	detail := h.Div(h.Class("cui-console-page"),
		console.ConsoleStatus{Label: "No response", Detail: "Timed out · 10 s"},
	)
	example.Check(t, detail, "console/console-status/examples/detail.html")
}

func TestStatusPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleStatus {
		return console.ConsoleStatus{Label: "x", Detail: ""}
	}, []example.Break[console.ConsoleStatus]{
		{Want: "console.ConsoleStatus needs a Label", Change: func(p *console.ConsoleStatus) { p.Label = "" }},
	})
}
