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
	sidebar := console.ConsoleSidebar{AriaLabel: "CrashCart", Home: "#home", Wordmark: "CrashCart", CloseAriaLabel: "Close", Footer: nil, Children: g.Group{
		console.ConsoleMenuItem{Href: "#projects", Label: "All Projects", Icon: lucide.Icon{}, Current: true},
		console.ConsoleMenuItem{Href: "#issues", Label: "All Issues", Icon: lucide.Icon{}, Current: false},
	}}
	return h.Div(h.Class("cui-console-page"), console.ConsoleContent{Sidebar: &sidebar, Children: g.Group{
		console.ConsoleSidebarOpen{AriaLabel: "Menu", Class: ""},
		console.ConsoleSectionHeader{Level: console.ConsoleSectionHeaderLevelPage, Title: "All Projects", Detail: "Every project you can see, the ones that need you first.", Children: nil},
	}})
}

func rhythmNodes() g.Node {
	return h.Div(h.Class("cui-console-page"), console.ConsoleContent{Sidebar: nil, Children: g.Group{
		console.ConsoleSectionHeader{Level: console.ConsoleSectionHeaderLevelPage, Title: "Findings", Detail: "", Children: nil},
		console.ConsoleSection{Title: "Needs You", Detail: "", Controls: nil, Children: g.Group{h.P(g.Text("A section’s first block.")), h.P(g.Text("Its second."))}},
		console.ConsoleSection{Title: "Closed", Detail: "", Controls: nil, Children: h.P(g.Text("The next section’s block."))},
	}})
}

func TestContentRendersTheExamples(t *testing.T) {
	example.Check(t, contentNodes(), "console/console-content/examples/content.html")
	example.Check(t, rhythmNodes(), "console/console-content/examples/rhythm.html")
}
