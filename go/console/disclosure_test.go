package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestDisclosureRendersTheExamples(t *testing.T) {
	lines := h.Div(h.Class("cui-console-page grid gap-3"),
		console.ConsoleDisclosure{Attrs: nil, Summary: "12 system frames", Children: h.P(g.Text("libsystem_kernel.dylib __pthread_kill + 8"))},
		console.ConsoleDisclosure{Attrs: nil, Summary: "Source", Children: h.P(g.Text("let total = items.reduce(0, +)"))},
	)
	example.Check(t, lines, "console/console-disclosure/examples/disclosure.html")
}

func TestDisclosurePanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleDisclosure {
		return console.ConsoleDisclosure{Attrs: nil, Summary: "x", Children: nil}
	}, []example.Break[console.ConsoleDisclosure]{
		{Want: "console.ConsoleDisclosure needs a Summary", Change: func(p *console.ConsoleDisclosure) { p.Summary = "" }},
	})
}
