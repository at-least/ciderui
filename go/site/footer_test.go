package site_test

import (
	"io"
	"strings"
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
				{Label: "Store", Href: "/store/"},
				{Label: "Laptops", Href: "/laptops/"},
				{Label: "Phones", Href: "/phones/"},
			}},
			{ID: "footer-wallet", Title: "Payments", Links: []site.Link{
				{Label: "Wallet", Href: "/wallet/"},
				{Label: "Cider Card", Href: "/cider-card/"},
			}},
		},
		{
			{ID: "footer-account", Title: "Account", Links: []site.Link{
				{Label: "Your Account", Href: "/account/"},
				{Label: "Order History", Href: "/shop/account/"},
			}},
		},
	}
	links := []site.Link{
		{Label: "Privacy Policy", Href: "/privacy/"},
		{Label: "Terms of Use", Href: "/terms/"},
		{Label: "Site Map", Href: "/sitemap/"},
	}
	example.Check(t, site.Footer{Notes: notes, NotesAriaLabel: "Footnotes", Directory: directory, DirectoryAriaLabel: "Directory", Shop: shop, Copyright: g.Text("Copyright © 2026 Example Inc. All rights reserved."), Links: links, Locale: &site.Link{Label: "United States", Href: "/choose-country/"}, Class: ""}, "site/footer/examples/footer.html")
}

func TestFooterNoNotesNoDirectoryNoShop(t *testing.T) {
	got := example.Render(site.Footer{Notes: nil, NotesAriaLabel: "", Directory: nil, DirectoryAriaLabel: "", Shop: nil, Copyright: g.Text("© 2026"), Links: nil, Locale: nil, Class: ""})
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

// Notes, Shop and Copyright are the page's nodes: rendering one twice may
// cost, or render differently the second time
func TestFooterItsNotesAndShopRenderedOnce(t *testing.T) {
	var notes, shop, copyright int
	example.Render(site.Footer{Notes: counted{&notes}, NotesAriaLabel: "Footnotes", Directory: nil, DirectoryAriaLabel: "", Shop: counted{&shop}, Copyright: counted{&copyright}, Links: nil, Locale: nil, Class: ""})
	if notes != 1 || shop != 1 || copyright != 1 {
		t.Errorf("notes rendered %d times, shop %d, copyright %d (want 1, 1, 1)", notes, shop, copyright)
	}
}

func TestFooterWritesItsWordsInThePagesLanguage(t *testing.T) {
	got := example.Render(site.Footer{Notes: g.Text("x"), NotesAriaLabel: "附註", Directory: [][]site.FooterSection{{{ID: "x", Title: "x", Links: nil}}}, DirectoryAriaLabel: "網站導覽", Shop: nil, Copyright: g.Text("©"), Links: nil, Locale: nil, Class: ""})
	for _, want := range []string{`<section class="cui-footer-notes" aria-label="附註">`, `<nav class="cui-footer-directory" aria-label="網站導覽">`} {
		if !strings.Contains(got, want) {
			t.Errorf("%s has no %s", got, want)
		}
	}
}

func TestFooterPanics(t *testing.T) {
	example.Panics(t, func() site.Footer {
		return site.Footer{Notes: nil, NotesAriaLabel: "", Directory: nil, DirectoryAriaLabel: "", Shop: nil, Copyright: g.Text("©"), Links: nil, Locale: nil, Class: ""}
	}, []example.Break[site.Footer]{
		{Want: "site.Footer needs a Copyright", Change: func(p *site.Footer) { p.Copyright = nil }},
		{Want: "site.Footer's Notes need a NotesAriaLabel", Change: func(p *site.Footer) { p.Notes = g.Text("x") }},
		{Want: "site.Footer's Directory needs a DirectoryAriaLabel", Change: func(p *site.Footer) { p.Directory = [][]site.FooterSection{{{ID: "x", Title: "x", Links: nil}}} }},
		{Want: "site.FooterSection needs an ID", Change: func(p *site.Footer) {
			p.DirectoryAriaLabel, p.Directory = "x", [][]site.FooterSection{{{ID: "", Title: "x", Links: nil}}}
		}},
		{Want: "site.FooterSection needs a Title", Change: func(p *site.Footer) {
			p.DirectoryAriaLabel, p.Directory = "x", [][]site.FooterSection{{{ID: "x", Title: "", Links: nil}}}
		}},
	})
}
