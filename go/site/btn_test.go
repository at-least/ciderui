package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func link(label, href string, variant site.BtnVariant, size site.BtnSize) site.Btn {
	return site.Btn{Href: href, Label: label, Submit: false, Variant: variant, Size: size, Block: false, Attrs: nil, Class: ""}
}

func TestBtnRendersTheExamples(t *testing.T) {
	buy := "#buy"
	primary, secondary := site.BtnVariantPrimary, site.BtnVariantSecondary
	medium := site.BtnSizeMedium
	block := link("Buy", buy, primary, medium)
	block.Block = true
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
		{"sizes", g.Group{link("Buy", buy, primary, medium), link("Buy", buy, primary, site.BtnSizeLarge)}},
		{"variants", g.Group{link("Buy", buy, primary, medium), link("Buy", buy, secondary, medium)}},
		{"disabled", g.Group{disabled(primary), disabled(secondary)}},
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
