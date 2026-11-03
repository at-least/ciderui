package site_test

import (
	"io"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestFooterRendersTheExample(t *testing.T) {
	notes := h.P(g.Text("Credit for a trade-in depends on the device. Details at "), h.A(h.Href("/shop/trade-in/"), g.Text("example.com/shop/trade-in")), g.Text("."))
	shop := g.Group{g.Text("Want to see it first? "), h.A(h.Href("/retail/"), g.Text("Visit a store")), g.Text(" near you.")}
	directory := [][]site.FooterSection{
		{
			{ID: "footer-shop", Title: "Products", Links: []site.Link{
				{Text: "Store", Href: "/store/"},
				{Text: "Laptops", Href: "/laptops/"},
				{Text: "Phones", Href: "/phones/"},
			}},
			{ID: "footer-wallet", Title: "Payments", Links: []site.Link{
				{Text: "Wallet", Href: "/wallet/"},
				{Text: "Cider Card", Href: "/cider-card/"},
			}},
		},
		{
			{ID: "footer-account", Title: "Account", Links: []site.Link{
				{Text: "Your Account", Href: "/account/"},
				{Text: "Order History", Href: "/shop/account/"},
			}},
		},
	}
	links := []site.Link{
		{Text: "Privacy Policy", Href: "/privacy/"},
		{Text: "Terms of Use", Href: "/terms/"},
		{Text: "Site Map", Href: "/sitemap/"},
	}
	example.Check(t, site.Footer{Notes: notes, Directory: directory, Shop: shop, Copyright: "Copyright © 2026 Example Inc. All rights reserved.", Links: links, Locale: &site.Link{Text: "United States", Href: "/choose-country/"}, Class: ""}, "site/footer/examples/footer.html")
}

func TestFooterNoNotesNoDirectoryNoShop(t *testing.T) {
	got := example.Render(site.Footer{Notes: nil, Directory: nil, Shop: nil, Copyright: "© 2026", Links: nil, Locale: nil, Class: ""})
	want := `<footer class="cui-footer"><div class="cui-footer-content"><div class="cui-footer-legal"><div class="cui-footer-row"><p class="cui-footer-copyright">© 2026</p></div></div></div></footer>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

type counted struct{ n *int }

func (c counted) Render(w io.Writer) error {
	*c.n++
	_, err := io.WriteString(w, "x")
	return err
}

// Notes and Shop are the page's nodes: rendering one twice may cost, or
// render differently the second time
func TestFooterItsNotesAndShopRenderedOnce(t *testing.T) {
	var notes, shop int
	example.Render(site.Footer{Notes: counted{&notes}, Directory: nil, Shop: counted{&shop}, Copyright: "©", Links: nil, Locale: nil, Class: ""})
	if notes != 1 || shop != 1 {
		t.Errorf("notes rendered %d times, shop %d (want 1, 1)", notes, shop)
	}
}

func TestFooterASectionWithNoIDPanics(t *testing.T) {
	mustPanic(t, "FooterSection needs an ID", func() {
		example.Render(site.Footer{Notes: nil, Directory: [][]site.FooterSection{{{ID: "", Title: "x", Links: nil}}}, Shop: nil, Copyright: "©", Links: nil, Locale: nil, Class: ""})
	})
}

func TestFooterAnEmptyCopyrightPanics(t *testing.T) {
	mustPanic(t, "Footer needs a Copyright", func() {
		example.Render(site.Footer{Notes: nil, Directory: nil, Shop: nil, Copyright: "", Links: nil, Locale: nil, Class: ""})
	})
}

func TestFooterASectionWithNoTitlePanics(t *testing.T) {
	mustPanic(t, "FooterSection needs a Title", func() {
		example.Render(site.Footer{Notes: nil, Directory: [][]site.FooterSection{{{ID: "x", Title: "", Links: nil}}}, Shop: nil, Copyright: "©", Links: nil, Locale: nil, Class: ""})
	})
}
