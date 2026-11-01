package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestSidebarRendersTheExample(t *testing.T) {
	account := h.A(h.Class("cui-console-sidebar-account"), h.Href("#account"), lucide.CircleUser.Class("cui-console-sidebar-avatar"), g.Text("Alex Morgan"))
	sidebar := console.ConsoleSidebar{ID: "sidebar", Label: "CrashCart", Home: "#home", Wordmark: "CrashCart", Close: "Close", Class: "", Footer: account, Children: g.Group{
		console.ConsoleSidebarItem{Href: "#projects", Label: "All Projects", Icon: lucide.LayoutGrid, Current: false, Secondary: "", Class: ""},
		console.ConsoleSidebarItem{Href: "#issues", Label: "All Issues", Icon: lucide.Bug, Current: true, Secondary: "", Class: ""},
		h.P(h.Class("cui-console-sidebar-heading"), g.Text("Projects")),
		console.ConsoleSidebarItem{Href: "#shop", Label: "Shop", Icon: icon.Icon{}, Current: false, Secondary: "8", Class: ""},
		console.ConsoleSidebarItem{Href: "#demo", Label: "Demo Shop", Icon: icon.Icon{}, Current: false, Secondary: "12", Class: ""},
		h.Hr(h.Class("cui-console-sidebar-rule")),
		console.ConsoleSidebarItem{Href: "#alerts", Label: "Alerts", Icon: lucide.Bell, Current: false, Secondary: "", Class: ""},
	}}
	example.Check(t, sidebar, "console/console-sidebar/examples/sidebar.html")
}

func TestSidebarWithNoCloseLabelPanics(t *testing.T) {
	mustPanic(t, "needs a Close", func() {
		example.Render(console.ConsoleSidebar{ID: "s", Label: "x", Home: "/", Wordmark: "x", Close: "", Class: "", Footer: nil, Children: nil})
	})
}
