package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// TileSize's zero value: a promo tile, two to a row.
type TileSize struct{ value string }

// TileSizeLarge: a full-width tile.
var TileSizeLarge = TileSize{"lg"}

type Tile struct {
	Size    TileSize
	Rounded bool
	Dark    bool
	Link    *Link
	// Media: an <img> carrying cui-tile-media.
	Media    g.Node
	Headline string
	Subhead  string
	Callout  string
	Class    string
	// Children: cui-btn buttons, data-size="cta" in a [TileSizeLarge] tile.
	Children g.Node
}

func (t Tile) Render(w io.Writer) error {
	if t.Headline == "" {
		panic("site.Tile needs a Headline")
	}
	block := "cui-tile"
	if t.Rounded {
		block += " cui-tile-rounded"
	}
	if t.Dark {
		block += " cui-dark"
	}
	return h.Div(class.Root(block, t.Class), g.If(t.Size.value != "", h.Data("size", t.Size.value)),
		t.Media,
		g.Iff(t.Link != nil, func() g.Node {
			return h.A(h.Class("cui-tile-link"), h.Href(t.Link.Href), h.Aria("label", t.Link.Text))
		}),
		h.Div(h.Class("cui-tile-content"),
			h.H2(h.Class("cui-tile-headline"), g.Text(t.Headline)),
			g.If(t.Subhead != "", h.P(h.Class("cui-tile-subhead"), g.Text(t.Subhead))),
			g.If(t.Callout != "", h.P(h.Class("cui-tile-callout"), g.Text(t.Callout))),
			h.Div(h.Class("cui-tile-ctas"), t.Children),
		),
	).Render(w)
}
