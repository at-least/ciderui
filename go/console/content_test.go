package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func contentNodes() g.Node {
	sidebar := console.ConsoleSidebar{AriaLabel: "CrashCart", Home: "#home", Wordmark: "CrashCart", SkipLabel: "Skip to Content", OpenAriaLabel: "Menu", CloseAriaLabel: "Close", Footer: nil, Children: g.Group{
		console.ConsoleMenuItem{Href: "#projects", Label: "All Projects", Icon: lucide.Icon{}, Current: true},
		console.ConsoleMenuItem{Href: "#issues", Label: "All Issues", Icon: lucide.Icon{}, Current: false},
	}}
	return h.Div(h.Class("cui-console-page"), console.ConsoleContent{Sidebar: &sidebar, Back: console.ConsoleContentBack{Href: "", Label: ""}, Title: "All Projects", Detail: "Every project you can see, the ones that need you first.", Controls: nil, Children: nil})
}

func rhythmNodes() g.Node {
	return h.Div(h.Class("cui-console-page"), console.ConsoleContent{Sidebar: nil, Back: console.ConsoleContentBack{Href: "", Label: ""}, Title: "Findings", Detail: "", Controls: nil, Children: g.Group{
		console.ConsoleSection{Title: "Needs You", Detail: "", Children: g.Group{h.P(g.Text("A section’s first block.")), h.P(g.Text("Its second."))}},
		console.ConsoleSection{Title: "Closed", Detail: "", Children: h.P(g.Text("The next section’s block."))},
	}})
}

func TestContentRendersTheExamples(t *testing.T) {
	example.Check(t, contentNodes(), "console/console-content/examples/content.html")
	example.Check(t, rhythmNodes(), "console/console-content/examples/rhythm.html")
}

func TestContentPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleContent {
		return console.ConsoleContent{Sidebar: nil, Back: console.ConsoleContentBack{Href: "", Label: ""}, Title: "x", Detail: "", Controls: nil, Children: nil}
	}, []example.Break[console.ConsoleContent]{
		{Want: "console.ConsoleContent needs a Title", Change: func(p *console.ConsoleContent) { p.Title = "" }},
		{Want: "console.ConsoleContent needs a Back with both an Href and a Label, or neither", Change: func(p *console.ConsoleContent) { p.Back.Href = "/" }},
		{Want: "console.ConsoleContent needs a Back with both an Href and a Label, or neither", Change: func(p *console.ConsoleContent) { p.Back.Label = "Home" }},
	})
}
