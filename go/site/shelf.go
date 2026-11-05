package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// With Paddles, the page's script scrolls the row and disables the paddle
// at an end: Shelf renders only the start state.
type Shelf struct {
	ID       string
	Headline string
	Subhead  string
	Paddles  bool
	Class    string
	// Children: <li>s, each holding a [Card].
	Children g.Node
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
			h.Ul(h.Class("cui-shelf-scroller"), h.ID(s.ID), s.Children),
			g.If(s.Paddles, g.Group{
				h.Button(h.Class("cui-gallery-paddle cui-shelf-paddle"), h.Data("size", "lg"), h.Type("button"), h.Aria("controls", s.ID), h.Aria("label", "Previous"), h.Disabled(), glyph.PaddlePrevious),
				h.Button(h.Class("cui-gallery-paddle cui-shelf-paddle"), h.Data("size", "lg"), h.Type("button"), h.Aria("controls", s.ID), h.Aria("label", "Next"), glyph.PaddleNext),
			}),
		),
	).Render(w)
}
