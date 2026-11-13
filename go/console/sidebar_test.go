package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func sidebarNodes() console.ConsoleSidebar {
	signOut := h.Form(h.Action("#sign-out"), h.Method("post"),
		console.ConsoleBtn{Href: "", Label: "Sign Out", AriaLabel: "", Icon: lucide.Icon{}, Submit: true, Variant: console.ConsoleBtnVariantPlain, Size: console.ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: ""},
	)
	return console.ConsoleSidebar{AriaLabel: "CrashCart", Home: "#home", Wordmark: "CrashCart", SkipLabel: "Skip to Content", OpenAriaLabel: "Menu", CloseAriaLabel: "Close", Footer: signOut, Sections: []console.ConsoleMenuSection{
		{Title: "", Items: []console.ConsoleMenuItem{
			{Href: "#projects", Label: "All Projects", Icon: lucide.LayoutGrid, Current: false},
			{Href: "#issues", Label: "All Issues", Icon: lucide.Bug, Current: true},
		}},
		{Title: "Projects", Items: []console.ConsoleMenuItem{
			{Href: "#shop", Label: "Shop", Icon: lucide.Icon{}, Current: false},
			{Href: "#demo", Label: "Demo Shop", Icon: lucide.Icon{}, Current: false},
		}},
		{Title: "", Items: []console.ConsoleMenuItem{
			{Href: "#alerts", Label: "Alerts", Icon: lucide.Bell, Current: false},
		}},
	}}
}

func TestSidebarRendersTheExamples(t *testing.T) {
	example.Check(t, sidebarNodes(), "console/console-sidebar/examples/sidebar.html")
}

func TestContentWritesTheSidebarsOpenButtonExample(t *testing.T) {
	sidebar := sidebarNodes()
	got := example.Render(console.ConsoleContent{Sidebar: &sidebar, Back: nil, Title: "x", Detail: "", Controls: nil, Children: nil})
	nav, found := strings.CutPrefix(got, contentOpen)
	nav, main, foundMain := strings.Cut(nav, contentMain)
	open, _, foundHeader := strings.Cut(main, `<header class="cui-console-header">`)
	if !found || !foundMain || !foundHeader {
		t.Fatalf("the content is not its sidebar, then its <main> with the open button first: %s", got)
	}
	example.Check(t, g.Group{g.Raw(open), g.Raw(nav)}, "console/console-sidebar/examples/open.html")
}

func TestSidebarPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleSidebar {
		return console.ConsoleSidebar{AriaLabel: "x", Home: "/", Wordmark: "x", SkipLabel: "Skip", OpenAriaLabel: "Menu", CloseAriaLabel: "Close", Footer: nil, Sections: nil}
	}, []example.Break[console.ConsoleSidebar]{
		{Want: "console.ConsoleSidebar needs an AriaLabel", Change: func(p *console.ConsoleSidebar) { p.AriaLabel = "" }},
		{Want: "console.ConsoleSidebar needs a Home", Change: func(p *console.ConsoleSidebar) { p.Home = "" }},
		{Want: "console.ConsoleSidebar needs a Wordmark", Change: func(p *console.ConsoleSidebar) { p.Wordmark = "" }},
		{Want: "console.ConsoleSidebar needs a SkipLabel", Change: func(p *console.ConsoleSidebar) { p.SkipLabel = "" }},
		{Want: "console.ConsoleSidebar needs an OpenAriaLabel", Change: func(p *console.ConsoleSidebar) { p.OpenAriaLabel = "" }},
		{Want: "console.ConsoleSidebar needs a CloseAriaLabel", Change: func(p *console.ConsoleSidebar) { p.CloseAriaLabel = "" }},
	})
}
