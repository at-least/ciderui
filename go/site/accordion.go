package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Accordion struct {
	// Large: the bag page's outer accordion.
	Large bool
	Class string
	// Children: [AccordionItem]s.
	Children g.Node
}

func (a Accordion) Render(w io.Writer) error {
	return h.Div(class.Root("cui-accordion", a.Class), g.If(a.Large, h.Data("size", "lg")), a.Children).Render(w)
}

type AccordionItem struct {
	Question string
	Open     bool
	Class    string
	Children g.Node
}

func (i AccordionItem) Render(w io.Writer) error {
	if i.Question == "" {
		panic("site.AccordionItem needs a Question")
	}
	return h.Details(class.Root("cui-accordion-item", i.Class), g.If(i.Open, h.Open()),
		h.Summary(h.Class("cui-accordion-question"), g.Text(i.Question), glyph.Disclosure),
		h.Div(h.Class("cui-accordion-answer"), i.Children),
	).Render(w)
}
