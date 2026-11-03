package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Card struct {
	Dark bool
	Href string
	// Media: an <img> carrying cui-card-media.
	Media       g.Node
	Eyebrow     string
	Tag         bool
	Headline    string
	Description string
	Price       string
	Class       string
}

func (c Card) Render(w io.Writer) error {
	if c.Headline == "" {
		panic("site.Card needs a Headline")
	}
	if c.Tag && c.Eyebrow == "" {
		panic("site.Card's Tag needs an Eyebrow")
	}
	block := "cui-card"
	if c.Dark {
		block = "cui-card cui-dark"
	}
	return h.Div(class.Root(block, c.Class),
		c.Media,
		h.Div(h.Class("cui-card-content"),
			g.If(c.Eyebrow != "", h.P(h.Class("cui-card-eyebrow"), g.If(c.Tag, h.Data("variant", "tag")), g.Text(c.Eyebrow))),
			h.H3(h.Class("cui-card-headline"), h.A(h.Class("cui-card-link"), h.Href(c.Href), g.Text(c.Headline))),
			g.If(c.Description != "", h.P(h.Class("cui-card-description"), g.Text(c.Description))),
			g.If(c.Price != "", h.P(h.Class("cui-card-price"), g.Text(c.Price))),
		),
	).Render(w)
}
