package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func sectionNodes() g.Node {
	return h.Div(h.Class("cui-console-page max-w-4xl"),
		console.ConsoleSection{
			Title:    "Alerts",
			Detail:   "Who hears about a new issue, and how.",
			Children: g.Group{h.P(g.Text("A section’s first block.")), h.P(g.Text("Its second."))},
		},
	)
}

func TestSectionRendersTheExamples(t *testing.T) {
	example.Check(t, sectionNodes(), "console/console-section/examples/section.html")
}

func TestSectionPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleSection {
		return console.ConsoleSection{Title: "x", Detail: "", Children: nil}
	}, []example.Break[console.ConsoleSection]{
		{Want: "console.ConsoleSection needs a Title", Change: func(p *console.ConsoleSection) { p.Title = "" }},
	})
}
