package site_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/icon"
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
	images := func(regular, compact icon.Icon) string {
		return example.Render(g.Group{regular.Class("cui-globalnav-image"), compact.Class("cui-globalnav-image-compact")})
	}
	want := fmt.Sprintf(`href="/laptops">%s<span><mark>La</mark>ptops</span>`, images(glyph.Arrow, glyph.ArrowCompact))
	if !strings.Contains(suggestions, want) {
		t.Errorf("the suggestions have no %s", want)
	}
}

func TestGlobalNavASearchWithNoQuickLinksHasNoHeading(t *testing.T) {
	got := example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.FlyoutSearch{Action: "/find?a&b", AriaLabel: "Search", Placeholder: "Search", ClearAriaLabel: "Clear search", Heading: "", Links: nil}, Bag: nil, Text: globalNavText, Clear: false, Class: ""})
	if !strings.Contains(got, `<form class="cui-globalnav-searchfield" action="/find?a&amp;b" method="get" role="search">`) {
		t.Errorf("got %s", got)
	}
	if strings.Contains(got, "cui-globalnav-searchheading") || strings.Contains(got, "cui-globalnav-searchlinks") {
		t.Errorf("got %s", got)
	}
}

func TestGlobalNavASearchWithNoPagePanics(t *testing.T) {
	mustPanic(t, "site.FlyoutSearch needs an Action", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.FlyoutSearch{Action: "", AriaLabel: "x", Placeholder: "x", ClearAriaLabel: "x", Heading: "", Links: nil}, Bag: nil, Text: globalNavText, Clear: false, Class: ""})
	})
}

func TestGlobalNavClearAndThePagesUtilities(t *testing.T) {
	got := example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: nil, Text: globalNavText, Clear: true, Class: "sticky top-0"})
	if !strings.HasPrefix(got, `<nav class="cui-globalnav sticky top-0" data-variant="clear" aria-label="Global">`) {
		t.Errorf("got %s", got)
	}
	if strings.Contains(got, "cui-globalnav-icon") {
		t.Errorf("got %s", got)
	}
}

func TestGlobalNavEscapesItsText(t *testing.T) {
	got := example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/?a&b", LogoAriaLabel: `A "B"`, Links: []site.NavLink{{Label: "<b>", Href: "/x", Action: nil}}, Search: nil, Bag: nil, Text: globalNavText, Clear: false, Class: ""})
	for _, want := range []string{
		`href="/?a&amp;b" aria-label="A &quot;B&quot;"`,
		`<a class="cui-globalnav-link" href="/x">&lt;b&gt;</a>`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("%s has no %s", got, want)
		}
	}
}

func TestGlobalNavNoLinksPanics(t *testing.T) {
	mustPanic(t, "needs Links", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: nil, Search: nil, Bag: nil, Text: globalNavText, Clear: false, Class: ""})
	})
}

func TestGlobalNavAnEmptyHomePanics(t *testing.T) {
	mustPanic(t, "GlobalNav needs a Home", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: nil, Text: globalNavText, Clear: false, Class: ""})
	})
}

func TestGlobalNavAnEmptyLogoAriaLabelPanics(t *testing.T) {
	mustPanic(t, "GlobalNav needs a LogoAriaLabel", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "", Links: globalNavLinks, Search: nil, Bag: nil, Text: globalNavText, Clear: false, Class: ""})
	})
}

func TestGlobalNavALinkWithNoLabelPanics(t *testing.T) {
	mustPanic(t, "GlobalNav's Links need their Label and Href", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Label: "", Href: "/x", Action: nil}}, Search: nil, Bag: nil, Text: globalNavText, Clear: false, Class: ""})
	})
}

func TestGlobalNavAFlyoutGroupWithNoHeadingPanics(t *testing.T) {
	flyout := []site.FlyoutGroup{{Heading: "", Large: nil, Links: nil}}
	mustPanic(t, "NavFlyout's Groups need their Heading", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Label: "Store", Href: "/store", Action: site.NavFlyout{Groups: flyout, AriaLabel: "Store menu"}}}, Search: nil, Bag: nil, Text: globalNavText, Clear: false, Class: ""})
	})
}

