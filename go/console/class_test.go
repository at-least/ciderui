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

func TestEachPartWritesThePagesClass(t *testing.T) {
	const page = example.PageClass
	player := pausedPlayer()
	player.Class = page
	example.PagesClass(t, ".", []string{"ConsolePage"}, []example.ClassRow{
		{Root: "cui-console-btn", Node: console.ConsoleBtn{Href: "", Label: "x", AriaLabel: "", Icon: lucide.Icon{}, Submit: false, Variant: console.ConsoleBtnVariantGray, Size: console.ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: page}},
		{Root: "cui-console-badge", Node: console.ConsoleBadge{Label: "x", Color: console.ConsoleColorNeutral, Class: page}},
		{Root: "cui-console-banner", Node: console.ConsoleBanner{Color: console.ConsoleColorInfo, Title: "x", Body: nil, Class: page, Children: nil}},
		{Root: "cui-console-chart", Node: console.ConsoleChart{AriaLabel: "x", Scale: nil, Empty: "", Axis: nil, Class: page, Children: nil}},
		{Root: "cui-console-dialog", Node: console.ConsoleDialog{ID: "x", Title: g.Text("x"), Attrs: nil, Class: page, Children: nil}},
		{Root: "cui-console-disclosure", Node: console.ConsoleDisclosure{Summary: "x", SummaryClass: "", Class: page, Children: nil}},
		{Root: "cui-console-disclosure-summary", Node: console.ConsoleDisclosure{Summary: "x", SummaryClass: page, Class: "", Children: nil}},
		{Root: "cui-console-empty", Node: console.ConsoleEmpty{Size: console.ConsoleEmptySizeSmall, Icon: lucide.Icon{}, Title: "x", Body: "", Class: page}},
		{Root: "cui-console-select", Node: console.ConsoleSelect{Class: page, Children: h.Select(h.Class("cui-console-field-input"))}},
		{Root: "cui-console-list-item", Node: console.ConsoleListItem{Href: "", Label: "x", Detail: "", Class: page, Children: nil}},
		{Root: "cui-console-list-item", Node: console.ConsoleListItem{Href: "#a", Label: "x", Detail: "", Class: page, Children: nil}},
		{Root: "cui-console-menu-item", Node: console.ConsoleMenuItem{Href: "#a", Label: "x", Icon: lucide.Icon{}, Current: false, Class: page}},
		{Root: "cui-console-menu-item", Node: console.ConsoleMenuItem{Href: "", Label: "x", Icon: lucide.Icon{}, Current: false, Class: page}},
		{Root: "cui-console-popover", Node: console.ConsoleChoiceMenu{ID: "x", Class: page, Columns: choiceMenu().Columns}},
		{Root: "cui-console-btn cui-console-btn-icon cui-console-play", Node: console.ConsolePlay{AriaLabel: "x", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: page}},
		{Root: "cui-console-section-header", Node: console.ConsoleSectionHeader{Title: g.Text("x"), Class: page, Children: nil}},
		{Root: "cui-console-sidebar", Node: console.ConsoleSidebar{AriaLabel: "x", Home: "/", Wordmark: "x", CloseAriaLabel: "x", Class: page, Footer: nil, Children: nil}},
		{Root: "cui-console-btn cui-console-btn-icon cui-console-sidebar-open", Node: console.ConsoleSidebarOpen{AriaLabel: "x", Class: page}},
		{Root: "cui-console-status", Node: console.ConsoleStatus{Label: "x", Color: console.ConsoleColorNeutral, Detail: "", Class: page}},
		{Root: "cui-console-strip", Node: console.ConsoleStrip{AriaLabel: "x", Class: page, Groups: validStrip().Groups}},
		{Root: "cui-console-table", Node: console.ConsoleTable{Columns: []console.ConsoleTableColumn{{Label: "x", Figure: false}}, Rows: nil, BodyAttrs: nil, Class: page}},
		{Root: "cui-console-tabs", Node: console.ConsoleTabs{AriaLabel: "x", Tabs: nil, Class: page}},
		{Root: "cui-console-player", Node: player},
	})
}

func TestEachPartWritesThePagesWords(t *testing.T) {
	intro := intro()
	intro.Intro = &console.ConsolePlayerIntro{Title: "The Feast", Detail: "2 min", Summary: "A short one.", PlayLabel: "Watch", RestartLabel: "From the Top"}
	end := end()
	end.End = &console.ConsolePlayerEnd{Title: "All Done", ReplayLabel: "Watch Again"}
	tip := validStrip()
	tip.Groups[0].Items[0].Tip = "Hello there."
	skip := consolePage()
	skip.SkipLabel, skip.Children = "Jump to Main", h.Main(h.ID("main"), h.TabIndex("-1"))
	for _, c := range []struct {
		part  string
		node  g.Node
		words []string
	}{
		{"ConsoleBanner", console.ConsoleBanner{Color: console.ConsoleColorInfo, Title: "Synced", Body: nil, Class: "", Children: nil}, []string{">Synced<"}},
		{"ConsoleChart", console.ConsoleChart{AriaLabel: "Events: 9", Scale: &console.ConsoleChartScale{Peak: "peak 9", Half: "4.5"}, Empty: "Nothing yet", Axis: &console.ConsoleChartAxis{Start: "Nov 1", End: "Nov 7"}, Class: "", Children: nil},
			[]string{`aria-label="Events: 9"`, ">peak 9<", ">4.5<", ">Nothing yet<", ">Nov 1<", ">Nov 7<"}},
		{"ConsoleField", console.ConsoleField{ID: "name", Label: "Name", Error: "Too short.", Control: input()}, []string{">Name<", ">Too short.<"}},
		{"ConsolePlay", console.ConsolePlay{AriaLabel: "Listen", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: ""}, []string{`aria-label="Listen"`}},
		{"ConsolePlay", console.ConsolePlay{AriaLabel: "", Label: "Hear It", PlayingLabel: "Halt", Size: console.ConsoleBtnSizeMedium, Class: ""}, []string{">Hear It<", ">Halt<"}},
		{"ConsolePage", skip, []string{">Jump to Main<"}},
		{"ConsolePlayerIntro", intro, []string{">The Feast<", ">2 min<", ">A short one.<", ">Watch<", ">From the Top<"}},
		{"ConsolePlayerEnd", end, []string{">All Done<", ">Watch Again<"}},
		{"ConsoleSidebar", console.ConsoleSidebar{AriaLabel: "Main", Home: "/home", Wordmark: "Muse", CloseAriaLabel: "Hide", Class: "", Footer: nil, Children: nil}, []string{`aria-label="Main"`, `href="/home"`, ">Muse<", `aria-label="Hide"`}},
		{"ConsoleSidebarOpen", console.ConsoleSidebarOpen{AriaLabel: "Show Menu", Class: ""}, []string{`aria-label="Show Menu"`}},
		{"ConsoleStatus", console.ConsoleStatus{Label: "Down", Color: console.ConsoleColorError, Detail: "500 · 2 s", Class: ""}, []string{">Down<", ">500 · 2 s<"}},
		{"ConsoleStrip", tip, []string{">Hello there.<"}},
	} {
		got := example.Render(c.node)
		for _, word := range c.words {
			if !strings.Contains(got, word) {
				t.Errorf("%s does not write the page's %s: %s", c.part, word, got)
			}
		}
	}
}
