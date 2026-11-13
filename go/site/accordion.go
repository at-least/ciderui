package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Accordion writes no element of its own: its items are the page's to place.
type Accordion struct {
	Items []AccordionItem
}

func (a Accordion) Render(w io.Writer) error {
	return g.Map(a.Items, AccordionItem.node).Render(w)
}

type AccordionItem struct {
	Question string
	Children g.Node
}

func (i AccordionItem) node() g.Node {
	need("site.AccordionItem", "a Question", i.Question)
	return h.Details(h.Class("cui-accordion-item"),
		h.Summary(h.Class("cui-accordion-question"), g.Text(i.Question), glyph.Disclosure),
		h.Div(h.Class("cui-accordion-answer"), i.Children),
	)
}