func TestGlobalNavABagWithNoMessagePanics(t *testing.T) {
	bag := site.FlyoutBag{AriaLabel: "Bag", Message: "", Line: nil, Heading: "", Links: nil}
	mustPanic(t, "site.FlyoutBag needs a Message", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: bag, Text: globalNavText, Clear: false, Class: ""})
	})
}

func TestGlobalNavABagLineWithNoLabelPanics(t *testing.T) {
	bag := site.FlyoutBag{AriaLabel: "Bag", Message: "Nothing here yet.", Line: &site.BagLine{Label: "", Href: "/signin", Rest: ""}, Heading: "", Links: nil}
	mustPanic(t, "Line needs its link's Label and Href", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: bag, Text: globalNavText, Clear: false, Class: ""})
	})
}

func TestGlobalNavBagLinksWithNoHeadingPanics(t *testing.T) {
	bag := site.FlyoutBag{AriaLabel: "Bag", Message: "Nothing here yet.", Line: nil, Heading: "", Links: []site.BagLink{{Icon: site.BagLinkIconOrders, Label: "Orders", Href: "/orders"}}}
	mustPanic(t, "FlyoutBag's Links need their Heading", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: bag, Text: globalNavText, Clear: false, Class: ""})
	})
}

func TestGlobalNavABagLinkWithNoLabelPanics(t *testing.T) {
	bag := site.FlyoutBag{AriaLabel: "Bag", Message: "Nothing here yet.", Line: nil, Heading: "Account", Links: []site.BagLink{{Icon: site.BagLinkIconOrders, Label: "", Href: "/orders"}}}
	mustPanic(t, "FlyoutBag's Links need their Label and Href", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: bag, Text: globalNavText, Clear: false, Class: ""})
	})
}

func TestGlobalNavABagLinkWithNoIconPanics(t *testing.T) {
	bag := site.FlyoutBag{AriaLabel: "Bag", Message: "Nothing here yet.", Line: nil, Heading: "Account", Links: []site.BagLink{{Icon: site.BagLinkIcon{}, Label: "Orders", Href: "/orders"}}}
	mustPanic(t, "FlyoutBag's Links need their Icon", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: bag, Text: globalNavText, Clear: false, Class: ""})
	})
}

func TestGlobalNavLargeLinksAfterTheFirstGroupPanics(t *testing.T) {
	flyout := []site.FlyoutGroup{
		{Heading: "Shop", Large: nil, Links: nil},
		{Heading: "More", Large: []site.Link{{Label: "x", Href: "/x"}}, Links: nil},
	}
	mustPanic(t, "site.NavFlyout's Large links are only in its first group", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Label: "Store", Href: "/store", Action: site.NavFlyout{Groups: flyout, AriaLabel: "Store menu"}}}, Search: nil, Bag: nil, Text: globalNavText, Clear: false, Class: ""})
	})
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

func TestGlobalNavAnEmptyTextPanics(t *testing.T) {
	for field, blank := range map[string]func(*site.GlobalNavText){
		"AriaLabel":      func(x *site.GlobalNavText) { x.AriaLabel = "" },
		"MenuAriaLabel":  func(x *site.GlobalNavText) { x.MenuAriaLabel = "" },
		"CloseAriaLabel": func(x *site.GlobalNavText) { x.CloseAriaLabel = "" },
	} {
		text := globalNavText
		blank(&text)
		mustPanic(t, "site.GlobalNav's Text needs its "+field, func() {
			example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: nil, Text: text, Clear: false, Class: ""})
		})
	}
}

// a nav with no flyout draws no back button, so writes no word for it
func TestGlobalNavMainMenuAriaLabelIsAFlyouts(t *testing.T) {
	text := globalNavText
	text.MainMenuAriaLabel = ""
	example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: nil, Text: text, Clear: false, Class: ""})
	mustPanic(t, "site.GlobalNav's Text needs its MainMenuAriaLabel", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Label: "Store", Href: "/store", Action: site.NavFlyout{Groups: globalNavStore, AriaLabel: "Store menu"}}}, Search: nil, Bag: nil, Text: text, Clear: false, Class: ""})
	})
}

