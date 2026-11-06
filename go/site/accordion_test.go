package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestAccordionRendersTheExamples(t *testing.T) {
	example.Check(t, site.Accordion{Items: []site.AccordionItem{
		{Question: "How long does delivery take?", Open: true, Class: "", Children: h.P(g.Text("Add your address at checkout to see a delivery date."))},
		{Question: "Can I collect my order in person?", Open: false, Class: "", Children: h.P(g.Text("Yes. Pick a store when you check out."))},
	}}, "site/accordion/examples/accordion.html")
}

func TestAccordionPanics(t *testing.T) {
	example.Panics(t, func() site.Accordion {
		return site.Accordion{Items: []site.AccordionItem{{Question: "x", Open: false, Class: "", Children: g.Text("x")}}}
	}, []example.Break[site.Accordion]{
		{Want: "site.AccordionItem needs a Question", Change: func(p *site.Accordion) { p.Items[0].Question = "" }},
	})
}
