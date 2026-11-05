package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Shelf struct {
	ID       string
	Headline string
	Subhead  string
	// Paddles: nil for none.
	Paddles *Paddles
	Class   string
	// Items: each a [Card].
	Items []g.Node
}

func (s Shelf) Render(w io.Writer) error {
	if s.ID == "" {
		panic("site.Shelf needs an ID")
	}
	if s.Headline == "" {
		panic("site.Shelf needs a Headline")
	}
	return h.Div(class.Root("cui-shelf", s.Class),
		h.Div(h.Class("cui-shelf-header"),
			h.H2(h.Class("cui-shelf-headline"), g.Text(s.Headline)),
			g.If(s.Subhead != "", g.Group{g.Text(" "), h.Span(h.Class("cui-shelf-subhead"), g.Text(s.Subhead))}),
		),
		h.Div(h.Class("cui-shelf-row"),
			h.Ul(h.Class("cui-shelf-scroller"), h.ID(s.ID), g.Map(s.Items, func(item g.Node) g.Node { return h.Li(item) })),
			s.Paddles.node("site.Shelf", "cui-gallery-paddle cui-shelf-paddle", s.ID),
		),
	).Render(w)
}
