package site_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/internal/glyph"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
)

var globalNavLinks = []site.NavLink{
	{Label: "Store", Href: "/store", Action: nil},
	{Label: "Laptops", Href: "/laptops", Action: nil},
	{Label: "Tablets", Href: "/tablets", Action: nil},
	{Label: "Support", Href: "/support", Action: nil},
}

var globalNavText = site.GlobalNavText{AriaLabel: "Global", MenuAriaLabel: "Menu", CloseAriaLabel: "Close", MainMenuAriaLabel: "Main menu"}

func globalNavLogo() g.Node {
	return g.Raw(`<svg aria-hidden="true" width="14" height="44" viewBox="0 0 14 44"><circle cx="7" cy="22" r="6" fill="currentColor"/></svg>`)
}

func TestGlobalNavRendersTheExample(t *testing.T) {
	example.Check(t, site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.LinkedSearch{Href: "/search", AriaLabel: "Search"}, Bag: site.LinkedBag{Href: "/bag", AriaLabel: "Bag"}, Text: globalNavText, Clear: false, Class: ""}, "site/globalnav/examples/globalnav.html")
}

var globalNavStore = []site.FlyoutGroup{
	{Heading: "Shop", Large: []site.Link{
		{Label: "New Arrivals", Href: "/store"},
		{Label: "Laptops", Href: "/shop/laptops"},
		{Label: "Tablets", Href: "/shop/tablets"},
		{Label: "Phones", Href: "/shop/phones"},
	}, Links: []site.Link{{Label: "Accessories", Href: "/shop/accessories"}}},
	{Heading: "Shortcuts", Large: nil, Links: []site.Link{
		{Label: "Visit a Store", Href: "/retail"},
		{Label: "Track an Order", Href: "/shop/orders"},
		{Label: "Trade In", Href: "/shop/trade-in"},
	}},
}

var globalNavLaptops = []site.FlyoutGroup{
	{Heading: "Laptops", Large: []site.Link{
		{Label: "All Laptops", Href: "/laptops"},
		{Label: "Cider Book 13", Href: "/cider-book-13"},
		{Label: "Cider Book 15", Href: "/cider-book-15"},
		{Label: "Cider Desk", Href: "/cider-desk"},
	}, Links: []site.Link{{Label: "Compare Models", Href: "/laptops/compare"}}},
	{Heading: "Buy", Large: nil, Links: []site.Link{
		{Label: "Buy a Laptop", Href: "/shop/laptops"},
		{Label: "Laptop Accessories", Href: "/shop/laptops/accessories"},
		{Label: "Financing", Href: "/shop/financing"},
	}},
	{Heading: "More", Large: nil, Links: []site.Link{
		{Label: "Laptop Support", Href: "/support/laptops"},
		{Label: "Cider OS", Href: "/cider-os"},
		{Label: "Sync Across Devices", Href: "/sync"},
	}},
}

func TestGlobalNavRendersTheFlyoutExample(t *testing.T) {
	flyouts := []site.NavLink{
		{Label: "Store", Href: "/store", Action: site.NavFlyout{Groups: globalNavStore, AriaLabel: "Store menu"}},
		{Label: "Laptops", Href: "/laptops", Action: site.NavFlyout{Groups: globalNavLaptops, AriaLabel: "Laptops menu"}},
		{Label: "Tablets", Href: "/tablets", Action: nil},
		{Label: "Support", Href: "/support", Action: nil},
	}
	example.Check(t, site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: flyouts, Search: site.LinkedSearch{Href: "/search", AriaLabel: "Search"}, Bag: site.LinkedBag{Href: "/bag", AriaLabel: "Bag"}, Text: globalNavText, Clear: false, Class: ""}, "site/globalnav/examples/flyout.html")
}

