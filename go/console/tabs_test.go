package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestTabsRendersTheExamples(t *testing.T) {
	underlined := []console.ConsoleTab{
		{Label: "Overview", Href: "#overview", Current: true},
		{Label: "Issues", Href: "#issues", Current: false},
		{Label: "Releases", Href: "#releases", Current: false},
		{Label: "Jobs", Href: "", Current: false},
	}
	example.Check(t, console.ConsoleTabs{AriaLabel: "Project", Tabs: underlined}, "console/console-tabs/examples/tabs.html")
}

func TestTabWithNoHrefIsNotAvailableEvenIfCurrent(t *testing.T) {
	got := example.Render(console.ConsoleTabs{AriaLabel: "Project", Tabs: []console.ConsoleTab{{Label: "Jobs", Href: "", Current: true}}})
	if want := `<a class="cui-console-tab" aria-disabled="true">Jobs</a>`; !strings.Contains(got, want) {
		t.Errorf("a tab with no Href is a link or marked current, not unavailable:\n got: %s\nwant: …%s…", got, want)
	}
}

func TestTabsPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleTabs {
		return console.ConsoleTabs{AriaLabel: "X", Tabs: []console.ConsoleTab{{Label: "A", Href: "#a", Current: true}, {Label: "B", Href: "#b", Current: false}}}
	}, []example.Break[console.ConsoleTabs]{
		{Want: "console.ConsoleTabs needs an AriaLabel", Change: func(p *console.ConsoleTabs) { p.AriaLabel = "" }},
		{Want: "console.ConsoleTabs has 2 current tabs: one at most", Change: func(p *console.ConsoleTabs) { p.Tabs[1].Current = true }},
		{Want: "console.ConsoleTab needs a Label", Change: func(p *console.ConsoleTabs) { p.Tabs[1].Label = "" }},
	})
}

func tabsChoice() console.ConsoleTabsChoice {
	return console.ConsoleTabsChoice{AriaLabel: "About", Name: "about", Attrs: nil, Options: []console.ConsoleTabsOption{
		{Value: "picture", Label: "The picture", Checked: true},
		{Value: "line", Label: "The line", Checked: false},
	}}
}

func TestTabsChoiceRendersTheExample(t *testing.T) {
	example.Check(t, tabsChoice(), "console/console-tabs/examples/radios.html")
}

func TestTabsChoiceWritesThePagesAttrs(t *testing.T) {
	c := tabsChoice()
	c.Attrs = h.Data("report-about", "")
	if got := example.Render(c); !strings.HasPrefix(got, `<div class="cui-console-tabs" data-variant="filled" role="group" aria-label="About" data-report-about="">`) {
		t.Errorf("the choice does not write the page's Attrs on its group: %s", got)
	}
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
