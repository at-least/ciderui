package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestEachPartWritesThePagesClass(t *testing.T) {
	const page = example.PageClass
	player := pausedPlayer()
	player.Class = page
	example.PagesClass(t, ".", []string{"ConsolePage"}, []example.ClassRow{
		{Root: "cui-console-btn", Node: console.ConsoleBtn{Href: "", Label: "x", AriaLabel: "", Icon: icon.Icon{}, Submit: false, Variant: console.ConsoleBtnVariantGray, Size: console.ConsoleBtnSizeMedium, Error: false, Capsule: false, Attrs: nil, Class: page}},
		{Root: "cui-console-badge", Node: console.ConsoleBadge{Label: "x", Color: console.ConsoleColorNeutral, Class: page}},
		{Root: "cui-console-banner", Node: console.ConsoleBanner{Color: console.ConsoleColorInfo, Icon: icon.Icon{}, Title: "x", Class: page, Children: nil}},
		{Root: "cui-console-card", Node: console.ConsoleMetric{Title: "x", Value: g.Text("1"), Change: "", ChangeColor: console.ConsoleChangeColorNeutral, Caption: "", Href: "", Class: page}},
		{Root: "cui-console-card", Node: console.ConsoleMetric{Title: "x", Value: g.Text("1"), Change: "", ChangeColor: console.ConsoleChangeColorNeutral, Caption: "", Href: "#a", Class: page}},
		{Root: "cui-console-chart", Node: console.ConsoleChart{AriaLabel: "x", Scale: nil, Empty: "", Axis: nil, Legend: nil, Class: page, Children: nil}},
		{Root: "cui-console-chart-sparkline", Node: console.ConsoleSparkline{AriaLabel: "x", Width: 1, Height: 1, Values: []float64{0, 1}, Class: page}},
		{Root: "cui-console-chip", Node: console.ConsoleChip{Key: "k", Value: "v", Href: "", Remove: nil, Class: page}},
		{Root: "cui-console-chip", Node: console.ConsoleChip{Key: "k", Value: "v", Href: "#a", Remove: nil, Class: page}},
		{Root: "cui-console-dialog", Node: console.ConsoleDialog{ID: "x", Title: g.Text("x"), Attrs: nil, Class: page, Children: nil}},
		{Root: "cui-console-disclosure", Node: console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantLine, Open: false, Summary: "x", SummaryClass: "", Class: page, Children: nil}},
		{Root: "cui-console-disclosure-summary", Node: console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantLine, Open: false, Summary: "x", SummaryClass: page, Class: "", Children: nil}},
		{Root: "cui-console-disclosure", Node: console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantButton, Open: false, Summary: "x", SummaryClass: "", Class: page, Children: nil}},
		{Root: "cui-console-btn", Node: console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantButton, Open: false, Summary: "x", SummaryClass: page, Class: "", Children: nil}},
		{Root: "cui-console-empty", Node: console.ConsoleEmpty{Size: console.ConsoleEmptySizeSmall, Backed: false, Icon: icon.Icon{}, Heading: "x", Caption: "", Class: page, Children: nil}},
		{Root: "cui-console-select", Node: console.ConsoleSelect{Class: page, Children: h.Select(h.Class("cui-console-field"))}},
		{Root: "cui-console-error", Node: console.ConsoleError{ControlID: "x", Label: "x", Class: page}},
		{Root: "cui-console-head", Node: console.ConsoleHead{Name: "x", Href: "/", TabsAriaLabel: "x", Tabs: nil, Class: page}},
		{Root: "cui-console-list-item", Node: console.ConsoleListItem{Href: "", ID: "", Label: "x", Secondary: "", Class: page, Children: nil}},
		{Root: "cui-console-list-item", Node: console.ConsoleListItem{Href: "#a", ID: "", Label: "x", Secondary: "", Class: page, Children: nil}},
		{Root: "cui-console-listing", Node: console.ConsoleListing{FirstLine: 1, Lines: []console.ConsoleListingLine{{Label: "a", Highlighted: false}}, Class: page}},
		{Root: "cui-console-menu", Node: console.ConsoleMenu{AriaLabel: "x", Class: page, Children: nil}},
		{Root: "cui-console-menu-item", Node: console.ConsoleMenuItem{Href: "#a", Label: "x", Icon: icon.Icon{}, Current: false, Secondary: "", Class: page}},
		{Root: "cui-console-menu-item", Node: console.ConsoleMenuItem{Href: "", Label: "x", Icon: icon.Icon{}, Current: false, Secondary: "", Class: page}},
		{Root: "cui-console-menu-group", Node: console.ConsoleMenuGroup{Label: "x", Icon: icon.Icon{}, Open: false, Class: page, Items: []console.ConsoleMenuItem{validMenuItem()}}},
		{Root: "cui-console-popover", Node: console.ConsoleChoiceMenu{ID: "x", Class: page, Columns: choiceMenu().Columns}},
		{Root: "cui-console-pager", Node: console.ConsolePager{AriaLabel: "x", Summary: "x", First: "", FirstLabel: "x", Previous: "", PreviousAriaLabel: "x", Next: "", NextAriaLabel: "x", Class: page}},
		{Root: "cui-console-btn cui-console-btn-icon cui-console-play", Node: console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "x", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: page}},
		{Root: "cui-console-radio", Node: console.ConsoleRadio{Name: "x", Value: "x", Checked: false, Caption: "", Attrs: nil, Class: page, Children: g.Text("x")}},
		{Root: "cui-console-section-head", Node: console.ConsoleSectionHead{Title: g.Text("x"), Class: page, Children: nil}},
		{Root: "cui-console-sidebar", Node: console.ConsoleSidebar{ID: "x", AriaLabel: "x", Home: "/", Wordmark: "x", CloseAriaLabel: "x", Class: page, Footer: nil, Children: nil}},
		{Root: "cui-console-status", Node: console.ConsoleStatus{Label: "x", Color: console.ConsoleColorNeutral, Class: page}},
		{Root: "cui-console-strip", Node: console.ConsoleStrip{AriaLabel: "x", Class: page, Groups: validStrip().Groups}},
		{Root: "cui-console-table-scroll", Node: console.ConsoleTable{AriaLabel: "x", Columns: []console.ConsoleTableColumn{{Label: "x", Figure: false}}, Rows: nil, BodyAttrs: nil, Class: page}},
		{Root: "cui-console-tabs", Node: console.ConsoleTabs{AriaLabel: "x", Variant: console.ConsoleTabsVariantUnderlined, Tabs: nil, Class: page}},
		{Root: "cui-console-tabs", Node: console.ConsoleTabs{AriaLabel: "x", Variant: console.ConsoleTabsVariantFilled, Tabs: nil, Class: page}},
		{Root: "cui-console-tabs", Node: console.ConsoleTabButtons{AriaLabel: "x", Name: "x", Buttons: nil, Class: page}},
		{Root: "cui-console-thumb", Node: console.ConsoleThumb{Href: "#a", Src: "/a.png", Alt: "x", Caption: "x", Lazy: false, Class: page}},
		{Root: "cui-console-player", Node: player},
	})
}
