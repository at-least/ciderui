package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestSidebarRendersTheExample(t *testing.T) {
	account := h.A(h.Class("cui-console-sidebar-account"), h.Href("#account"), lucide.CircleUser.Class("cui-console-sidebar-avatar"), g.Text("Alex Morgan"))
	sidebar := console.ConsoleSidebar{ID: "sidebar", AriaLabel: "CrashCart", Home: "#home", Wordmark: "CrashCart", CloseAriaLabel: "Close", Class: "", Footer: account, Children: g.Group{
		console.ConsoleMenuItem{Href: "#projects", Label: "All Projects", Icon: lucide.LayoutGrid, Current: false, Class: ""},
		console.ConsoleMenuItem{Href: "#issues", Label: "All Issues", Icon: lucide.Bug, Current: true, Class: ""},
		h.P(h.Class("cui-console-menu-heading"), g.Text("Projects")),
		console.ConsoleMenuItem{Href: "#shop", Label: "Shop", Icon: lucide.Icon{}, Current: false, Class: ""},
		console.ConsoleMenuItem{Href: "#demo", Label: "Demo Shop", Icon: lucide.Icon{}, Current: false, Class: ""},
		h.Hr(h.Class("cui-console-menu-rule")),
		console.ConsoleMenuItem{Href: "#alerts", Label: "Alerts", Icon: lucide.Bell, Current: false, Class: ""},
	}}
	example.Check(t, sidebar, "console/console-sidebar/examples/sidebar.html")
}

func TestSidebarPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleSidebar {
		return console.ConsoleSidebar{ID: "s", AriaLabel: "x", Home: "/", Wordmark: "x", CloseAriaLabel: "Close", Class: "", Footer: nil, Children: nil}
	}, []example.Break[console.ConsoleSidebar]{
		{Want: "console.ConsoleSidebar needs an ID", Change: func(p *console.ConsoleSidebar) { p.ID = "" }},
		{Want: "console.ConsoleSidebar needs an AriaLabel", Change: func(p *console.ConsoleSidebar) { p.AriaLabel = "" }},
		{Want: "console.ConsoleSidebar needs a Wordmark", Change: func(p *console.ConsoleSidebar) { p.Wordmark = "" }},
		{Want: "console.ConsoleSidebar needs a CloseAriaLabel", Change: func(p *console.ConsoleSidebar) { p.CloseAriaLabel = "" }},
	})
}
