package site

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Accordion is cui-accordion (components/site/accordion): apple.com's FAQ,
// questions under rules, each opening its answer below it with no script
// (<details>), its chevron turning up; with Large, the bag's large question.
type Accordion struct {
	// Large is whether its questions are large (data-size="lg", the bag's
	// outer one).
	Large bool
	// Class is the page's utilities on the accordion; "" for none.
	Class string
	// Children are its items, [AccordionItem]s.
	Children g.Node
}

// Render writes the accordion.
func (a Accordion) Render(w io.Writer) error {
	return h.Div(class.Root("cui-accordion", a.Class), g.If(a.Large, h.Data("size", "lg")), a.Children).Render(w)
}

// AccordionItem is a question of an [Accordion] and its answer, a <details>.
type AccordionItem struct {
	// Question is the question. Required: an empty one panics.
	Question string
	// Open is whether its answer is open.
	Open bool
	// Class is the page's utilities on its root; "" for none.
	Class string
	// Children is the answer.
	Children g.Node
}

// Render writes the item.
func (i AccordionItem) Render(w io.Writer) error {
	if i.Question == "" {
		panic("site.AccordionItem needs a Question")
	}
	return h.Details(class.Root("cui-accordion-item", i.Class), g.If(i.Open, h.Open()),
		h.Summary(h.Class("cui-accordion-question"), g.Text(i.Question), icon.Disclosure),
		h.Div(h.Class("cui-accordion-answer"), i.Children),
	).Render(w)
}
