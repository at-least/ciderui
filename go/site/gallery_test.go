package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestGalleryRendersTheExample(t *testing.T) {
	card := func(href, headline, price string) site.Card {
		return site.Card{Dark: false, Href: href, Media: nil, Eyebrow: "", Tag: false, Headline: headline, Description: "", Price: price, Class: ""}
	}
	example.Check(t, site.Gallery{ID: "gallery", Paddles: true, Dots: 4, Current: 1, Class: "", Children: g.Group{
		h.Li(h.ID("gallery-1"), card("/phone/", "Cider Phone 3", "From $1199")),
		h.Li(h.ID("gallery-2"), card("/phone-fold/", "Cider Phone Fold", "From $899")),
		h.Li(h.ID("gallery-3"), card("/watch/", "Cider Watch", "From $399")),
		h.Li(h.ID("gallery-4"), card("/buds/", "Cider Buds", "From $129")),
	}}, "site/gallery/examples/gallery.html")
}

func TestGalleryNoPaddlesNoDots(t *testing.T) {
	got := example.Render(site.Gallery{ID: "g", Paddles: false, Dots: 0, Current: 0, Class: "cui-dark", Children: g.Text("x")})
	want := `<div class="cui-gallery cui-dark"><ul class="cui-gallery-scroller" id="g">x</ul></div>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestGalleryAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "needs an ID", func() {
		example.Render(site.Gallery{ID: "", Paddles: false, Dots: 0, Current: 0, Class: "", Children: g.Text("x")})
	})
}

func TestGalleryACurrentDotOutsidePanics(t *testing.T) {
	mustPanic(t, "outside its Dots", func() {
		example.Render(site.Gallery{ID: "g", Paddles: false, Dots: 3, Current: 4, Class: "", Children: g.Text("x")})
	})
}
