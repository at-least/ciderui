package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func page() site.Page {
	return site.Page{Lang: "zh-Hant", Title: "課程", Head: nil, Class: "", Skip: "", BodyClass: "", Children: h.P(g.Text("x"))}
}

func TestPageWritesTheDocument(t *testing.T) {
	want := `<!doctype html><html lang="zh-Hant"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>課程</title><link rel="stylesheet" href="` + ciderui.SiteStylesheet() + `"></head><body><p>x</p></body></html>`
	if got := example.Render(page()); got != want {
		t.Errorf("got %s\nwant %s", got, want)
	}
}

func TestPageWritesItsClassesAndHeadAfterTheStylesheet(t *testing.T) {
	p := page()
	p.Class = "scroll-pt-16"
	p.BodyClass = "flex"
	p.Head = h.Link(h.Rel("stylesheet"), h.Href("/own.css"))
	got := example.Render(p)
	for _, want := range []string{`<html lang="zh-Hant" class="scroll-pt-16">`, `<body class="flex">`, `.css"><link rel="stylesheet" href="/own.css"></head>`} {
		if !strings.Contains(got, want) {
			t.Errorf("no %s in %s", want, got)
		}
	}
}

// the link before all else the keyboard reads, before the page's children
func TestPageWritesItsSkipLinkFirst(t *testing.T) {
	p := page()
	p.Skip = "跳到主要內容"
	got := example.Render(p)
	want := `<body><a class="cui-btn cui-skip" href="#main">跳到主要內容</a><p>x</p>`
	if !strings.Contains(got, want) {
		t.Errorf("no %s in %s", want, got)
	}
}

func TestPagePanics(t *testing.T) {
	example.Panics(t, page, []example.Break[site.Page]{
		{Want: "site.Page needs a Lang", Change: func(p *site.Page) { p.Lang = "" }},
		{Want: "site.Page needs a Title", Change: func(p *site.Page) { p.Title = "" }},
	})
}
