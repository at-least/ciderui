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
}

// Tile is one of [Tiles]: the whole tile is its title's link.
type Tile struct {
	Href string
	// Image: the picture's URL; its alt is empty, the Title naming the link.
	Image  string
	Title  string
	Detail string
}

func (t Tiles) Render(w io.Writer) error {
	if len(t.Items) == 0 {
		panic("site.Tiles needs Items: a page with none says so in its own words")
	}
	return h.Ul(h.Class("cui-tiles"), g.Map(t.Items, Tile.node)).Render(w)
}

func (t Tile) node() g.Node {
	if t.Href == "" || t.Image == "" || t.Title == "" {
		panic("site.Tile needs an Href, an Image and a Title")
	}
	return h.Li(h.Class("cui-tile"),
		h.Img(h.Class("cui-tile-media"), h.Src(t.Image), h.Alt(""), h.Loading("lazy"), g.Attr("decoding", "async")),
		h.A(h.Class("cui-tile-title"), h.Href(t.Href), g.Text(t.Title)),
		g.If(t.Detail != "", h.P(h.Class("cui-tile-detail"), g.Text(t.Detail))),
	)
}
