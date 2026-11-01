package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Card is cui-card (components/site/card): the store's card, its picture
// filling it and its text at its top, an eyebrow (or a tag), its headline, a
// description and a price; the whole card its headline's link.
type Card struct {
	// Dark is whether it is a dark card (cui-dark): black, its text white.
	Dark bool
	// Href is its headline's link, the whole card's.
	Href string
	// Media is its picture (cui-card-media on an <img>); nil for none.
	Media g.Node
	// Eyebrow is the eyebrow above the headline; "" for none.
	Eyebrow string
	// Tag is whether the eyebrow is a tag ("NEW"), in the tag's color. A Tag
	// with no Eyebrow panics.
	Tag bool
	// Headline is the headline, an <h3>, its text the link's. Required: an
	// empty one panics.
	Headline string
	// Description is the description; "" for none.
	Description string
	// Price is the price; "" for none.
	Price string
	// Class is the page's utilities on the card (another width: w-full); ""
	// for none.
	Class string
}

// Render writes the card.
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
