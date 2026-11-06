package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestDisclosureRendersTheExamples(t *testing.T) {
	lines := h.Div(h.Class("cui-console-page grid gap-3"),
		console.ConsoleDisclosure{Open: false, Summary: "12 system frames", SummaryClass: "", Class: "", Children: h.P(g.Text("libsystem_kernel.dylib __pthread_kill + 8"))},
		console.ConsoleDisclosure{Open: true, Summary: "Source", SummaryClass: "", Class: "", Children: h.P(g.Text("let total = items.reduce(0, +)"))},
	)
	example.Check(t, lines, "console/console-disclosure/examples/disclosure.html")
}

func TestDisclosureThePagesUtilitiesGoOnEachPart(t *testing.T) {
	got := example.Render(console.ConsoleDisclosure{Open: false, Summary: "Raw", SummaryClass: "text-xs", Class: "mt-6", Children: nil})
	if !strings.HasPrefix(got, `<details class="cui-console-disclosure mt-6"><summary class="cui-console-disclosure-summary text-xs"><svg`) {
		t.Errorf("got %s", got)
	}
}

func TestDisclosurePanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleDisclosure {
		return console.ConsoleDisclosure{Open: false, Summary: "x", SummaryClass: "", Class: "", Children: nil}
	}, []example.Break[console.ConsoleDisclosure]{
		{Want: "console.ConsoleDisclosure needs a Summary", Change: func(p *console.ConsoleDisclosure) { p.Summary = "" }},
	})
}
