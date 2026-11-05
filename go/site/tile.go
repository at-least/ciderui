package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Tile struct {
	// Large: a full-width tile; else a promo tile, two to a row.
	Large   bool
	Rounded bool
	Link    *Link
	// Media: an <img> carrying cui-tile-media.
	Media    g.Node
	Headline string
	Subhead  string
	Callout  string
	Class    string
	// Children: cui-btn buttons, data-size="cta" in a Large tile.
	Children g.Node
}

func (t Tile) Render(w io.Writer) error {
	if t.Headline == "" {
		panic("site.Tile needs a Headline")
	}
	if t.Link != nil && (t.Link.Label == "" || t.Link.Href == "") {
		panic("site.Tile's Link needs its Label and Href")
	}
	block := "cui-tile"
	if t.Rounded {
		block += " cui-tile-rounded"
	}
	return h.Div(class.Root(block, t.Class), g.If(t.Large, h.Data("size", "lg")),
		t.Media,
		g.Iff(t.Link != nil, func() g.Node {
			return h.A(h.Class("cui-tile-link"), h.Href(t.Link.Href), h.Aria("label", t.Link.Label))
		}),
		h.Div(h.Class("cui-tile-content"),
			h.H2(h.Class("cui-tile-headline"), g.Text(t.Headline)),
			g.If(t.Subhead != "", h.P(h.Class("cui-tile-subhead"), g.Text(t.Subhead))),
			g.If(t.Callout != "", h.P(h.Class("cui-tile-callout"), g.Text(t.Callout))),
			h.Div(h.Class("cui-tile-ctas"), t.Children),
		),
	).Render(w)
}
