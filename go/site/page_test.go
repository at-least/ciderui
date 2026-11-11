package site_test

import (
	"testing"

	"github.com/at-least/ciderui"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func page() site.Page {
	return site.Page{Lang: "zh-Hant", Title: "課程", Head: nil, Children: h.P(g.Text("x"))}
}

func TestPageWritesTheDocumentItsHeadAfterTheStylesheet(t *testing.T) {
	p := page()
	p.Head = h.Link(h.Rel("stylesheet"), h.Href("/own.css"))
	want := `<!doctype html><html lang="zh-Hant"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>課程</title><link rel="stylesheet" href="` + ciderui.SiteStylesheet() + `"><link rel="stylesheet" href="/own.css"></head><body><p>x</p></body></html>`
	if got := example.Render(p); got != want {
		t.Errorf("the page is not the document with its head after the site stylesheet\n got: %s\nwant: %s", got, want)
	}
}

func TestPagePanics(t *testing.T) {
	example.Panics(t, page, []example.Break[site.Page]{
		{Want: "site.Page needs a Lang", Change: func(p *site.Page) { p.Lang = "" }},
		{Want: "site.Page needs a Title", Change: func(p *site.Page) { p.Title = "" }},
		// @invariant: under go test a page checks itself: a skip link to no main
		{Want: `site.Page: href="#main" on <a class="cui-btn cui-skip">: this page has no id "main"`, Change: func(p *site.Page) {
			p.Children = globalNav([]site.GlobalNavLink{{Label: "Store", Href: "/store", Post: false}})
		}},
	})
}