func TestGlobalNavRendersTheBagExample(t *testing.T) {
	emptyBag := site.FlyoutBag{
		AriaLabel: "Bag",
		Message:   "Nothing here yet.",
		Line:      &site.BagLine{Label: "Sign in", Href: "/signin", Rest: " for items you saved earlier"},
		Heading:   "Account",
		Links: []site.BagLink{
			{Icon: site.BagLinkIconOrders, Label: "Orders", Href: "/orders"},
			{Icon: site.BagLinkIconSaves, Label: "Saved Items", Href: "/saves"},
			{Icon: site.BagLinkIconAccount, Label: "Settings", Href: "/account"},
			{Icon: site.BagLinkIconSignIn, Label: "Sign in", Href: "/signin"},
		},
	}
	example.Check(t, site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.LinkedSearch{Href: "/search", AriaLabel: "Search"}, Bag: emptyBag, Text: globalNavText, Clear: false, Class: ""}, "site/globalnav/examples/bag.html")
}

var globalNavQuick = []site.Link{
	{Label: "Visit a Store", Href: "/retail"},
	{Label: "Laptops", Href: "/laptops"},
	{Label: "Tablets", Href: "/tablets"},
	{Label: "Gift Cards", Href: "/gift-cards"},
	{Label: "Support", Href: "/support"},
}

func TestGlobalNavRendersTheSearchExample(t *testing.T) {
	example.Check(t, site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.FlyoutSearch{Action: "/search", AriaLabel: "Search", Placeholder: "Search", ClearAriaLabel: "Clear search", Heading: "Shortcuts", Links: globalNavQuick}, Bag: site.LinkedBag{Href: "/bag", AriaLabel: "Bag"}, Text: globalNavText, Clear: false, Class: ""}, "site/globalnav/examples/search.html")
}

// the suggestions' example is the search's as a site's script leaves it, its
// sections in the quick links' place, a link's row after the package's arrow
func TestGlobalNavTheSuggestionsExampleIsTheSearchWithItsSections(t *testing.T) {
	search := example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.FlyoutSearch{Action: "/search", AriaLabel: "Search", Placeholder: "Search", ClearAriaLabel: "Clear search", Heading: "Shortcuts", Links: globalNavQuick}, Bag: site.LinkedBag{Href: "/bag", AriaLabel: "Bag"}, Text: globalNavText, Clear: false, Class: ""})
	suggestions := example.Read("site/globalnav/examples/suggestions.html")
	sections := func(html string) (int, int) {
		start := strings.Index(html, `<h2 class="cui-globalnav-searchheading"`)
		end := strings.LastIndex(html, "</ul>")
		if start < 0 || end < 0 {
			t.Fatalf("no sections in %s", html)
		}
		return start, end + len("</ul>")
	}
	a, b := sections(search)
	c, d := sections(suggestions)
	if got := suggestions[:c] + search[a:b] + suggestions[d:]; got != search {
		t.Errorf("the suggestions with the search's sections\n got: %s\nwant: %s", got, search)
	}
	images := func(regular, compact glyph.Icon) string {
		return example.Render(g.Group{regular.Class("cui-globalnav-image"), compact.Class("cui-globalnav-image-compact")})
	}
	want := fmt.Sprintf(`href="/laptops">%s<span><mark>La</mark>ptops</span>`, images(glyph.Arrow, glyph.ArrowCompact))
	if !strings.Contains(suggestions, want) {
		t.Errorf("the suggestions have no %s", want)
	}
}

func TestGlobalNavASearchWithNoQuickLinksHasNoHeading(t *testing.T) {
	search := flyoutSearch()
	search.Action = "/find?a&b"
	got := example.Render(navWith(search, nil))
	if !strings.Contains(got, `<form class="cui-globalnav-searchfield" action="/find?a&amp;b" method="get" role="search">`) {
		t.Errorf("got %s", got)
	}
	if strings.Contains(got, "cui-globalnav-searchheading") || strings.Contains(got, "cui-globalnav-searchlinks") {
		t.Errorf("got %s", got)
	}
}

func TestGlobalNavClearAndNoIcons(t *testing.T) {
	n := validGlobalNav()
	n.Clear = true
	got := example.Render(n)
	if !strings.HasPrefix(got, `<nav class="cui-globalnav" data-variant="clear" aria-label="Global">`) || strings.Contains(got, "cui-globalnav-icon") {
		t.Errorf("got %s", got)
	}
}

