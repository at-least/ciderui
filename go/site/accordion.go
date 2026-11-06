package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
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
	Open     bool
	Class    string
	Children g.Node
}

func (i AccordionItem) node() g.Node {
	if i.Question == "" {
		panic("site.AccordionItem needs a Question")
	}
	return h.Details(class.Root("cui-accordion-item", i.Class), g.If(i.Open, h.Open()),
		h.Summary(h.Class("cui-accordion-question"), g.Text(i.Question), glyph.Disclosure),
		h.Div(h.Class("cui-accordion-answer"), i.Children),
	)
}
