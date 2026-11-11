package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func sectionNodes() g.Node {
	return h.Div(h.Class("cui-console-page max-w-4xl"),
		console.ConsoleSection{
			Title:    "Alerts",
			Detail:   "Who hears about a new issue, and how.",
			Controls: console.ConsoleBtn{Href: "", Label: "Add Alert", AriaLabel: "", Icon: lucide.Icon{}, Submit: false, Variant: console.ConsoleBtnVariantGray, Size: console.ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: ""},
			Children: g.Group{h.P(g.Text("A section’s first block.")), h.P(g.Text("Its second."))},
		},
	)
}

func TestSectionRendersTheExamples(t *testing.T) {
	example.Check(t, sectionNodes(), "console/console-section/examples/section.html")
}

func TestSectionPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleSection {
		return console.ConsoleSection{Title: "x", Detail: "", Controls: nil, Children: nil}
	}, []example.Break[console.ConsoleSection]{
		{Want: "console.ConsoleSection needs a Title", Change: func(p *console.ConsoleSection) { p.Title = "" }},
	})
}
