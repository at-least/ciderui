package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
)

func globalNavLogo() g.Node {
	return g.Raw(`<svg aria-hidden="true" width="14" height="44" viewBox="0 0 14 44"><circle cx="7" cy="22" r="6" fill="currentColor"/></svg>`)
}

func globalNav(links []site.GlobalNavLink) site.GlobalNav {
	return site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: links, AriaLabel: "Global", MenuAriaLabel: "Menu", CloseAriaLabel: "Close", Class: ""}
}

func TestGlobalNavRendersTheExamples(t *testing.T) {
	example.Check(t, globalNav([]site.GlobalNavLink{
		{Label: "Store", Href: "/store", Post: false},
		{Label: "Laptops", Href: "/laptops", Post: false},
		{Label: "Tablets", Href: "/tablets", Post: false},
		{Label: "Support", Href: "/support", Post: false},
	}), "site/globalnav/examples/globalnav.html")
	example.Check(t, globalNav([]site.GlobalNavLink{
		{Label: "Store", Href: "/store", Post: false},
		{Label: "Laptops", Href: "/laptops", Post: false},
		{Label: "Tablets", Href: "/tablets", Post: false},
		{Label: "Sign Out", Href: "/signout", Post: true},
	}), "site/globalnav/examples/post.html")
}

func TestGlobalNavWritesItsWordsInThePagesLanguage(t *testing.T) {
	got := example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/zh-tw/", LogoAriaLabel: "回首頁", Links: []site.GlobalNavLink{{Label: "商店", Href: "/store", Post: false}}, AriaLabel: "主要選單", MenuAriaLabel: "選單", CloseAriaLabel: "關閉", Class: ""})
	for _, want := range []string{
		`<nav class="cui-globalnav" aria-label="主要選單">`,
		`class="cui-globalnav-logo" href="/zh-tw/" aria-label="回首頁">`,
		`class="cui-globalnav-menubutton" popovertarget="cui-globalnav-menu" aria-label="選單">`,
		`popovertargetaction="hide" aria-label="關閉"`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("a word the nav writes is not the caller's: no %s in %s", want, got)
		}
	}
}

func TestGlobalNavPanics(t *testing.T) {
	example.Panics(t, func() site.GlobalNav {
		return globalNav([]site.GlobalNavLink{{Label: "Store", Href: "/store", Post: false}})
	}, []example.Break[site.GlobalNav]{
		{Want: "site.GlobalNav needs a Home", Change: func(p *site.GlobalNav) { p.Home = "" }},
		{Want: "site.GlobalNav needs a LogoAriaLabel", Change: func(p *site.GlobalNav) { p.LogoAriaLabel = "" }},
		{Want: "site.GlobalNav needs Links", Change: func(p *site.GlobalNav) { p.Links = nil }},
		{Want: "site.GlobalNav needs an AriaLabel", Change: func(p *site.GlobalNav) { p.AriaLabel = "" }},
		{Want: "site.GlobalNav needs a MenuAriaLabel", Change: func(p *site.GlobalNav) { p.MenuAriaLabel = "" }},
		{Want: "site.GlobalNav needs a CloseAriaLabel", Change: func(p *site.GlobalNav) { p.CloseAriaLabel = "" }},
		{Want: "site.GlobalNav's Links need their Label and Href", Change: func(p *site.GlobalNav) { p.Links[0].Label = "" }},
		{Want: "site.GlobalNav's Links need their Label and Href", Change: func(p *site.GlobalNav) { p.Links[0].Href = "" }},
	})
}
