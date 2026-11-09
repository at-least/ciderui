package console_test

import (
	"strings"
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
	example.Check(t, console.ConsoleTabs{AriaLabel: "Project", Tabs: underlined}, "console/console-tabs/examples/tabs.html")
}

// not the link a current tab is: no aria-current on what cannot be followed
func TestTabWithNoHrefIsNotAvailable(t *testing.T) {
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
		{Want: "console.ConsoleTabs has one current tab at most", Change: func(p *console.ConsoleTabs) { p.Tabs[1].Current = true }},
		{Want: "console.ConsoleTab needs a Label", Change: func(p *console.ConsoleTabs) { p.Tabs[1].Label = "" }},
	})
}
