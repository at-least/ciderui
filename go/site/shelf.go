package site

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Shelf is cui-shelf (components/site/shelf): the store's row of cards, a
// header (a headline and its secondary text after it) over a scroller across
// the window, its items snapping into place; with Paddles, a large paddle at
// each side, the previous disabled (the row at its start): the page's script
// scrolls the row and disables the one at an end.
type Shelf struct {
	// ID is the scroller's id, the paddles' aria-controls. Required: an empty
	// one panics.
	ID string
	// Headline is the headline, an <h2>. Required: an empty one panics.
	Headline string
	// Subhead is the secondary text after it; "" for none.
	Subhead string
	// Paddles is whether it shows its paddles (a page with a script to
	// scroll the row).
	Paddles bool
	// Class is the page's utilities on the shelf; "" for none.
	Class string
	// Children are its items, each an <li> (a [Card] in it).
	Children g.Node
}

// Render writes the shelf.
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
				h.Button(h.Class("cui-gallery-paddle cui-shelf-paddle"), h.Data("size", "lg"), h.Type("button"), h.Aria("controls", s.ID), h.Aria("label", "Previous"), h.Disabled(), icon.PaddlePrevious),
				h.Button(h.Class("cui-gallery-paddle cui-shelf-paddle"), h.Data("size", "lg"), h.Type("button"), h.Aria("controls", s.ID), h.Aria("label", "Next"), icon.PaddleNext),
			}),
		),
	).Render(w)
}
