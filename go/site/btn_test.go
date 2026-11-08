package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func link(label, href string, variant site.BtnVariant, size site.BtnSize) site.Btn {
	return site.Btn{Href: href, Label: label, Icon: nil, Submit: false, Variant: variant, Size: size, Block: false, Attrs: nil, Class: ""}
}

func TestBtnRendersTheExamples(t *testing.T) {
	buy := "#buy"
	primary, secondary := site.BtnVariantPrimary, site.BtnVariantSecondary
	medium := site.BtnSizeMedium
	block := link("Buy", buy, primary, medium)
	block.Block = true
	blockButton := link("Buy", "", primary, medium)
	blockButton.Block = true
	icon := link("Added", "", primary, medium)
	icon.Icon = lucide.Check
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
		{"block-button", blockButton},
		{"icon", icon},
		{"sizes", g.Group{link("Buy", buy, primary, medium), link("Buy", buy, primary, site.BtnSizeLarge)}},
		{"variants", g.Group{link("Buy", buy, primary, medium), link("Buy", buy, secondary, medium), link("Buy", buy, site.BtnVariantSecondaryNeutral, medium)}},
		{"disabled", g.Group{disabled(primary), disabled(secondary)}},
	} {
		example.Check(t, c.node, "site/btn/examples/"+c.path+".html")
	}
}

func TestBtnSubmits(t *testing.T) {
	b := link("Buy", "", site.BtnVariantPrimary, site.BtnSizeMedium)
	b.Submit = true
	want := `<button class="cui-btn" type="submit">Buy</button>`
	if got := example.Render(b); got != want {
		t.Errorf("the button does not submit its form\n got: %s\nwant: %s", got, want)
	}
}

func TestBtnAttrsGoOnItsLink(t *testing.T) {
	b := link("Buy", "/buy", site.BtnVariantPrimary, site.BtnSizeMedium)
	b.Attrs = h.Aria("current", "page")
	want := `<a class="cui-btn" href="/buy" aria-current="page">Buy</a>`
	if got := example.Render(b); got != want {
		t.Errorf("the link button lost its Attrs\n got: %s\nwant: %s", got, want)
	}
}

func TestBtnDrawsItsIconBeforeItsLabel(t *testing.T) {
	b := link("Sign in with Google", "#login", site.BtnVariantSecondaryNeutral, site.BtnSizeLarge)
	b.Icon = lucide.Check
	got := example.Render(b)
	if !strings.HasPrefix(got, `<a class="cui-btn" data-variant="secondary-neutral" data-size="lg" href="#login"><svg aria-hidden="true"`) ||
		!strings.HasSuffix(got, `</svg>Sign in with Google</a>`) {
		t.Errorf("the icon is not first inside the link, before its label: %s", got)
	}
}

func TestBtnPanics(t *testing.T) {
	example.Panics(t, func() site.Btn { return link("Buy", "#buy", site.BtnVariantPrimary, site.BtnSizeMedium) }, []example.Break[site.Btn]{
		{Want: "site.Btn needs a Label", Change: func(b *site.Btn) { b.Label = "" }},
		{Want: "site.Btn has an Href and Submit: a link sends no form", Change: func(b *site.Btn) { b.Submit = true }},
	})
}
