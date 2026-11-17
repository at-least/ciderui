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
		console.ConsoleBtn{Href: "", Label: "Sign Out", Submit: true, Variant: console.ConsoleBtnVariantPlain, Size: console.ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: ""},
	)
	return console.ConsoleSidebar{AriaLabel: "CrashCart", Home: "#home", Wordmark: "CrashCart", SkipLabel: "Skip to Content", OpenAriaLabel: "Menu", CloseAriaLabel: "Close", Footer: signOut, Current: "#issues", Sections: []console.ConsoleMenuSection{
		{Title: "", Items: []console.ConsoleMenuItem{
			{Href: "#projects", Label: "All Projects", Icon: lucide.LayoutGrid},
			{Href: "#issues", Label: "All Issues", Icon: lucide.Bug},
		}},
		{Title: "Projects", Items: []console.ConsoleMenuItem{
			{Href: "#shop", Label: "Shop", Icon: lucide.Icon{}},
			{Href: "#demo", Label: "Demo Shop", Icon: lucide.Icon{}},
		}},
		{Title: "", Items: []console.ConsoleMenuItem{
			{Href: "#alerts", Label: "Alerts", Icon: lucide.Bell},
		}},
	}}
}

func TestSidebarRendersTheExamples(t *testing.T) {
	example.Check(t, sidebarNodes(), "console/console-sidebar/examples/sidebar.html")
}

func TestContentWritesTheSidebarsOpenButtonExample(t *testing.T) {
	sidebar := sidebarNodes()
	got := example.Render(console.ConsoleContent{Sidebar: &sidebar, Back: nil, Title: "x", Detail: "", Actions: nil, Children: nil})
	nav, found := strings.CutPrefix(got, contentOpen)
	nav, main, foundMain := strings.Cut(nav, contentMain)
	open, _, foundHeader := strings.Cut(main, `<header class="cui-console-header">`)
	if !found || !foundMain || !foundHeader {
		t.Fatalf("the content is not its sidebar, then its <main> with the open button first: %s", got)
	}
	example.Check(t, g.Group{g.Raw(open), g.Raw(nav)}, "console/console-sidebar/examples/open.html")
}

func TestSidebarRulesOffEachUntitledSectionAfterTheFirst(t *testing.T) {
	item := func(href, label string) console.ConsoleMenuSection {
		return console.ConsoleMenuSection{Title: "", Items: []console.ConsoleMenuItem{{Href: href, Label: label, Icon: lucide.Icon{}}}}
	}
	sidebar := sidebarNodes()
	sidebar.Sections, sidebar.Current = []console.ConsoleMenuSection{item("#a", "A"), item("#b", "B")}, ""
	want := `data-variant="sidebar"><a class="cui-console-menu-item" href="#a"><span class="cui-console-menu-label">A</span></a>` +
		`<hr class="cui-console-menu-rule"><a class="cui-console-menu-item" href="#b"><span class="cui-console-menu-label">B</span></a></div>`
	if got := example.Render(sidebar); !strings.Contains(got, want) {
		t.Errorf("two untitled sections are not the first's items, a rule, then the second's:\n got: %s\nwant: …%s…", got, want)
	}
}

func TestSidebarPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleSidebar {
		return console.ConsoleSidebar{AriaLabel: "x", Home: "/", Wordmark: "x", SkipLabel: "Skip", OpenAriaLabel: "Menu", CloseAriaLabel: "Close", Footer: nil, Current: "#y", Sections: []console.ConsoleMenuSection{
			{Title: "", Items: []console.ConsoleMenuItem{{Href: "#x", Label: "x", Icon: lucide.Icon{}}}},
			{Title: "", Items: []console.ConsoleMenuItem{{Href: "#y", Label: "y", Icon: lucide.Icon{}}, {Href: "#z", Label: "z", Icon: lucide.Icon{}}}},
		}}
	}, []example.Break[console.ConsoleSidebar]{
		{Want: "console.ConsoleSidebar needs an AriaLabel", Change: func(p *console.ConsoleSidebar) { p.AriaLabel = "" }},
		{Want: "console.ConsoleSidebar needs a Home", Change: func(p *console.ConsoleSidebar) { p.Home = "" }},
		{Want: "console.ConsoleSidebar needs a Wordmark", Change: func(p *console.ConsoleSidebar) { p.Wordmark = "" }},
		{Want: "console.ConsoleSidebar needs a SkipLabel", Change: func(p *console.ConsoleSidebar) { p.SkipLabel = "" }},
		{Want: "console.ConsoleSidebar needs an OpenAriaLabel", Change: func(p *console.ConsoleSidebar) { p.OpenAriaLabel = "" }},
		{Want: "console.ConsoleSidebar needs a CloseAriaLabel", Change: func(p *console.ConsoleSidebar) { p.CloseAriaLabel = "" }},
		{Want: "console.ConsoleSidebar needs Sections", Change: func(p *console.ConsoleSidebar) { p.Sections = nil }},
		{Want: "console.ConsoleSidebar.Sections[1] needs Items", Change: func(p *console.ConsoleSidebar) { p.Sections[1].Items = nil }},
		{Want: "console.ConsoleSidebar.Sections[1].Items[1] needs an Href", Change: func(p *console.ConsoleSidebar) { p.Sections[1].Items[1].Href = "" }},
		{Want: "console.ConsoleSidebar.Sections[1].Items[1] needs a Label", Change: func(p *console.ConsoleSidebar) { p.Sections[1].Items[1].Label = "" }},
		{Want: `console.ConsoleSidebar.Current "#nowhere" matches no item's Href`, Change: func(p *console.ConsoleSidebar) { p.Current = "#nowhere" }},
	})
}
