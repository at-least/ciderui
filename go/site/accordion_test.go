package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestAccordionRendersTheExamples(t *testing.T) {
	example.Check(t, site.Accordion{Large: false, Class: "", Items: []site.AccordionItem{
		{Question: "How long does delivery take?", Open: true, Class: "", Children: h.P(g.Text("Add your address at checkout to see a delivery date."))},
		{Question: "Can I collect my order in person?", Open: false, Class: "", Children: h.P(g.Text("Yes. Pick a store when you check out."))},
	}}, "site/accordion/examples/accordion.html")
	example.Check(t, site.Accordion{Large: true, Class: "", Items: []site.AccordionItem{{Question: "Before You Buy", Open: false, Class: "", Children: h.P(g.Text("What people ask most, answered."))}}}, "site/accordion/examples/large.html")
}

func TestAccordionItemAnEmptyQuestionPanics(t *testing.T) {
	mustPanic(t, "needs a Question", func() {
		example.Render(site.Accordion{Large: false, Class: "", Items: []site.AccordionItem{{Question: "", Open: false, Class: "", Children: g.Text("x")}}})
	})
}

func TestAccordionItemTakesThePagesUtilities(t *testing.T) {
	got := example.Render(site.Accordion{Large: false, Class: "", Items: []site.AccordionItem{{Question: "x", Open: false, Class: "hidden", Children: g.Text("x")}}})
	if !strings.Contains(got, `<details class="cui-accordion-item hidden">`) {
		t.Errorf("got %s", got)
	}
}