func TestGlobalNavAFlyoutWithNoAriaLabelPanics(t *testing.T) {
	mustPanic(t, "site.NavFlyout needs an AriaLabel", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Label: "Store", Href: "/store", Action: site.NavFlyout{Groups: globalNavStore, AriaLabel: ""}}}, Search: nil, Bag: nil, Text: globalNavText, Clear: false, Class: ""})
	})
}

func TestGlobalNavAFlyoutWithNoGroupsPanics(t *testing.T) {
	mustPanic(t, "site.NavFlyout needs Groups", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Label: "Store", Href: "/store", Action: site.NavFlyout{Groups: nil, AriaLabel: "x"}}}, Search: nil, Bag: nil, Text: globalNavText, Clear: false, Class: ""})
	})
}

func TestGlobalNavSearchAndBagNeedTheirWords(t *testing.T) {
	render := func(search site.Search, bag site.Bag) func() {
		return func() {
			example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: search, Bag: bag, Text: globalNavText, Clear: false, Class: ""})
		}
	}
	mustPanic(t, "site.LinkedSearch needs an AriaLabel", render(site.LinkedSearch{Href: "/s", AriaLabel: ""}, nil))
	mustPanic(t, "site.FlyoutSearch needs an AriaLabel", render(site.FlyoutSearch{Action: "/s", AriaLabel: "", Placeholder: "x", ClearAriaLabel: "x", Heading: "", Links: nil}, nil))
	mustPanic(t, "site.FlyoutSearch needs a Placeholder", render(site.FlyoutSearch{Action: "/s", AriaLabel: "x", Placeholder: "", ClearAriaLabel: "x", Heading: "", Links: nil}, nil))
	mustPanic(t, "site.FlyoutSearch needs a ClearAriaLabel", render(site.FlyoutSearch{Action: "/s", AriaLabel: "x", Placeholder: "x", ClearAriaLabel: "", Heading: "", Links: nil}, nil))
	mustPanic(t, "site.LinkedBag needs an AriaLabel", render(nil, site.LinkedBag{Href: "/b", AriaLabel: ""}))
	mustPanic(t, "site.FlyoutBag needs an AriaLabel", render(nil, site.FlyoutBag{AriaLabel: "", Message: "x", Line: nil, Heading: "", Links: nil}))
}

func TestGlobalNavSearchAsksForTheSearchKey(t *testing.T) {
	got := example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.FlyoutSearch{Action: "/s", AriaLabel: "x", Placeholder: "x", ClearAriaLabel: "x", Heading: "", Links: nil}, Bag: nil, Text: globalNavText, Clear: false, Class: ""})
	if !strings.Contains(got, `type="text" name="q"`) || !strings.Contains(got, `enterkeyhint="search"`) {
		t.Errorf("got %s", got)
	}
}

func TestGlobalNavSearchLinksWithNoHeadingPanics(t *testing.T) {
	mustPanic(t, "site.FlyoutSearch's Links need their Heading", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.FlyoutSearch{Action: "/search", AriaLabel: "x", Placeholder: "x", ClearAriaLabel: "x", Heading: "", Links: globalNavQuick}, Bag: nil, Text: globalNavText, Clear: false, Class: ""})
	})
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

// a pointer to a variant also has its method: the sealed interfaces do not stop one
func TestGlobalNavAPointerToAVariantPanics(t *testing.T) {
	render := func(link site.NavAction, search site.Search, bag site.Bag) func() {
		return func() {
			example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Label: "x", Href: "/x", Action: link}}, Search: search, Bag: bag, Text: globalNavText, Clear: false, Class: ""})
		}
	}
	mustPanic(t, "site.NavLink's Action is neither a NavFlyout nor a NavPost", render(&site.NavPost{}, nil, nil))
	mustPanic(t, "site.GlobalNav's Search is neither a LinkedSearch nor a FlyoutSearch", render(nil, &site.LinkedSearch{Href: "/s", AriaLabel: "x"}, nil))
	mustPanic(t, "site.GlobalNav's Bag is neither a LinkedBag nor a FlyoutBag", render(nil, nil, &site.LinkedBag{Href: "/b", AriaLabel: "x"}))
}
