package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
)

func TestShelfRendersTheExamples(t *testing.T) {
	card := func(href, eyebrow, headline, price string) g.Node {
		return site.Card{Href: href, Media: nil, Eyebrow: eyebrow, Headline: headline, Description: "", Price: price, Class: ""}
	}
	for _, c := range []struct {
		paddles *site.Paddles
		file    string
	}{
		{nil, "site/shelf/examples/shelf.html"},
		{&site.Paddles{PreviousAriaLabel: "Previous", NextAriaLabel: "Next"}, "site/shelf/examples/paddles.html"},
	} {
		example.Check(t, site.Shelf{ID: "shelf", Headline: "New arrivals.", Subhead: "See what just came in.", Paddles: c.paddles, Class: "", Items: []g.Node{
			card("/phone/", "NEW", "Cider Phone 3", "From $1199"),
			card("/phone-fold/", "NEW", "Cider Phone Fold", "From $899"),
			card("/watch/", "", "Cider Watch", "From $399"),
			card("/buds/", "", "Cider Buds", "From $129"),
		}}, c.file)
	}
}

func TestShelfNoSubheadNoSpace(t *testing.T) {
	got := example.Render(site.Shelf{ID: "accessories", Headline: "Accessories.", Subhead: "", Paddles: nil, Class: "mt-8", Items: []g.Node{g.Text("x")}})
	want := `<div class="cui-shelf mt-8"><div class="cui-shelf-header"><h2 class="cui-shelf-headline">Accessories.</h2></div><div class="cui-shelf-row"><ul class="cui-shelf-scroller" id="accessories"><li>x</li></ul></div></div>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestShelfPanics(t *testing.T) {
	example.Panics(t, func() site.Shelf {
		return site.Shelf{ID: "x", Headline: "x", Subhead: "", Paddles: nil, Class: "", Items: nil}
	}, []example.Break[site.Shelf]{
		{Want: "site.Shelf needs an ID", Change: func(p *site.Shelf) { p.ID = "" }},
		{Want: "site.Shelf needs a Headline", Change: func(p *site.Shelf) { p.Headline = "" }},
		{Want: "site.Shelf's Paddles need a PreviousAriaLabel", Change: func(p *site.Shelf) { p.Paddles = &site.Paddles{PreviousAriaLabel: "", NextAriaLabel: "x"} }},
		{Want: "site.Shelf's Paddles need a NextAriaLabel", Change: func(p *site.Shelf) { p.Paddles = &site.Paddles{PreviousAriaLabel: "x", NextAriaLabel: ""} }},
	})
}
