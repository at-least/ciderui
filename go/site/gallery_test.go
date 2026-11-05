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

func TestGalleryAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "needs an ID", func() {
		example.Render(site.Gallery{ID: "", Items: nil, Paddles: nil, Current: 0, Class: ""})
	})
}

func TestGalleryACurrentDotOutsidePanics(t *testing.T) {
	for _, current := range []int{-1, 2} {
		mustPanic(t, "outside its Items", func() {
			example.Render(site.Gallery{ID: "g", Items: []site.GalleryItem{{Content: g.Text("x"), DotAriaLabel: "a"}}, Paddles: nil, Current: current, Class: ""})
		})
	}
}

func TestGalleryAnEmptyDotAriaLabelPanics(t *testing.T) {
	mustPanic(t, "site.GalleryItem needs a DotAriaLabel in a gallery with dots", func() {
		example.Render(site.Gallery{ID: "g", Items: []site.GalleryItem{{Content: g.Text("x"), DotAriaLabel: ""}}, Paddles: nil, Current: 1, Class: ""})
	})
}

func TestGalleryPaddlesNeedTheirWords(t *testing.T) {
	mustPanic(t, "site.Gallery's Paddles need a PreviousAriaLabel", func() {
		example.Render(site.Gallery{ID: "g", Items: nil, Paddles: &site.Paddles{PreviousAriaLabel: "", NextAriaLabel: "x"}, Current: 0, Class: ""})
	})
	mustPanic(t, "site.Gallery's Paddles need a NextAriaLabel", func() {
		example.Render(site.Gallery{ID: "g", Items: nil, Paddles: &site.Paddles{PreviousAriaLabel: "x", NextAriaLabel: ""}, Current: 0, Class: ""})
	})
}
