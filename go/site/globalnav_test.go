package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func globalNav(links []site.GlobalNavLink) site.GlobalNav {
	return site.GlobalNav{Wordmark: "Cider", Home: "/", Links: links, AriaLabel: "Global", SkipLabel: "Skip to Content", MenuAriaLabel: "Menu", CloseAriaLabel: "Close"}
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
	got := example.Render(site.GlobalNav{Wordmark: "西打", Home: "/zh-tw/", Links: []site.GlobalNavLink{{Label: "商店", Href: "/store", Post: false}}, AriaLabel: "主要選單", SkipLabel: "跳到主要內容", MenuAriaLabel: "選單", CloseAriaLabel: "關閉"})
	for _, want := range []string{
		`<a class="cui-btn cui-skip" href="#main">跳到主要內容</a><nav class="cui-globalnav" aria-label="主要選單">`,
		`<a class="cui-globalnav-wordmark" href="/zh-tw/">西打</a>`,
		`class="cui-globalnav-menubutton" popovertarget="cui-globalnav-menu" aria-label="選單">`,
		`popovertargetaction="hide" aria-label="關閉"`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("a word the nav writes is not the caller's: no %s in %s", want, got)
		}
	}
}

func TestGlobalNavPostLinkSendsAPostToItsHref(t *testing.T) {
	got := example.Render(globalNav([]site.GlobalNavLink{{Label: "Log Out", Href: "/account/logout", Post: true}}))
	for _, want := range []string{
		`<li><button class="cui-globalnav-link" type="submit" form="cui-globalnav-post-1">Log Out</button></li>`,
		`<form id="cui-globalnav-post-1" action="/account/logout" method="post" hidden></form>`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("the POST link does not submit its own label's form to its own Href: no %s in %s", want, got)
		}
	}
}

func TestGlobalNavPanics(t *testing.T) {
	example.Panics(t, func() site.GlobalNav {
		return globalNav([]site.GlobalNavLink{{Label: "Store", Href: "/store", Post: false}})
	}, []example.Break[site.GlobalNav]{
		{Want: "site.GlobalNav needs a Home", Change: func(p *site.GlobalNav) { p.Home = "" }},
		{Want: "site.GlobalNav needs a Wordmark", Change: func(p *site.GlobalNav) { p.Wordmark = "" }},
		{Want: "site.GlobalNav needs Links", Change: func(p *site.GlobalNav) { p.Links = nil }},
		{Want: "site.GlobalNav needs an AriaLabel", Change: func(p *site.GlobalNav) { p.AriaLabel = "" }},
		{Want: "site.GlobalNav needs a SkipLabel", Change: func(p *site.GlobalNav) { p.SkipLabel = "" }},
		{Want: "site.GlobalNav needs a MenuAriaLabel", Change: func(p *site.GlobalNav) { p.MenuAriaLabel = "" }},
		{Want: "site.GlobalNav needs a CloseAriaLabel", Change: func(p *site.GlobalNav) { p.CloseAriaLabel = "" }},
		{Want: "site.GlobalNav's Links need their Label and Href", Change: func(p *site.GlobalNav) { p.Links[0].Label = "" }},
		{Want: "site.GlobalNav's Links need their Label and Href", Change: func(p *site.GlobalNav) { p.Links[0].Href = "" }},
	})
}
