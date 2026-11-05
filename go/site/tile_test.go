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

func TestTileAnEmptyHeadlinePanics(t *testing.T) {
	mustPanic(t, "needs a Headline", func() {
		example.Render(site.Tile{Large: false, Rounded: true, Link: nil, Media: nil, Headline: "", Subhead: "", Callout: "", Class: "", Children: g.Text("x")})
	})
}

func TestTileTakesThePagesUtilities(t *testing.T) {
	got := example.Render(site.Tile{Large: false, Rounded: true, Link: nil, Media: nil, Headline: "x", Subhead: "", Callout: "", Class: "min-h-0", Children: g.Text("x")})
	if !strings.HasPrefix(got, `<div class="cui-tile cui-tile-rounded min-h-0">`) {
		t.Errorf("got %s", got)
	}
}

func TestTileAnEmptyLinkPanics(t *testing.T) {
	for _, link := range []site.Link{{Label: "", Href: "/x/"}, {Label: "x", Href: ""}} {
		mustPanic(t, "site.Tile's Link needs its Label and Href", func() {
			example.Render(site.Tile{Large: false, Rounded: false, Link: &link, Media: nil, Headline: "x", Subhead: "", Callout: "", Class: "", Children: g.Text("x")})
		})
	}
}
