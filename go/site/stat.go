package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Stat is cui-stat (components/site/stat): a product page's number, a caption
// above it and one below in the secondary text; with Large, AirPods Pro's.
type Stat struct {
	// Large is whether it is a large one (data-size="lg").
	Large bool
	// Before is the caption above the value; "" for none.
	Before string
	// Value is the value. Required: an empty one panics.
	Value string
	// After is the caption below the value; "" for none.
	After string
	// Class is the page's utilities on the stat; "" for none.
	Class string
}

// Render writes the stat.
func (s Stat) Render(w io.Writer) error {
	if s.Value == "" {
		panic("site.Stat needs a Value")
	}
	return h.P(class.Root("cui-stat", s.Class), g.If(s.Large, h.Data("size", "lg")),
		g.If(s.Before != "", h.Span(h.Class("cui-stat-caption"), g.Text(s.Before))),
		h.Span(h.Class("cui-stat-value"), g.Text(s.Value)),
		g.If(s.After != "", h.Span(h.Class("cui-stat-caption"), g.Text(s.After))),
	).Render(w)
}
