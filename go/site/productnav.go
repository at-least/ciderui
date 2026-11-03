package site

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type ProductNav struct {
	ID    string
	Label string
	// Paddles: only on a page with a script that scrolls the row.
	Paddles bool
	Class   string
	// Children: [ProductNavItem]s.
	Children g.Node
}

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

type ProductNavItem struct {
	Href  string
	Name  string
	Class string
	// Children: an <img> carrying cui-productnav-image, with its width and
	// height: its shape sets the card's width.
	Children g.Node
}

func (i ProductNavItem) Render(w io.Writer) error {
	if i.Name == "" {
		panic("site.ProductNavItem needs a Name")
	}
	return h.Li(g.If(i.Class != "", h.Class(i.Class)),
		h.A(h.Class("cui-productnav-item"), h.Href(i.Href), i.Children, h.Span(h.Class("cui-productnav-name"), g.Text(i.Name))),
	).Render(w)
}
