package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func link(label, href string, variant site.BtnVariant, size site.BtnSize) site.Btn {
	return site.Btn{Href: href, Label: label, Submit: false, Variant: variant, Size: size, Block: false, Rect: false, Attrs: nil, Class: ""}
}

func TestBtnRendersTheExamples(t *testing.T) {
	buy := "#buy"
	primary, secondary, neutral := site.BtnVariantPrimary, site.BtnVariantSecondary, site.BtnVariantNeutral
	medium := site.BtnSizeMedium
	block := link("Buy", buy, primary, medium)
	block.Block = true
	signIn := link("Sign In", "#sign-in", primary, site.BtnSizeXLarge)
	signIn.Rect = true
	store := link("Continue Shopping", "#store", secondary, site.BtnSizeXLarge)
	store.Rect = true
	disabled := func(variant site.BtnVariant) site.Btn {
		b := link("Buy", "", variant, medium)
		b.Attrs = h.Disabled()
		return b
	}
	for _, c := range []struct {
		path string
		node g.Node
	}{
		{"default", link("Buy", buy, primary, medium)},
		{"block", block},
		{"rect", g.Group{signIn, store}},
		{"cta", g.Group{link("Learn more", "#learn", primary, site.BtnSizeCTA), link("Buy", buy, secondary, site.BtnSizeCTA)}},
		{"sizes", g.Group{link("Buy", buy, primary, site.BtnSizeSmall), link("Buy", buy, primary, medium), link("Buy", buy, primary, site.BtnSizeLarge), link("Buy", buy, primary, site.BtnSizeXLarge)}},
		{"variants", g.Group{link("Buy", buy, primary, medium), link("Buy", buy, secondary, medium), link("Buy", buy, neutral, medium), link("Buy", buy, site.BtnVariantSecondaryNeutral, medium)}},
		{"disabled", g.Group{disabled(primary), disabled(secondary), disabled(neutral)}},
	} {
		example.Check(t, c.node, "site/button/examples/"+c.path+".html")
	}
}

func TestBtnSubmits(t *testing.T) {
	b := link("Buy", "", site.BtnVariantPrimary, site.BtnSizeMedium)
	b.Submit = true
	want := `<button class="cui-btn" type="submit">Buy</button>`
	if got := example.Render(b); got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestBtnPanics(t *testing.T) {
	example.Panics(t, func() site.Btn { return link("Buy", "#buy", site.BtnVariantPrimary, site.BtnSizeMedium) }, []example.Break[site.Btn]{
		{Want: "site.Btn needs a Label", Change: func(b *site.Btn) { b.Label = "" }},
		{Want: "site.Btn has an Href and Submit: a link sends no form", Change: func(b *site.Btn) { b.Submit = true }},
	})
}
