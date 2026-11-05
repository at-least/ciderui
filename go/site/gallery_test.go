package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
)

func TestGalleryRendersTheExample(t *testing.T) {
	card := func(href, headline, price, dot string) site.GalleryItem {
		return site.GalleryItem{Content: site.Card{Href: href, Media: nil, Eyebrow: "", Headline: headline, Description: "", Price: price, Class: ""}, DotAriaLabel: dot}
	}
	example.Check(t, site.Gallery{ID: "gallery", Items: []site.GalleryItem{
		card("/phone/", "Cider Phone 3", "From $1199", "Item 1"),
		card("/phone-fold/", "Cider Phone Fold", "From $899", "Item 2"),
		card("/watch/", "Cider Watch", "From $399", "Item 3"),
		card("/buds/", "Cider Buds", "From $129", "Item 4"),
	}, Paddles: &site.Paddles{PreviousAriaLabel: "Previous", NextAriaLabel: "Next"}, Current: 1, Class: ""}, "site/gallery/examples/gallery.html")
}

func TestGalleryNoPaddlesNoDots(t *testing.T) {
	got := example.Render(site.Gallery{ID: "g", Items: []site.GalleryItem{{Content: g.Text("x"), DotAriaLabel: ""}}, Paddles: nil, Current: 0, Class: "cui-dark"})
	want := `<div class="cui-gallery cui-dark"><ul class="cui-gallery-scroller" id="g"><li id="g-1">x</li></ul></div>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestGalleryPanics(t *testing.T) {
	example.Panics(t, func() site.Gallery {
		return site.Gallery{ID: "g", Items: []site.GalleryItem{{Content: g.Text("x"), DotAriaLabel: "a"}}, Paddles: nil, Current: 1, Class: ""}
	}, []example.Break[site.Gallery]{
		{Want: "site.Gallery needs an ID", Change: func(p *site.Gallery) { p.ID = "" }},
		{Want: "site.Gallery's Current dot is outside its Items", Change: func(p *site.Gallery) { p.Current = -1 }},
		{Want: "site.Gallery's Current dot is outside its Items", Change: func(p *site.Gallery) { p.Current = 2 }},
		{Want: "site.GalleryItem needs a DotAriaLabel in a gallery with dots", Change: func(p *site.Gallery) { p.Items[0].DotAriaLabel = "" }},
		{Want: "site.Gallery's Paddles need a PreviousAriaLabel", Change: func(p *site.Gallery) { p.Paddles = &site.Paddles{PreviousAriaLabel: "", NextAriaLabel: "x"} }},
		{Want: "site.Gallery's Paddles need a NextAriaLabel", Change: func(p *site.Gallery) { p.Paddles = &site.Paddles{PreviousAriaLabel: "x", NextAriaLabel: ""} }},
	})
}
