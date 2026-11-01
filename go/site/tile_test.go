package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestTileRendersTheExamples(t *testing.T) {
	example.Check(t, site.Tile{Size: site.TileSize{}, Rounded: false, Dark: false, Link: &site.Link{Text: "Upgrade Program", Href: "/upgrade/"}, Media: nil, Headline: "Upgrade Program", Subhead: "Get the latest phone every year.", Callout: "", Class: "",
		Children: h.A(h.Class("cui-btn"), h.Href("/upgrade/"), g.Text("Learn more")),
	}, "site/tile/examples/tile.html")
	example.Check(t, site.Tile{Size: site.TileSizeLarge, Rounded: false, Dark: false, Link: &site.Link{Text: "Cider Phone Fold", Href: "/phone-fold/"}, Media: nil, Headline: "Cider Phone Fold", Subhead: "Two screens, one phone.", Callout: "Ships October 30", Class: "",
		Children: g.Group{
			h.A(h.Class("cui-btn"), h.Data("size", "cta"), h.Href("/phone-fold/"), g.Text("Learn more")),
			h.A(h.Class("cui-btn"), h.Data("size", "cta"), h.Data("variant", "secondary"), h.Href("/buy/"), g.Text("Pre-order")),
		},
	}, "site/tile/examples/hero.html")
}

func TestTileAnEmptyHeadlinePanics(t *testing.T) {
	mustPanic(t, "needs a Headline", func() {
		example.Render(site.Tile{Size: site.TileSize{}, Rounded: true, Dark: false, Link: nil, Media: nil, Headline: "", Subhead: "", Callout: "", Class: "", Children: g.Text("x")})
	})
}

func TestTileTakesThePagesUtilities(t *testing.T) {
	got := example.Render(site.Tile{Size: site.TileSize{}, Rounded: true, Dark: false, Link: nil, Media: nil, Headline: "x", Subhead: "", Callout: "", Class: "min-h-0", Children: g.Text("x")})
	if !strings.HasPrefix(got, `<div class="cui-tile cui-tile-rounded min-h-0">`) {
		t.Errorf("got %s", got)
	}
}

func TestTileRoundedDarkThenThePagesUtilities(t *testing.T) {
	got := example.Render(site.Tile{Size: site.TileSize{}, Rounded: true, Dark: true, Link: nil, Media: nil, Headline: "x", Subhead: "", Callout: "", Class: "min-h-0", Children: g.Text("x")})
	if !strings.HasPrefix(got, `<div class="cui-tile cui-tile-rounded cui-dark min-h-0">`) {
		t.Errorf("got %s", got)
	}
}
