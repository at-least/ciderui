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
	}
	example.Check(t, console.ConsoleTabs{AriaLabel: "Project", Tabs: underlined}, "console/console-tabs/examples/tabs.html")
}

func TestTabsOfOneTabRenders(t *testing.T) {
	got := example.Render(console.ConsoleTabs{AriaLabel: "Project", Tabs: []console.ConsoleTab{{Label: "Overview", Href: "#overview", Current: false}}})
	if want := `<nav class="cui-console-tabs" aria-label="Project"><a class="cui-console-tab" href="#overview">Overview</a></nav>`; got != want {
		t.Errorf("tabs of one tab are not that tab:\n got: %s\nwant: %s", got, want)
	}
}

func TestTabsPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleTabs {
		return console.ConsoleTabs{AriaLabel: "X", Tabs: []console.ConsoleTab{{Label: "A", Href: "#a", Current: true}, {Label: "B", Href: "#b", Current: false}}}
	}, []example.Break[console.ConsoleTabs]{
		{Want: "console.ConsoleTabs needs an AriaLabel", Change: func(p *console.ConsoleTabs) { p.AriaLabel = "" }},
		{Want: "console.ConsoleTabs needs Tabs", Change: func(p *console.ConsoleTabs) { p.Tabs = nil }},
		{Want: "console.ConsoleTabs has 2 current tabs: one at most", Change: func(p *console.ConsoleTabs) { p.Tabs[1].Current = true }},
		{Want: "console.ConsoleTab needs a Label", Change: func(p *console.ConsoleTabs) { p.Tabs[1].Label = "" }},
		{Want: "console.ConsoleTab needs an Href", Change: func(p *console.ConsoleTabs) { p.Tabs[1].Href = "" }},
	})
}

func tabsChoice() console.ConsoleTabsChoice {
	return console.ConsoleTabsChoice{AriaLabel: "About", Name: "about", Options: []console.ConsoleTabsOption{
		{Value: "picture", Label: "The picture", Checked: true},
		{Value: "line", Label: "The line", Checked: false},
	}}
}

func TestTabsChoiceRendersTheExample(t *testing.T) {
	example.Check(t, tabsChoice(), "console/console-tabs/examples/radios.html")
}

func TestTabsChoicePanics(t *testing.T) {
	example.Panics(t, tabsChoice, []example.Break[console.ConsoleTabsChoice]{
		{Want: "console.ConsoleTabsChoice needs an AriaLabel", Change: func(p *console.ConsoleTabsChoice) { p.AriaLabel = "" }},
		{Want: "console.ConsoleTabsChoice needs a Name", Change: func(p *console.ConsoleTabsChoice) { p.Name = "" }},
		{Want: "console.ConsoleTabsChoice needs two Options at least: one is no choice", Change: func(p *console.ConsoleTabsChoice) { p.Options = p.Options[:1] }},
		{Want: "console.ConsoleTabsOption needs a Label", Change: func(p *console.ConsoleTabsChoice) { p.Options[1].Label = "" }},
		{Want: "console.ConsoleTabsChoice needs one Checked option, has 0: a radio group starts with its value", Change: func(p *console.ConsoleTabsChoice) { p.Options[0].Checked = false }},
		{Want: "console.ConsoleTabsChoice needs one Checked option, has 2: a radio group starts with its value", Change: func(p *console.ConsoleTabsChoice) { p.Options[1].Checked = true }},
	})
}
