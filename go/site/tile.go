package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// TileSize is a [Tile]'s size (data-size). Its zero value is a promo tile,
// two to a row from 735px: 580px, 490px, 500px tall.
type TileSize struct{ value string }

// TileSizeLarge is a full-width tile (data-size="lg"): 692px, 650px, 500px,
// its type a step larger.
var TileSizeLarge = TileSize{"lg"}

// Tile is cui-tile (components/site/tile): a block of its own ground, its
// headline, subhead, callout and buttons centered at its top, its picture
// filling it; with Link, the whole tile a link, its buttons above it.
type Tile struct {
	// Size is its size.
	Size TileSize
	// Rounded is whether its corners are rounded (cui-tile-rounded, 28px).
	Rounded bool
	// Dark is whether it is a dark block (cui-dark): black, its text light.
	Dark bool
	// Link is the whole tile's link, its Text the link's name (aria-label);
	// nil for none.
	Link *Link
	// Media is its picture (cui-tile-media on an <img>); nil for none.
	Media g.Node
	// Headline is the headline, an <h2>. Required: an empty one panics.
	Headline string
	// Subhead is the subhead; "" for none.
	Subhead string
	// Callout is the callout, in the secondary text; "" for none.
	Callout string
	// Class is the page's utilities on the tile (its minimum height removed:
	// min-h-0); "" for none.
	Class string
	// Children are the buttons (cui-btn, data-size="cta" in a large tile).
	Children g.Node
}

// Render writes the tile.
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
