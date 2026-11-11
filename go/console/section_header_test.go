package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestSectionHeaderRendersTheExamples(t *testing.T) {
	page := h.Div(h.Class("cui-console-page max-w-4xl"),
		console.ConsoleSectionHeader{
			Level:  console.ConsoleSectionHeaderLevelPage,
			Back:   console.ConsoleSectionHeaderBack{Href: "", Label: ""},
			Title:  "Issues",
			Detail: "New and regressed issues in this project.",
			Children: console.ConsoleTabs{AriaLabel: "Time window", Tabs: []console.ConsoleTab{
				{Label: "24h", Href: "#24h", Current: false},
				{Label: "7d", Href: "#7d", Current: true},
				{Label: "30d", Href: "#30d", Current: false},
			}},
		},
		h.P(g.Text("The page’s content.")),
	)
	example.Check(t, page, "console/console-section-header/examples/section-header.html")
	section := h.Div(h.Class("cui-console-page max-w-4xl"),
		console.ConsoleSectionHeader{Level: console.ConsoleSectionHeaderLevelSection, Back: console.ConsoleSectionHeaderBack{Href: "", Label: ""}, Title: "Alerts", Detail: "Who hears about a new issue, and how.", Children: nil},
		h.P(g.Text("The section’s content.")),
	)
	example.Check(t, section, "console/console-section-header/examples/section.html")
	back := h.Div(h.Class("cui-console-page max-w-4xl"),
		console.ConsoleSectionHeader{Level: console.ConsoleSectionHeaderLevelPage, Back: console.ConsoleSectionHeaderBack{Href: "#needs-you", Label: "Needs You"}, Title: "Checkout times out", Detail: "", Children: nil},
	)
	example.Check(t, back, "console/console-section-header/examples/back.html")
}

func TestSectionHeaderWithNoDetailWritesNoLine(t *testing.T) {
	got := example.Render(console.ConsoleSectionHeader{Level: console.ConsoleSectionHeaderLevelSection, Back: console.ConsoleSectionHeaderBack{Href: "", Label: ""}, Title: "Alerts", Detail: "", Children: nil})
	if want := `<header class="cui-console-section-header"><div><h2 class="cui-console-h4-emph">Alerts</h2></div></header>`; got != want {
		t.Errorf("a header with no Detail wrote more than its title:\n got: %s\nwant: %s", got, want)
	}
}

func TestSectionHeaderPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleSectionHeader {
		return console.ConsoleSectionHeader{Level: console.ConsoleSectionHeaderLevelPage, Back: console.ConsoleSectionHeaderBack{Href: "", Label: ""}, Title: "x", Detail: "", Children: nil}
	}, []example.Break[console.ConsoleSectionHeader]{
		{Want: "console.ConsoleSectionHeader needs a Title", Change: func(p *console.ConsoleSectionHeader) { p.Title = "" }},
		{Want: "console.ConsoleSectionHeader needs a Back with both an Href and a Label, or neither", Change: func(p *console.ConsoleSectionHeader) { p.Back.Href = "/" }},
		{Want: "console.ConsoleSectionHeader needs a Back with both an Href and a Label, or neither", Change: func(p *console.ConsoleSectionHeader) { p.Back.Label = "Home" }},
		{Want: "console.ConsoleSectionHeader's Back is a page's: a section is on the page it belongs to", Change: func(p *console.ConsoleSectionHeader) {
			p.Level, p.Back = console.ConsoleSectionHeaderLevelSection, console.ConsoleSectionHeaderBack{Href: "/", Label: "Home"}
		}},
		{Want: "console.ConsoleSectionHeader needs a Level: ConsoleSectionHeaderLevelPage (a whole page) or ConsoleSectionHeaderLevelSection (one of its sections)", Change: func(p *console.ConsoleSectionHeader) { p.Level = console.ConsoleSectionHeaderLevel{} }},
	})
}
