package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestHeadRendersTheExample(t *testing.T) {
	tabs := []console.ConsoleTab{
		{Label: "Overview", Href: "#overview", Current: false},
		{Label: "Issues", Href: "#issues", Current: true},
		{Label: "Releases", Href: "#releases", Current: false},
		{Label: "Jobs", Href: "#jobs", Current: false},
	}
	head := h.Div(h.Class("cui-console-page max-w-4xl"),
		console.ConsoleHead{Name: "Shop", Href: "#overview", TabsAriaLabel: "Project", Tabs: tabs, Class: ""},
		h.P(g.Text("The page’s content.")),
	)
	example.Check(t, head, "console/console-head/examples/head.html")
}

func TestHeadPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleHead {
		return console.ConsoleHead{Name: "x", Href: "/", TabsAriaLabel: "x", Tabs: nil, Class: ""}
	}, []example.Break[console.ConsoleHead]{
		{Want: "console.ConsoleHead needs a Name", Change: func(p *console.ConsoleHead) { p.Name = "" }},
	})
}
