package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

var localNavLinks = []site.Link{
	{Label: "Overview", Href: "/cider-book/"},
	{Label: "Tech Specs", Href: "/cider-book/specs/"},
	{Label: "Compare", Href: "/cider-book/compare/"},
	{Label: "Switch to Cider Book", Href: "/cider-book/switch/"},
}

func TestLocalNavRendersTheExample(t *testing.T) {
	example.Check(t, site.LocalNav{Title: "Cider Book 13", Href: "/cider-book/", Links: localNavLinks, Current: "/cider-book/", Button: &site.Link{Label: "Buy", Href: "/buy/cider-book/"}, AriaLabel: "Local", MenuAriaLabel: "Menu", Class: ""}, "site/localnav/examples/localnav.html")
}

func TestLocalNavNoLinksNoMenu(t *testing.T) {
	got := example.Render(site.LocalNav{Title: "AirTag", Href: "/airtag/", Links: nil, Current: "", Button: nil, AriaLabel: "Local", MenuAriaLabel: "", Class: "cui-dark"})
	want := `<nav class="cui-localnav cui-dark" aria-label="Local"><div class="cui-localnav-content">` +
		`<a class="cui-localnav-title" href="/airtag/">AirTag</a><div class="cui-localnav-actions"></div></div></nav>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestLocalNavWritesItsWordsInThePagesLanguage(t *testing.T) {
	got := example.Render(site.LocalNav{Title: "x", Href: "/x/", Links: localNavLinks, Current: "", Button: nil, AriaLabel: "本頁", MenuAriaLabel: "選單", Class: ""})
	for _, want := range []string{`<nav class="cui-localnav" aria-label="本頁">`, `id="cui-localnav-menu" aria-label="選單">`} {
		if !strings.Contains(got, want) {
			t.Errorf("%s has no %s", got, want)
		}
	}
}

func TestLocalNavPanics(t *testing.T) {
	example.Panics(t, func() site.LocalNav {
		return site.LocalNav{Title: "x", Href: "/x/", Links: []site.Link{{Label: "x", Href: "/x/"}}, Current: "", Button: nil, AriaLabel: "Local", MenuAriaLabel: "Menu", Class: ""}
	}, []example.Break[site.LocalNav]{
		{Want: "site.LocalNav needs a Title", Change: func(p *site.LocalNav) { p.Title = "" }},
		{Want: "site.LocalNav needs an Href", Change: func(p *site.LocalNav) { p.Href = "" }},
		{Want: "site.LocalNav needs an AriaLabel", Change: func(p *site.LocalNav) { p.AriaLabel = "" }},
		{Want: "site.LocalNav's Links need a MenuAriaLabel", Change: func(p *site.LocalNav) { p.MenuAriaLabel = "" }},
		{Want: "site.LocalNav's Links need their Label and Href", Change: func(p *site.LocalNav) { p.Links[0].Label = "" }},
		{Want: "site.LocalNav's Links need their Label and Href", Change: func(p *site.LocalNav) { p.Links[0].Href = "" }},
		{Want: "site.LocalNav's Current is none of its Links'", Change: func(p *site.LocalNav) { p.Current = "/elsewhere/" }},
		{Want: "site.LocalNav's Button needs its Label and Href", Change: func(p *site.LocalNav) { p.Button = &site.Link{Label: "", Href: "/buy/"} }},
		{Want: "site.LocalNav's Button needs its Label and Href", Change: func(p *site.LocalNav) { p.Button = &site.Link{Label: "Buy", Href: ""} }},
	})
}
