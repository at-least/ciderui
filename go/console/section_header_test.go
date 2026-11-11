package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestSectionHeaderRendersTheExamples(t *testing.T) {
	none := console.ConsoleContentBack{Href: "", Label: ""}
	page := h.Div(h.Class("cui-console-page max-w-4xl"),
		console.PageHeader(none, "Issues", "New and regressed issues in this project.", console.ConsoleTabs{AriaLabel: "Time window", Tabs: []console.ConsoleTab{
			{Label: "24h", Href: "#24h", Current: false},
			{Label: "7d", Href: "#7d", Current: true},
			{Label: "30d", Href: "#30d", Current: false},
		}}),
		h.P(g.Text("The page’s content.")),
	)
	example.Check(t, page, "console/console-section-header/examples/section-header.html")
	section := h.Div(h.Class("cui-console-page max-w-4xl"),
		console.SectionHeader("Alerts", "Who hears about a new issue, and how.", nil),
		h.P(g.Text("The section’s content.")),
	)
	example.Check(t, section, "console/console-section-header/examples/section.html")
	back := h.Div(h.Class("cui-console-page max-w-4xl"),
		console.PageHeader(console.ConsoleContentBack{Href: "#needs-you", Label: "Needs You"}, "Checkout times out", "", nil),
	)
	example.Check(t, back, "console/console-section-header/examples/back.html")
}

func TestSectionHeaderWithNoDetailWritesNoLine(t *testing.T) {
	got := example.Render(console.SectionHeader("Alerts", "", nil))
	if want := `<header class="cui-console-section-header"><div><h2 class="cui-console-h4-emph">Alerts</h2></div></header>`; got != want {
		t.Errorf("a header with no Detail wrote more than its title:\n got: %s\nwant: %s", got, want)
	}
}
