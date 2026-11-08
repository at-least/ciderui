package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestSectionHeaderRendersTheExample(t *testing.T) {
	header := h.Div(h.Class("cui-console-page max-w-4xl"),
		console.ConsoleSectionHeader{
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
	example.Check(t, header, "console/console-section-header/examples/section-header.html")
}

func TestSectionHeaderPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleSectionHeader {
		return console.ConsoleSectionHeader{Title: g.Text("x"), Class: "", Children: nil}
	}, []example.Break[console.ConsoleSectionHeader]{
		{Want: "console.ConsoleSectionHeader needs a Title", Change: func(p *console.ConsoleSectionHeader) { p.Title = nil }},
	})
}
