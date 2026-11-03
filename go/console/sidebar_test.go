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
		console.ConsoleMenuItem{Href: "#projects", Label: "All Projects", Icon: lucide.LayoutGrid, Current: false, Disabled: false, Secondary: "", Class: ""},
		console.ConsoleMenuItem{Href: "#issues", Label: "All Issues", Icon: lucide.Bug, Current: true, Disabled: false, Secondary: "", Class: ""},
		h.P(h.Class("cui-console-menu-heading"), g.Text("Projects")),
		console.ConsoleMenuItem{Href: "#shop", Label: "Shop", Icon: icon.Icon{}, Current: false, Disabled: false, Secondary: "8", Class: ""},
		console.ConsoleMenuItem{Href: "#demo", Label: "Demo Shop", Icon: icon.Icon{}, Current: false, Disabled: false, Secondary: "12", Class: ""},
		h.Hr(h.Class("cui-console-menu-rule")),
		console.ConsoleMenuItem{Href: "#alerts", Label: "Alerts", Icon: lucide.Bell, Current: false, Disabled: false, Secondary: "", Class: ""},
	}}
	example.Check(t, sidebar, "console/console-sidebar/examples/sidebar.html")
}

func TestSidebarWithNoCloseLabelPanics(t *testing.T) {
	mustPanic(t, "needs a Close", func() {
		example.Render(console.ConsoleSidebar{ID: "s", Label: "x", Home: "/", Wordmark: "x", Close: "", Class: "", Footer: nil, Children: nil})
	})
}

func TestSidebarWithNoIDPanics(t *testing.T) {
	mustPanic(t, "ConsoleSidebar needs an ID", func() {
		example.Render(console.ConsoleSidebar{ID: "", Label: "x", Home: "/", Wordmark: "x", Close: "Close", Class: "", Footer: nil, Children: nil})
	})
}

func TestSidebarWithNoLabelPanics(t *testing.T) {
	mustPanic(t, "ConsoleSidebar needs a Label", func() {
		example.Render(console.ConsoleSidebar{ID: "s", Label: "", Home: "/", Wordmark: "x", Close: "Close", Class: "", Footer: nil, Children: nil})
	})
}

func TestSidebarWithNoWordmarkPanics(t *testing.T) {
	mustPanic(t, "ConsoleSidebar needs a Wordmark", func() {
		example.Render(console.ConsoleSidebar{ID: "s", Label: "x", Home: "/", Wordmark: "", Close: "Close", Class: "", Footer: nil, Children: nil})
	})
}
