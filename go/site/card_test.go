package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	h "maragu.dev/gomponents/html"
)

func TestCardRendersTheExamples(t *testing.T) {
	example.Check(t, site.Card{Dark: false, Href: "/phone/", Media: nil, Eyebrow: "NEW", Tag: true, Headline: "Cider Phone 3", Description: "Our most capable phone.", Price: "From $1199", Class: ""}, "site/card/examples/card.html")
	example.Check(t, site.Card{Dark: false, Href: "/shop/specialist/", Media: nil, Eyebrow: "EXPERT HELP", Tag: false, Headline: "Talk it through with an expert.", Description: "", Price: "", Class: ""}, "site/card/examples/eyebrow.html")
}

func TestCardWritesItsPictureAndThePagesUtilities(t *testing.T) {
	media := h.Img(h.Class("cui-card-media"), h.Src("/phone.jpg"), h.Alt(""))
	got := example.Render(site.Card{Dark: false, Href: "/phone/", Media: media, Eyebrow: "", Tag: false, Headline: "Cider Phone", Description: "", Price: "", Class: "w-full"})
	want := `<div class="cui-card w-full"><img class="cui-card-media" src="/phone.jpg" alt=""><div class="cui-card-content"><h3 class="cui-card-headline"><a class="cui-card-link" href="/phone/">Cider Phone</a></h3></div></div>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestCardAnEmptyHeadlinePanics(t *testing.T) {
	mustPanic(t, "needs a Headline", func() {
		example.Render(site.Card{Dark: false, Href: "/x/", Media: nil, Eyebrow: "", Tag: false, Headline: "", Description: "", Price: "", Class: ""})
	})
}

func TestCardATagWithNoEyebrowPanics(t *testing.T) {
	mustPanic(t, "Tag needs an Eyebrow", func() {
		example.Render(site.Card{Dark: false, Href: "/x/", Media: nil, Eyebrow: "", Tag: true, Headline: "x", Description: "", Price: "", Class: ""})
	})
}

func TestCardDarkThenThePagesUtilities(t *testing.T) {
	got := example.Render(site.Card{Dark: true, Href: "/x/", Media: nil, Eyebrow: "", Tag: false, Headline: "x", Description: "", Price: "", Class: "w-full"})
	if !strings.HasPrefix(got, `<div class="cui-card cui-dark w-full">`) {
		t.Errorf("got %s", got)
	}
}
