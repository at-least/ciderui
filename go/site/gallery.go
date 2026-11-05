package site

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// GalleryItem: DotAriaLabel, what its dot reads as; "" in a gallery with no dots.
type GalleryItem struct {
	Content      g.Node
	DotAriaLabel string
}

// The page's script scrolls the row from the paddles and moves the dots'
// aria-current: Gallery renders only the start state.
type Gallery struct {
	ID    string
	Items []GalleryItem
	// Paddles: nil for none.
	Paddles *Paddles
	// Current: the shown item's dot, from 1; 0 for no dots.
	Current int
	Class   string
}

func (x Gallery) Render(w io.Writer) error {
	if x.ID == "" {
		panic("site.Gallery needs an ID")
	}
	if x.Current < 0 || x.Current > len(x.Items) {
		panic("site.Gallery's Current dot is outside its Items")
	}
	var items, dots g.Group
	for i, item := range x.Items {
		id := x.ID + "-" + strconv.Itoa(i+1)
		items = append(items, h.Li(h.ID(id), item.Content))
		if x.Current == 0 {
			continue
		}
		if item.DotAriaLabel == "" {
			panic("site.GalleryItem needs a DotAriaLabel in a gallery with dots")
		}
		dots = append(dots, h.Li(h.A(h.Class("cui-gallery-dot"), h.Href("#"+id), h.Aria("label", item.DotAriaLabel), g.If(i+1 == x.Current, h.Aria("current", "true")))))
	}
	return h.Div(class.Root("cui-gallery", x.Class),
		h.Ul(h.Class("cui-gallery-scroller"), h.ID(x.ID), items),
		g.If(x.Paddles != nil, h.Div(h.Class("cui-gallery-nav"), x.Paddles.node("site.Gallery", "cui-gallery-paddle", x.ID))),
		g.If(x.Current > 0, h.Ul(h.Class("cui-gallery-dots"), dots)),
	).Render(w)
}
