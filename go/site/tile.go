package site

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Tiles is apple.com's search results' grid: tiles three across from 735px,
// two below, each a square picture over its title and a line under it.
type Tiles struct {
	Items []Tile
	// Attrs: nodes on the grid's <ul>, a page script's data-* (an order it
	// sets per reader).
	Attrs g.Node
}

// Tile is one of [Tiles]: the whole tile is its title's link.
type Tile struct {
	Href string
	// Image: the picture's URL; its alt is empty, the Title naming the link.
	// "" draws the square in the alternate ground.
	Image  string
	Title  string
	Detail string
	// Attrs: nodes on the tile's <li>, a page script's data-*.
	Attrs g.Node
}

func (t Tiles) Render(w io.Writer) error {
	if len(t.Items) == 0 {
		panic("site.Tiles needs Items: a page with none says so in its own words")
	}
	return h.Ul(h.Class("cui-tiles"), t.Attrs, g.Map(t.Items, Tile.node)).Render(w)
}

func (t Tile) node() g.Node {
	if t.Href == "" || t.Title == "" {
		panic("site.Tile needs an Href and a Title")
	}
	var media g.Node = h.Span(h.Class("cui-tile-media"))
	if t.Image != "" {
		media = h.Img(h.Class("cui-tile-media"), h.Src(t.Image), h.Alt(""), h.Loading("lazy"), g.Attr("decoding", "async"))
	}
	return h.Li(h.Class("cui-tile"), t.Attrs,
		media,
		h.A(h.Class("cui-tile-title"), h.Href(t.Href), g.Text(t.Title)),
		g.If(t.Detail != "", h.P(h.Class("cui-tile-detail"), g.Text(t.Detail))),
	)
}
