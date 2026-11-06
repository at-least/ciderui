package console_test

import (
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
		{Root: "cui-console-btn", Node: console.ConsoleBtn{Href: "", Label: "x", AriaLabel: "", Icon: lucide.Icon{}, Submit: false, Variant: console.ConsoleBtnVariantGray, Size: console.ConsoleBtnSizeMedium, Error: false, Attrs: nil, Class: page}},
		{Root: "cui-console-badge", Node: console.ConsoleBadge{Label: "x", Color: console.ConsoleColorNeutral, Class: page}},
		{Root: "cui-console-banner", Node: console.ConsoleBanner{Color: console.ConsoleColorInfo, Icon: lucide.Icon{}, Title: "x", Class: page, Children: nil}},
		{Root: "cui-console-chart", Node: console.ConsoleChart{AriaLabel: "x", Scale: nil, Empty: "", Axis: nil, Class: page, Children: nil}},
		{Root: "cui-console-dialog", Node: console.ConsoleDialog{ID: "x", Title: g.Text("x"), Attrs: nil, Class: page, Children: nil}},
		{Root: "cui-console-disclosure", Node: console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantLine, Open: false, Summary: "x", SummaryClass: "", Class: page, Children: nil}},
		{Root: "cui-console-btn", Node: console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantButton, Open: false, Summary: "x", SummaryClass: page, Class: "", Children: nil}},
		{Root: "cui-console-disclosure", Node: console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantButton, Open: false, Summary: "x", SummaryClass: "", Class: page, Children: nil}},
		{Root: "cui-console-disclosure-summary", Node: console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantLine, Open: false, Summary: "x", SummaryClass: page, Class: "", Children: nil}},
		{Root: "cui-console-empty", Node: console.ConsoleEmpty{Size: console.ConsoleEmptySizeSmall, Backed: false, Icon: lucide.Icon{}, Heading: "x", Caption: "", Class: page, Children: nil}},
		{Root: "cui-console-select", Node: console.ConsoleSelect{Class: page, Children: h.Select(h.Class("cui-console-field"))}},
		{Root: "cui-console-list-item", Node: console.ConsoleListItem{Href: "", ID: "", Label: "x", Secondary: "", Class: page, Children: nil}},
		{Root: "cui-console-list-item", Node: console.ConsoleListItem{Href: "#a", ID: "", Label: "x", Secondary: "", Class: page, Children: nil}},
		{Root: "cui-console-menu-item", Node: console.ConsoleMenuItem{Href: "#a", Label: "x", Icon: lucide.Icon{}, Current: false, Class: page}},
		{Root: "cui-console-menu-item", Node: console.ConsoleMenuItem{Href: "", Label: "x", Icon: lucide.Icon{}, Current: false, Class: page}},
		{Root: "cui-console-popover", Node: console.ConsoleChoiceMenu{ID: "x", Class: page, Columns: choiceMenu().Columns}},
		{Root: "cui-console-btn cui-console-btn-icon cui-console-play", Node: console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "x", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: page}},
		{Root: "cui-console-section-head", Node: console.ConsoleSectionHead{Title: g.Text("x"), Class: page, Children: nil}},
		{Root: "cui-console-sidebar", Node: console.ConsoleSidebar{ID: "x", AriaLabel: "x", Home: "/", Wordmark: "x", CloseAriaLabel: "x", Class: page, Footer: nil, Children: nil}},
		{Root: "cui-console-status", Node: console.ConsoleStatus{Label: "x", Color: console.ConsoleColorNeutral, Class: page}},
		{Root: "cui-console-strip", Node: console.ConsoleStrip{AriaLabel: "x", Class: page, Groups: validStrip().Groups}},
		{Root: "cui-console-table-scroll", Node: console.ConsoleTable{AriaLabel: "x", Columns: []console.ConsoleTableColumn{{Label: "x", Figure: false}}, Rows: nil, BodyAttrs: nil, Class: page}},
		{Root: "cui-console-tabs", Node: console.ConsoleTabs{AriaLabel: "x", Variant: console.ConsoleTabsVariantUnderlined, Tabs: nil, Class: page}},
		{Root: "cui-console-tabs", Node: console.ConsoleTabs{AriaLabel: "x", Variant: console.ConsoleTabsVariantFilled, Tabs: nil, Class: page}},
		{Root: "cui-console-tabs", Node: console.ConsoleTabButtons{AriaLabel: "x", Name: "x", Buttons: nil, Class: page}},
		{Root: "cui-console-player", Node: player},
	})
}
