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
	example.Check(t, console.ConsoleTabs{Label: "Project", Variant: console.ConsoleTabsVariantUnderlined, Tabs: underlined, Class: ""}, "console/console-tabs/examples/tabs.html")
	filled := []console.ConsoleTab{
		{Label: "24h", Href: "#24h", Current: false},
		{Label: "7d", Href: "#7d", Current: true},
		{Label: "30d", Href: "#30d", Current: false},
		{Label: "90d", Href: "#90d", Current: false},
	}
	example.Check(t, console.ConsoleTabs{Label: "Period", Variant: console.ConsoleTabsVariantFilled, Tabs: filled, Class: ""}, "console/console-tabs/examples/filled.html")
}

func TestTabsTwoCurrentTabsPanic(t *testing.T) {
	tabs := []console.ConsoleTab{{Label: "A", Href: "#a", Current: true}, {Label: "B", Href: "#b", Current: true}}
	mustPanic(t, "one current tab at most", func() {
		example.Render(console.ConsoleTabs{Label: "X", Variant: console.ConsoleTabsVariantUnderlined, Tabs: tabs, Class: ""})
	})
}

// not in the Rust crate, whose enum has no zero value
func TestTabsTheZeroVariantPanics(t *testing.T) {
	mustPanic(t, "needs a Variant", func() {
		example.Render(console.ConsoleTabs{Label: "X", Variant: console.ConsoleTabsVariant{}, Tabs: nil, Class: ""})
	})
}
