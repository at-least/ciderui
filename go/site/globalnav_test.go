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

func globalNav(links []site.NavLink) site.GlobalNav {
	return site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: links, AriaLabel: "Global", MenuAriaLabel: "Menu", CloseAriaLabel: "Close", Class: ""}
}

func TestGlobalNavRendersTheExample(t *testing.T) {
	example.Check(t, globalNav([]site.NavLink{
		{Label: "Store", Href: "/store", Post: false},
		{Label: "Laptops", Href: "/laptops", Post: false},
		{Label: "Tablets", Href: "/tablets", Post: false},
		{Label: "Support", Href: "/support", Post: false},
	}), "site/globalnav/examples/globalnav.html")
}

func TestGlobalNavRendersThePostExample(t *testing.T) {
	example.Check(t, globalNav([]site.NavLink{
		{Label: "Store", Href: "/store", Post: false},
		{Label: "Laptops", Href: "/laptops", Post: false},
		{Label: "Tablets", Href: "/tablets", Post: false},
		{Label: "Sign Out", Href: "/signout", Post: true},
	}), "site/globalnav/examples/post.html")
}

// every word the nav writes of its own is the caller's
func TestGlobalNavWritesItsWordsInThePagesLanguage(t *testing.T) {
	got := example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Label: "商店", Href: "/store", Post: false}}, AriaLabel: "主要選單", MenuAriaLabel: "選單", CloseAriaLabel: "關閉", Class: ""})
	for _, want := range []string{
		`<nav class="cui-globalnav" aria-label="主要選單">`,
		`class="cui-globalnav-menubutton" popovertarget="cui-globalnav-menu" aria-label="選單">`,
		`popovertargetaction="hide" aria-label="關閉"`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("%s has no %s", got, want)
		}
	}
}

func TestGlobalNavPanics(t *testing.T) {
	example.Panics(t, func() site.GlobalNav { return globalNav([]site.NavLink{{Label: "Store", Href: "/store", Post: false}}) }, []example.Break[site.GlobalNav]{
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