// every word the nav writes of its own is the caller's: none of the English
// is left when every one is another language's
func TestGlobalNavWritesItsTextInThePagesLanguage(t *testing.T) {
	text := site.GlobalNavText{AriaLabel: "主要選單", MenuAriaLabel: "選單", CloseAriaLabel: "關閉", MainMenuAriaLabel: "主選單"}
	flyout := []site.FlyoutGroup{{Heading: "選購", Large: nil, Links: []site.Link{{Label: "配件", Href: "/shop/accessories"}}}}
	links := []site.NavLink{{Label: "商店", Href: "/store", Action: site.NavFlyout{Groups: flyout, AriaLabel: "商店選單"}}}
	search := site.FlyoutSearch{Action: "/search", AriaLabel: "搜尋", Placeholder: "搜尋網站", ClearAriaLabel: "清除搜尋", Heading: "快速連結", Links: []site.Link{{Label: "門市", Href: "/retail"}}}
	bag := site.FlyoutBag{AriaLabel: "購物袋", Message: "購物袋是空的。", Line: nil, Heading: "", Links: nil}
	got := example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: links, Search: search, Bag: bag, Text: text, Clear: false, Class: ""})
	for _, want := range []string{
		`<nav class="cui-globalnav" aria-label="主要選單">`,
		`class="cui-globalnav-menubutton" popovertarget="cui-globalnav-menu" aria-label="選單">`,
		`class="cui-globalnav-backbutton" popovertarget="cui-globalnav-flyout-1" popovertargetaction="hide" aria-label="主選單">`,
		`class="cui-globalnav-flyoutbutton" popovertarget="cui-globalnav-flyout-1" aria-label="商店選單">`,
		`popovertarget="cui-globalnav-search" aria-label="搜尋">`,
		`placeholder="搜尋網站" aria-label="搜尋"`,
		`type="reset" aria-label="清除搜尋">`,
		`id="cui-globalnav-searchheading">快速連結</h2>`,
		`popovertarget="cui-globalnav-bag" aria-label="購物袋">`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("%s has no %s", got, want)
		}
	}
	if n := strings.Count(got, `aria-label="關閉"`); n != 4 {
		t.Errorf("%d close buttons say 關閉 in %s, want the menu's, the flyout's, the search's and the bag's", n, got)
	}
	for _, english := range []string{"Global", "Menu", "Close", "Main menu", "Search", "Clear search", "Shortcuts", "Bag", " menu"} {
		if strings.Contains(got, `"`+english+`"`) || strings.Contains(got, ">"+english+"<") {
			t.Errorf("%q left in %s", english, got)
		}
	}
}

func TestGlobalNavRendersThePostExample(t *testing.T) {
	links := []site.NavLink{
		{Label: "Store", Href: "/store", Action: nil},
		{Label: "Laptops", Href: "/laptops", Action: nil},
		{Label: "Tablets", Href: "/tablets", Action: nil},
		{Label: "Sign Out", Href: "/signout", Action: site.NavPost{}},
	}
	example.Check(t, site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: links, Search: site.LinkedSearch{Href: "/search", AriaLabel: "Search"}, Bag: site.LinkedBag{Href: "/bag", AriaLabel: "Bag"}, Text: globalNavText, Clear: false, Class: ""}, "site/globalnav/examples/post.html")
}

func validGlobalNav() site.GlobalNav {
	return site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Label: "Store", Href: "/store", Action: nil}}, Search: nil, Bag: nil, Text: globalNavText, Clear: false, Class: ""}
}

func flyoutOf(groups ...site.FlyoutGroup) site.NavFlyout {
	return site.NavFlyout{Groups: groups, AriaLabel: "Store menu"}
}

func flyoutSearch() site.FlyoutSearch {
	return site.FlyoutSearch{Action: "/s", AriaLabel: "x", Placeholder: "x", ClearAriaLabel: "x", Heading: "", Links: nil}
}

