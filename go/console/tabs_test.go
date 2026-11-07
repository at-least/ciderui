package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
)

func TestTabsRendersTheExamples(t *testing.T) {
	underlined := []console.ConsoleTab{
		{Label: "Overview", Href: "#overview", Current: true},
		{Label: "Issues", Href: "#issues", Current: false},
		{Label: "Releases", Href: "#releases", Current: false},
		{Label: "Jobs", Href: "", Current: false},
	}
	example.Check(t, console.ConsoleTabs{AriaLabel: "Project", Variant: console.ConsoleTabsVariantUnderlined, Tabs: underlined, Class: ""}, "console/console-tabs/examples/tabs.html")
	filled := []console.ConsoleTab{
		{Label: "24h", Href: "#24h", Current: false},
		{Label: "7d", Href: "#7d", Current: true},
		{Label: "30d", Href: "#30d", Current: false},
		{Label: "90d", Href: "#90d", Current: false},
	}
	example.Check(t, console.ConsoleTabs{AriaLabel: "Period", Variant: console.ConsoleTabsVariantFilled, Tabs: filled, Class: ""}, "console/console-tabs/examples/filled.html")
}

func TestTabsPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleTabs {
		return console.ConsoleTabs{AriaLabel: "X", Variant: console.ConsoleTabsVariantUnderlined, Tabs: []console.ConsoleTab{{Label: "A", Href: "#a", Current: true}, {Label: "B", Href: "#b", Current: false}}, Class: ""}
	}, []example.Break[console.ConsoleTabs]{
		{Want: "console.ConsoleTabs needs an AriaLabel", Change: func(p *console.ConsoleTabs) { p.AriaLabel = "" }},
		{Want: "console.ConsoleTabs needs a Variant", Change: func(p *console.ConsoleTabs) { p.Variant = console.ConsoleTabsVariant{} }},
		{Want: "console.ConsoleTabs has one current tab at most", Change: func(p *console.ConsoleTabs) { p.Tabs[1].Current = true }},
		{Want: "console.ConsoleTab needs a Label", Change: func(p *console.ConsoleTabs) { p.Tabs[1].Label = "" }},
	})
}
