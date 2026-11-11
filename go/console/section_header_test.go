package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// @invariant: the section-header examples show a header alone, so each is cut
// from the part that writes it, between that part's own markup around it
func headerOf(t *testing.T, part g.Node, before, after string) g.Node {
	t.Helper()
	got := example.Render(part)
	inner, foundBefore := strings.CutPrefix(got, before)
	inner, foundAfter := strings.CutSuffix(inner, after)
	if !foundBefore || !foundAfter {
		t.Fatalf("the header is not the only thing between %s and %s: %s", before, after, got)
	}
	return g.Raw(inner)
}

const (
	contentOpen   = `<div class="cui-console-content">`
	contentMain   = `<main id="main" tabindex="-1" class="cui-console-content-main">`
	contentBefore = contentOpen + contentMain
	contentAfter  = `</main></div>`
)

func TestContentAndSectionWriteTheSectionHeaderExamples(t *testing.T) {
	none := console.ConsoleContentBack{Href: "", Label: ""}
	tabs := console.ConsoleTabs{AriaLabel: "Time window", Tabs: []console.ConsoleTab{
		{Label: "24h", Href: "#24h", Current: false},
		{Label: "7d", Href: "#7d", Current: true},
		{Label: "30d", Href: "#30d", Current: false},
	}}
	page := headerOf(t, console.ConsoleContent{Sidebar: nil, Back: none, Title: "Issues", Detail: "New and regressed issues in this project.", Controls: tabs, Children: nil}, contentBefore, contentAfter)
	example.Check(t, h.Div(h.Class("cui-console-page max-w-4xl"), page, h.P(g.Text("The page’s content."))), "console/console-section-header/examples/section-header.html")
	section := headerOf(t, console.ConsoleSection{Title: "Alerts", Detail: "Who hears about a new issue, and how.", Controls: nil, Children: nil}, `<section class="cui-console-section">`, `</section>`)
	example.Check(t, h.Div(h.Class("cui-console-page max-w-4xl"), section, h.P(g.Text("The section’s content."))), "console/console-section-header/examples/section.html")
	back := headerOf(t, console.ConsoleContent{Sidebar: nil, Back: console.ConsoleContentBack{Href: "#needs-you", Label: "Needs You"}, Title: "Checkout times out", Detail: "", Controls: nil, Children: nil}, contentBefore, contentAfter)
	example.Check(t, h.Div(h.Class("cui-console-page max-w-4xl"), back), "console/console-section-header/examples/back.html")
}