func flyoutBag() site.FlyoutBag {
	return site.FlyoutBag{AriaLabel: "Bag", Message: "x", Line: nil, Heading: "Account", Links: []site.BagLink{{Icon: site.BagLinkIconOrders, Label: "Orders", Href: "/orders"}}}
}

func TestGlobalNavPanics(t *testing.T) {
	example.Panics(t, validGlobalNav, []example.Break[site.GlobalNav]{
		{Want: "site.GlobalNav needs a Home", Change: func(p *site.GlobalNav) { p.Home = "" }},
		{Want: "site.GlobalNav needs a LogoAriaLabel", Change: func(p *site.GlobalNav) { p.LogoAriaLabel = "" }},
		{Want: "site.GlobalNav needs Links", Change: func(p *site.GlobalNav) { p.Links = nil }},
		{Want: "site.GlobalNav's Text needs its AriaLabel", Change: func(p *site.GlobalNav) { p.Text.AriaLabel = "" }},
		{Want: "site.GlobalNav's Text needs its MenuAriaLabel", Change: func(p *site.GlobalNav) { p.Text.MenuAriaLabel = "" }},
		{Want: "site.GlobalNav's Text needs its CloseAriaLabel", Change: func(p *site.GlobalNav) { p.Text.CloseAriaLabel = "" }},
		{Want: "site.GlobalNav's Links need their Label and Href", Change: func(p *site.GlobalNav) { p.Links[0].Label = "" }},
		{Want: "site.GlobalNav's Links need their Label and Href", Change: func(p *site.GlobalNav) { p.Links[0].Href = "" }},
		{Want: "site.GlobalNav's Text needs its MainMenuAriaLabel", Change: func(p *site.GlobalNav) {
			p.Text.MainMenuAriaLabel, p.Links[0].Action = "", flyoutOf(site.FlyoutGroup{Heading: "x", Large: nil, Links: nil})
		}},
		{Want: "site.NavFlyout needs an AriaLabel", Change: func(p *site.GlobalNav) {
			p.Links[0].Action = site.NavFlyout{Groups: []site.FlyoutGroup{{Heading: "x", Large: nil, Links: nil}}, AriaLabel: ""}
		}},
		{Want: "site.NavFlyout needs Groups", Change: func(p *site.GlobalNav) { p.Links[0].Action = flyoutOf() }},
		{Want: "site.NavFlyout's Groups need their Heading", Change: func(p *site.GlobalNav) {
			p.Links[0].Action = flyoutOf(site.FlyoutGroup{Heading: "", Large: nil, Links: nil})
		}},
		{Want: "site.NavFlyout's Large links are only in its first group", Change: func(p *site.GlobalNav) {
			p.Links[0].Action = flyoutOf(site.FlyoutGroup{Heading: "a", Large: nil, Links: nil}, site.FlyoutGroup{Heading: "b", Large: []site.Link{{Label: "x", Href: "/x"}}, Links: nil})
		}},
		{Want: "site.NavLink's Action is neither a NavFlyout nor a NavPost", Change: func(p *site.GlobalNav) { p.Links[0].Action = &site.NavPost{} }},
		{Want: "site.LinkedSearch needs an AriaLabel", Change: func(p *site.GlobalNav) { p.Search = site.LinkedSearch{Href: "/s", AriaLabel: ""} }},
		{Want: "site.GlobalNav's Search is neither a LinkedSearch nor a FlyoutSearch", Change: func(p *site.GlobalNav) { p.Search = &site.LinkedSearch{Href: "/s", AriaLabel: "x"} }},
		{Want: "site.LinkedBag needs an AriaLabel", Change: func(p *site.GlobalNav) { p.Bag = site.LinkedBag{Href: "/b", AriaLabel: ""} }},
		{Want: "site.GlobalNav's Bag is neither a LinkedBag nor a FlyoutBag", Change: func(p *site.GlobalNav) { p.Bag = &site.LinkedBag{Href: "/b", AriaLabel: "x"} }},
	})
}

