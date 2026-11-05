package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type ProductNav struct {
	ID        string
	AriaLabel string
	// Paddles: nil for none.
	Paddles *Paddles
	Class   string
	Items   []ProductNavItem
}

func (n ProductNav) Render(w io.Writer) error {
	if n.ID == "" {
		panic("site.ProductNav needs an ID")
	}
	if n.AriaLabel == "" {
		panic("site.ProductNav needs an AriaLabel")
	}
	return h.Nav(class.Root("cui-shelf cui-productnav", n.Class), h.Aria("label", n.AriaLabel),
		h.Div(h.Class("cui-shelf-row"),
			h.Ul(h.Class("cui-shelf-scroller"), h.ID(n.ID), g.Map(n.Items, ProductNavItem.node)),
			n.Paddles.node("site.ProductNav", "cui-gallery-paddle cui-shelf-paddle", n.ID),
		),
	).Render(w)
}

type ProductNavItem struct {
	Href  string
	Name  string
	Class string
	// Children: an <img> carrying cui-productnav-image, with its width and
	// height: its shape sets the card's width.
	Children g.Node
}

func (i ProductNavItem) node() g.Node {
	if i.Name == "" {
		panic("site.ProductNavItem needs a Name")
	}
	return h.Li(g.If(i.Class != "", h.Class(i.Class)),
		h.A(h.Class("cui-productnav-item"), h.Href(i.Href), i.Children, h.Span(h.Class("cui-productnav-name"), g.Text(i.Name))),
	)
}
