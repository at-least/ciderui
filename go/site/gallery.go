package site

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// The page's script scrolls the row from the paddles, disables the paddle
// at an end and moves the dots' aria-current: Gallery renders only the
// start state.
type Gallery struct {
	// ID: each item's id must be <ID>-<n>, n from 1, for its dot's link.
	ID      string
	Paddles bool
	// Dots: one per item, or 0.
	Dots    int
	Current int
	Class   string
	// Children: <li>s.
	Children g.Node
}

func (x Gallery) Render(w io.Writer) error {
	if x.ID == "" {
		panic("site.Gallery needs an ID")
	}
	if x.Dots != 0 && (x.Current < 1 || x.Current > x.Dots) {
		panic("site.Gallery's Current dot is outside its Dots")
	}
	var dots g.Group
	for n := 1; n <= x.Dots; n++ {
		dots = append(dots, h.Li(h.A(h.Class("cui-gallery-dot"), h.Href("#"+x.ID+"-"+strconv.Itoa(n)), h.Aria("label", "Item "+strconv.Itoa(n)), g.If(n == x.Current, h.Aria("current", "true")))))
	}
	return h.Div(class.Root("cui-gallery", x.Class),
		h.Ul(h.Class("cui-gallery-scroller"), h.ID(x.ID), x.Children),
		g.If(x.Paddles, h.Div(h.Class("cui-gallery-nav"),
			h.Button(h.Class("cui-gallery-paddle"), h.Type("button"), h.Aria("controls", x.ID), h.Aria("label", "Previous"), h.Disabled(), glyph.PaddlePrevious),
			h.Button(h.Class("cui-gallery-paddle"), h.Type("button"), h.Aria("controls", x.ID), h.Aria("label", "Next"), glyph.PaddleNext),
		)),
		g.If(x.Dots > 0, h.Ul(h.Class("cui-gallery-dots"), dots)),
	).Render(w)
}
