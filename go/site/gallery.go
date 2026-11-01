package site

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Gallery is cui-gallery (components/site/gallery): a row of items across the
// window, each snapping into place where the content starts; with Paddles,
// its paddles' row, the previous disabled (the row at its start); with Dots,
// a dot for each item, a link to it, the current one marked. The page's
// script scrolls the row from the paddles, disables the one at an end and
// moves the mark: the gallery shows the state the page sets.
type Gallery struct {
	// ID is the row's id: the paddles' aria-controls, and the dots' links to
	// its items, #<id>-1, #<id>-2, …. Required: an empty one panics.
	ID string
	// Paddles is whether it shows its paddles (a page with a script to
	// scroll the row).
	Paddles bool
	// Dots is how many dots, one for each item (0 for none).
	Dots int
	// Current is the current dot, counted from 1: with dots, one outside
	// them panics.
	Current int
	// Class is the page's utilities on the gallery (cui-dark); "" for none.
	Class string
	// Children are its items, each an <li>, with the id its dot links to.
	Children g.Node
}

// Render writes the gallery.
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
			h.Button(h.Class("cui-gallery-paddle"), h.Type("button"), h.Aria("controls", x.ID), h.Aria("label", "Previous"), h.Disabled(), icon.PaddlePrevious),
			h.Button(h.Class("cui-gallery-paddle"), h.Type("button"), h.Aria("controls", x.ID), h.Aria("label", "Next"), icon.PaddleNext),
		)),
		g.If(x.Dots > 0, h.Ul(h.Class("cui-gallery-dots"), dots)),
	).Render(w)
}
