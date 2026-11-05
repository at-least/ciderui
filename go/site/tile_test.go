package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestTileRendersTheExamples(t *testing.T) {
	example.Check(t, site.Tile{Large: false, Rounded: false, Link: &site.Link{Label: "Upgrade Program", Href: "/upgrade/"}, Media: nil, Headline: "Upgrade Program", Subhead: "Get the latest phone every year.", Callout: "", Class: "",
		Children: h.A(h.Class("cui-btn"), h.Href("/upgrade/"), g.Text("Learn more")),
	}, "site/tile/examples/tile.html")
	example.Check(t, site.Tile{Large: true, Rounded: false, Link: &site.Link{Label: "Cider Phone Fold", Href: "/phone-fold/"}, Media: nil, Headline: "Cider Phone Fold", Subhead: "Two screens, one phone.", Callout: "Ships October 30", Class: "",
		Children: g.Group{
			h.A(h.Class("cui-btn"), h.Data("size", "cta"), h.Href("/phone-fold/"), g.Text("Learn more")),
			h.A(h.Class("cui-btn"), h.Data("size", "cta"), h.Data("variant", "secondary"), h.Href("/buy/"), g.Text("Pre-order")),
		},
	}, "site/tile/examples/hero.html")
}

func TestTilePanics(t *testing.T) {
	example.Panics(t, func() site.Tile {
		return site.Tile{Large: false, Rounded: false, Link: nil, Media: nil, Headline: "x", Subhead: "", Callout: "", Class: "", Children: g.Text("x")}
	}, []example.Break[site.Tile]{
		{Want: "site.Tile needs a Headline", Change: func(p *site.Tile) { p.Headline = "" }},
		{Want: "site.Tile's Link needs its Label and Href", Change: func(p *site.Tile) { p.Link = &site.Link{Label: "", Href: "/x/"} }},
		{Want: "site.Tile's Link needs its Label and Href", Change: func(p *site.Tile) { p.Link = &site.Link{Label: "x", Href: ""} }},
	})
}

func TestTileRoundedWithAPictureAndNoSubhead(t *testing.T) {
	media := h.Img(h.Src("/a.png"), h.Alt(""))
	got := example.Render(site.Tile{Large: false, Rounded: true, Link: nil, Media: media, Headline: "x", Subhead: "", Callout: "", Class: "", Children: g.Text("y")})
	want := `<div class="cui-tile cui-tile-rounded"><img src="/a.png" alt=""><div class="cui-tile-content"><h2 class="cui-tile-headline">x</h2><div class="cui-tile-ctas">y</div></div></div>`
	if got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
}
