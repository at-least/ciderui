package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	h "maragu.dev/gomponents/html"
)

func TestCardRendersTheExamples(t *testing.T) {
	example.Check(t, site.Card{Href: "/phone/", Media: nil, Eyebrow: "NEW", Headline: "Cider Phone 3", Description: "Our most capable phone.", Price: "From $1199", Class: ""}, "site/card/examples/card.html")
	example.Check(t, site.Card{Href: "/shop/specialist/", Media: nil, Eyebrow: "EXPERT HELP", Headline: "Talk it through with an expert.", Description: "", Price: "", Class: ""}, "site/card/examples/eyebrow.html")
}

func TestCardWritesItsPicture(t *testing.T) {
	media := h.Img(h.Class("cui-card-media"), h.Src("/phone.jpg"), h.Alt(""))
	got := example.Render(site.Card{Href: "/phone/", Media: media, Eyebrow: "", Headline: "Cider Phone", Description: "", Price: "", Class: ""})
	want := `<div class="cui-card"><img class="cui-card-media" src="/phone.jpg" alt=""><div class="cui-card-content"><h3 class="cui-card-headline"><a class="cui-card-link" href="/phone/">Cider Phone</a></h3></div></div>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestCardPanics(t *testing.T) {
	example.Panics(t, func() site.Card {
		return site.Card{Href: "/x/", Media: nil, Eyebrow: "", Headline: "x", Description: "", Price: "", Class: ""}
	}, []example.Break[site.Card]{
		{Want: "site.Card needs a Headline", Change: func(p *site.Card) { p.Headline = "" }},
		{Want: "site.Card needs an Href", Change: func(p *site.Card) { p.Href = "" }},
	})
}
