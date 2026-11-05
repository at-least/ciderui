package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Card struct {
	Href string
	// Media: an <img> carrying cui-card-media.
	Media       g.Node
	Eyebrow     string
	Headline    string
	Description string
	Price       string
	Class       string
}

func (c Card) Render(w io.Writer) error {
	if c.Href == "" {
		panic("site.Card needs an Href")
	}
	if c.Headline == "" {
		panic("site.Card needs a Headline")
	}
	return h.Div(class.Root("cui-card", c.Class),
		c.Media,
		h.Div(h.Class("cui-card-content"),
			g.If(c.Eyebrow != "", h.P(h.Class("cui-card-eyebrow"), g.Text(c.Eyebrow))),
			h.H3(h.Class("cui-card-headline"), h.A(h.Class("cui-card-link"), h.Href(c.Href), g.Text(c.Headline))),
			g.If(c.Description != "", h.P(h.Class("cui-card-description"), g.Text(c.Description))),
			g.If(c.Price != "", h.P(h.Class("cui-card-price"), g.Text(c.Price))),
		),
	).Render(w)
}
