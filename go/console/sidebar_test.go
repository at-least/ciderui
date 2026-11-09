package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func sidebarNodes() g.Node {
	signOut := h.Form(h.Action("#sign-out"), h.Method("post"),
		console.ConsoleBtn{Href: "", Label: "Sign Out", AriaLabel: "", Icon: lucide.Icon{}, Submit: true, Variant: console.ConsoleBtnVariantPlain, Size: console.ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: ""},
	)
	return console.ConsoleSidebar{AriaLabel: "CrashCart", Home: "#home", Wordmark: "CrashCart", CloseAriaLabel: "Close", Footer: signOut, Children: g.Group{
		console.ConsoleMenuItem{Href: "#projects", Label: "All Projects", Icon: lucide.LayoutGrid, Current: false},
		console.ConsoleMenuItem{Href: "#issues", Label: "All Issues", Icon: lucide.Bug, Current: true},
		h.P(h.Class("cui-console-menu-title"), g.Text("Projects")),
		console.ConsoleMenuItem{Href: "#shop", Label: "Shop", Icon: lucide.Icon{}, Current: false},
		console.ConsoleMenuItem{Href: "#demo", Label: "Demo Shop", Icon: lucide.Icon{}, Current: false},
		h.Hr(h.Class("cui-console-menu-rule")),
		console.ConsoleMenuItem{Href: "#alerts", Label: "Alerts", Icon: lucide.Bell, Current: false},
	}}
}

func TestSidebarRendersTheExamples(t *testing.T) {
	example.Check(t, sidebarNodes(), "console/console-sidebar/examples/sidebar.html")
	example.Check(t, g.Group{
		console.ConsoleSidebarOpen{AriaLabel: "Menu", Class: ""},
		sidebarNodes(),
	}, "console/console-sidebar/examples/open.html")
}

func TestSidebarOpenPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleSidebarOpen {
		return console.ConsoleSidebarOpen{AriaLabel: "Menu", Class: ""}
	}, []example.Break[console.ConsoleSidebarOpen]{
		{Want: "console.ConsoleSidebarOpen needs an AriaLabel", Change: func(p *console.ConsoleSidebarOpen) { p.AriaLabel = "" }},
	})
}

func TestSidebarPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleSidebar {
		return console.ConsoleSidebar{AriaLabel: "x", Home: "/", Wordmark: "x", CloseAriaLabel: "Close", Footer: nil, Children: nil}
	}, []example.Break[console.ConsoleSidebar]{
		{Want: "console.ConsoleSidebar needs an AriaLabel", Change: func(p *console.ConsoleSidebar) { p.AriaLabel = "" }},
		{Want: "console.ConsoleSidebar needs a Wordmark", Change: func(p *console.ConsoleSidebar) { p.Wordmark = "" }},
		{Want: "console.ConsoleSidebar needs a CloseAriaLabel", Change: func(p *console.ConsoleSidebar) { p.CloseAriaLabel = "" }},
	})
}
