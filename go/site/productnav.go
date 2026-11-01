package site

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ProductNav is cui-productnav (components/site/productnav): the store's
// shelf of small cards ([ProductNavItem]s), each a product's image above its
// name, on a cui-shelf with no header; with Paddles, the shelf's large
// paddles.
type ProductNav struct {
	// ID is the scroller's id, the paddles' aria-controls. Required: an empty
	// one panics.
	ID string
	// Label is the nav's name (aria-label). Required: an empty one panics.
	Label string
	// Paddles is whether it shows its paddles (a page with a script to
	// scroll the row).
	Paddles bool
	// Class is the page's utilities on the nav; "" for none.
	Class string
	// Children are its cards, [ProductNavItem]s.
	Children g.Node
}

// Render writes the nav.
func (n ProductNav) Render(w io.Writer) error {
	if n.ID == "" {
		panic("site.ProductNav needs an ID")
	}
	if n.Label == "" {
		panic("site.ProductNav needs a Label")
	}
	return h.Nav(class.Root("cui-shelf cui-productnav", n.Class), h.Aria("label", n.Label),
		h.Div(h.Class("cui-shelf-row"),
			h.Ul(h.Class("cui-shelf-scroller"), h.ID(n.ID), n.Children),
			g.If(n.Paddles, g.Group{
				h.Button(h.Class("cui-gallery-paddle cui-shelf-paddle"), h.Data("size", "lg"), h.Type("button"), h.Aria("controls", n.ID), h.Aria("label", "Previous"), h.Disabled(), icon.PaddlePrevious),
				h.Button(h.Class("cui-gallery-paddle cui-shelf-paddle"), h.Data("size", "lg"), h.Type("button"), h.Aria("controls", n.ID), h.Aria("label", "Next"), icon.PaddleNext),
			}),
		),
	).Render(w)
}

// ProductNavItem is a card of a [ProductNav]: a link, its image above its
// name.
type ProductNavItem struct {
	// Href is the link.
	Href string
	// Name is the product's name. Required: an empty one panics.
	Name string
	// Class is the page's utilities on its root, the <li>; "" for none.
	Class string
	// Children is its image, cui-productnav-image on an <img> (with its
	// width and height: its shape sets its width).
	Children g.Node
}

// Render writes the card.
func (i ProductNavItem) Render(w io.Writer) error {
	if i.Name == "" {
		panic("site.ProductNavItem needs a Name")
	}
	return h.Li(g.If(i.Class != "", h.Class(i.Class)),
		h.A(h.Class("cui-productnav-item"), h.Href(i.Href), i.Children, h.Span(h.Class("cui-productnav-name"), g.Text(i.Name))),
	).Render(w)
}
