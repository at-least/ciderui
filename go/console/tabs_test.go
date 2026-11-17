package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
)

func TestTabsRendersTheExamples(t *testing.T) {
	underlined := []console.ConsoleTab{
		{Label: "Overview", Href: "#overview"},
		{Label: "Issues", Href: "#issues"},
		{Label: "Releases", Href: "#releases"},
	}
	example.Check(t, console.ConsoleTabs{AriaLabel: "Project", Tabs: underlined, Current: "#overview"}, "console/console-tabs/examples/tabs.html")
}

func TestTabsOfOneTabRenders(t *testing.T) {
	got := example.Render(console.ConsoleTabs{AriaLabel: "Project", Tabs: []console.ConsoleTab{{Label: "Overview", Href: "#overview"}}, Current: ""})
	if want := `<nav class="cui-console-tabs" aria-label="Project"><a class="cui-console-tab" href="#overview">Overview</a></nav>`; got != want {
		t.Errorf("tabs of one tab are not that tab:\n got: %s\nwant: %s", got, want)
	}
}

func TestTabsPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleTabs {
		return console.ConsoleTabs{AriaLabel: "X", Tabs: []console.ConsoleTab{{Label: "A", Href: "#a"}, {Label: "B", Href: "#b"}}, Current: "#a"}
	}, []example.Break[console.ConsoleTabs]{
		{Want: "console.ConsoleTabs needs an AriaLabel", Change: func(p *console.ConsoleTabs) { p.AriaLabel = "" }},
		{Want: "console.ConsoleTabs needs Tabs", Change: func(p *console.ConsoleTabs) { p.Tabs = nil }},
		{Want: "console.ConsoleTabs needs a Current that is one of its Tabs' Href", Change: func(p *console.ConsoleTabs) { p.Current = "#c" }},
		{Want: "console.ConsoleTabs needs a Current that is one of its Tabs' Href", Change: func(p *console.ConsoleTabs) { p.Tabs[1].Href = "#a" }},
		{Want: "console.ConsoleTabs.Tabs[1] needs a Label", Change: func(p *console.ConsoleTabs) { p.Tabs[1].Label = "" }},
		{Want: "console.ConsoleTabs.Tabs[1] needs an Href", Change: func(p *console.ConsoleTabs) { p.Tabs[1].Href = "" }},
	})
}

func segmentedControl() console.ConsoleSegmentedControl {
	return console.ConsoleSegmentedControl{AriaLabel: "About", Name: "about", Segments: []console.ConsoleSegment{
		{Value: "picture", Label: "The picture"},
		{Value: "line", Label: "The line"},
	}, Value: "picture"}
}

func TestSegmentedControlRendersTheExample(t *testing.T) {
	example.Check(t, segmentedControl(), "console/console-tabs/examples/radios.html")
}

func TestSegmentedControlPanics(t *testing.T) {
	example.Panics(t, segmentedControl, []example.Break[console.ConsoleSegmentedControl]{
		{Want: "console.ConsoleSegmentedControl needs an AriaLabel", Change: func(p *console.ConsoleSegmentedControl) { p.AriaLabel = "" }},
		{Want: "console.ConsoleSegmentedControl needs a Name", Change: func(p *console.ConsoleSegmentedControl) { p.Name = "" }},
		{Want: "console.ConsoleSegmentedControl needs two Segments at least: one is no choice", Change: func(p *console.ConsoleSegmentedControl) { p.Segments = p.Segments[:1] }},
		{Want: "console.ConsoleSegmentedControl.Segments[1] needs a Label", Change: func(p *console.ConsoleSegmentedControl) { p.Segments[1].Label = "" }},
		{Want: "console.ConsoleSegmentedControl needs a Value that is one of its Segments' Value: a radio group starts with its value", Change: func(p *console.ConsoleSegmentedControl) { p.Value = "" }},
		{Want: "console.ConsoleSegmentedControl needs a Value that is one of its Segments' Value: a radio group starts with its value", Change: func(p *console.ConsoleSegmentedControl) { p.Segments[1].Value = "picture" }},
	})
}
