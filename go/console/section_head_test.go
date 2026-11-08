package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestSectionHeadRendersTheExample(t *testing.T) {
	head := h.Div(h.Class("cui-console-page max-w-4xl"),
		console.ConsoleSectionHead{
			Title: g.Group{h.H1(h.Class("cui-console-h2"), g.Text("Issues")), h.P(g.Text("New and regressed issues in this project."))},
			Class: "",
			Children: console.ConsoleTabs{AriaLabel: "Time window", Class: "", Tabs: []console.ConsoleTab{
				{Label: "24h", Href: "#24h", Current: false},
				{Label: "7d", Href: "#7d", Current: true},
				{Label: "30d", Href: "#30d", Current: false},
			}},
		},
		h.P(g.Text("The page’s content.")),
	)
	example.Check(t, head, "console/console-section-head/examples/section-head.html")
}

func TestSectionHeadPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleSectionHead {
		return console.ConsoleSectionHead{Title: g.Text("x"), Class: "", Children: nil}
	}, []example.Break[console.ConsoleSectionHead]{
		{Want: "console.ConsoleSectionHead needs a Title", Change: func(p *console.ConsoleSectionHead) { p.Title = nil }},
	})
}
