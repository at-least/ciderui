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
		console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantLine, Open: false, Summary: "12 system frames", SummaryClass: "", Class: "", Children: h.P(g.Text("libsystem_kernel.dylib __pthread_kill + 8"))},
		console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantLine, Open: true, Summary: "Source", SummaryClass: "", Class: "", Children: h.P(g.Text("let total = items.reduce(0, +)"))},
	)
	example.Check(t, lines, "console/console-disclosure/examples/disclosure.html")

	buttons := h.Div(h.Class("cui-console-page grid gap-3"),
		console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantButton, Open: false, Summary: "Filters", SummaryClass: "", Class: "", Children: h.P(g.Text("Level, release and the search box."))},
		console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantButton, Open: true, Summary: "Filters", SummaryClass: "", Class: "", Children: h.P(g.Text("Level, release and the search box."))},
	)
	example.Check(t, buttons, "console/console-disclosure/examples/button.html")

	plain := h.Div(h.Class("cui-console-page grid gap-3"),
		console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantPlainButton, Open: false, Summary: "Edit", SummaryClass: "", Class: "", Children: h.P(g.Text("Its name, its address and its digest."))},
		console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantPlainButton, Open: true, Summary: "Edit", SummaryClass: "", Class: "", Children: h.P(g.Text("Its name, its address and its digest."))},
	)
	example.Check(t, plain, "console/console-disclosure/examples/plain-button.html")
}

func TestDisclosureThePagesUtilitiesGoOnEachPart(t *testing.T) {
	got := example.Render(console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantLine, Open: false, Summary: "Raw", SummaryClass: "text-xs", Class: "mt-6", Children: nil})
	if !strings.HasPrefix(got, `<details class="cui-console-disclosure mt-6"><summary class="cui-console-disclosure-summary text-xs"><svg`) {
		t.Errorf("got %s", got)
	}
}

func TestDisclosureAnEmptySummaryPanics(t *testing.T) {
	mustPanic(t, "needs a Summary", func() {
		example.Render(console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantLine, Open: false, Summary: "", SummaryClass: "", Class: "", Children: nil})
	})
}