func navWith(search site.Search, bag site.Bag) site.GlobalNav {
	n := validGlobalNav()
	n.Search, n.Bag = search, bag
	return n
}

func TestFlyoutSearchPanics(t *testing.T) {
	broken := func(change func(*site.FlyoutSearch)) func(*site.GlobalNav) {
		return func(n *site.GlobalNav) {
			s := flyoutSearch()
			change(&s)
			n.Search = s
		}
	}
	example.Panics(t, func() site.GlobalNav { return navWith(flyoutSearch(), nil) }, []example.Break[site.GlobalNav]{
		{Want: "site.FlyoutSearch needs an Action", Change: broken(func(s *site.FlyoutSearch) { s.Action = "" })},
		{Want: "site.FlyoutSearch needs an AriaLabel", Change: broken(func(s *site.FlyoutSearch) { s.AriaLabel = "" })},
		{Want: "site.FlyoutSearch needs a Placeholder", Change: broken(func(s *site.FlyoutSearch) { s.Placeholder = "" })},
		{Want: "site.FlyoutSearch needs a ClearAriaLabel", Change: broken(func(s *site.FlyoutSearch) { s.ClearAriaLabel = "" })},
		{Want: "site.FlyoutSearch's Links need their Heading", Change: broken(func(s *site.FlyoutSearch) { s.Links = []site.Link{{Label: "x", Href: "/x"}} })},
	})
}

func TestFlyoutBagPanics(t *testing.T) {
	broken := func(change func(*site.FlyoutBag)) func(*site.GlobalNav) {
		return func(n *site.GlobalNav) {
			b := flyoutBag()
			change(&b)
			n.Bag = b
		}
	}
	example.Panics(t, func() site.GlobalNav { return navWith(nil, flyoutBag()) }, []example.Break[site.GlobalNav]{
		{Want: "site.FlyoutBag needs an AriaLabel", Change: broken(func(b *site.FlyoutBag) { b.AriaLabel = "" })},
		{Want: "site.FlyoutBag needs a Message", Change: broken(func(b *site.FlyoutBag) { b.Message = "" })},
		{Want: "site.FlyoutBag's Line needs its link's Label and Href", Change: broken(func(b *site.FlyoutBag) { b.Line = &site.BagLine{Label: "", Href: "/x", Rest: ""} })},
		{Want: "site.FlyoutBag's Line needs its link's Label and Href", Change: broken(func(b *site.FlyoutBag) { b.Line = &site.BagLine{Label: "x", Href: "", Rest: ""} })},
		{Want: "site.FlyoutBag's Links need their Heading", Change: broken(func(b *site.FlyoutBag) { b.Heading = "" })},
		{Want: "site.FlyoutBag's Links need their Label and Href", Change: broken(func(b *site.FlyoutBag) { b.Links[0].Label = "" })},
		{Want: "site.FlyoutBag's Links need their Label and Href", Change: broken(func(b *site.FlyoutBag) { b.Links[0].Href = "" })},
		{Want: "site.FlyoutBag's Links need their Icon", Change: broken(func(b *site.FlyoutBag) { b.Links[0].Icon = site.BagLinkIcon{} })},
	})
}

// a nav with no flyout draws no back button, so writes no word for it
func TestGlobalNavNoFlyoutNeedsNoMainMenuAriaLabel(t *testing.T) {
	n := validGlobalNav()
	n.Text.MainMenuAriaLabel = ""
	if got := example.Render(n); strings.Contains(got, "cui-globalnav-backbutton") {
		t.Errorf("got %s", got)
	}
}

func TestGlobalNavBagHeadingIsItsLinks(t *testing.T) {
	n := validGlobalNav()
	bag := flyoutBag()
	n.Bag = bag
	if got := example.Render(n); !strings.Contains(got, `id="cui-globalnav-bagheading">Account</h3>`) {
		t.Errorf("one link, no heading: %s", got)
	}
	bag.Links = nil
	n.Bag = bag
	if got := example.Render(n); strings.Contains(got, "cui-globalnav-bagheading") {
		t.Errorf("no link, a heading: %s", got)
	}
}
