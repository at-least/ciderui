package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestShelfRendersTheExamples(t *testing.T) {
	card := func(href, eyebrow string, tag bool, headline, price string) g.Node {
		return h.Li(site.Card{Dark: false, Href: href, Media: nil, Eyebrow: eyebrow, Tag: tag, Headline: headline, Description: "", Price: price, Class: ""})
	}
	for paddles, file := range map[bool]string{
		false: "site/shelf/examples/shelf.html",
		true:  "site/shelf/examples/paddles.html",
	} {
		example.Check(t, site.Shelf{ID: "shelf", Headline: "New arrivals.", Subhead: "See what just came in.", Paddles: paddles, Class: "", Children: g.Group{
			card("/phone/", "NEW", true, "Cider Phone 3", "From $1199"),
			card("/phone-fold/", "NEW", true, "Cider Phone Fold", "From $899"),
			card("/watch/", "", false, "Cider Watch", "From $399"),
			card("/buds/", "", false, "Cider Buds", "From $129"),
		}}, file)
	}
}

func TestShelfNoSubheadNoSpace(t *testing.T) {
	got := example.Render(site.Shelf{ID: "accessories", Headline: "Accessories.", Subhead: "", Paddles: false, Class: "mt-8", Children: g.Text("x")})
	want := `<div class="cui-shelf mt-8"><div class="cui-shelf-header"><h2 class="cui-shelf-headline">Accessories.</h2></div><div class="cui-shelf-row"><ul class="cui-shelf-scroller" id="accessories">x</ul></div></div>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestShelfAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "needs an ID", func() {
		example.Render(site.Shelf{ID: "", Headline: "x", Subhead: "", Paddles: false, Class: "", Children: g.Text("x")})
	})
}

func TestShelfAnEmptyHeadlinePanics(t *testing.T) {
	mustPanic(t, "needs a Headline", func() {
		example.Render(site.Shelf{ID: "x", Headline: "", Subhead: "", Paddles: false, Class: "", Children: g.Text("x")})
	})
}
