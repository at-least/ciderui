package site_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
)

var globalNavLinks = []site.NavLink{
	{Text: "Store", Href: "/store", Flyout: nil},
	{Text: "Laptops", Href: "/laptops", Flyout: nil},
	{Text: "Tablets", Href: "/tablets", Flyout: nil},
	{Text: "Support", Href: "/support", Flyout: nil},
}

func globalNavLogo() g.Node {
	return g.Raw(`<svg aria-hidden="true" width="14" height="44" viewBox="0 0 14 44"><circle cx="7" cy="22" r="6" fill="currentColor"/></svg>`)
}

func TestGlobalNavRendersTheExample(t *testing.T) {
	example.Check(t, site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.LinkedSearch{Href: "/search"}, Bag: site.LinkedBag{Href: "/bag"}, Dark: false, Clear: false, Class: ""}, "site/globalnav/examples/globalnav.html")
}

var globalNavStore = []site.FlyoutGroup{
	{Heading: "Shop", Large: []site.Link{
		{Text: "New Arrivals", Href: "/store"},
		{Text: "Laptops", Href: "/shop/laptops"},
		{Text: "Tablets", Href: "/shop/tablets"},
		{Text: "Phones", Href: "/shop/phones"},
	}, Links: []site.Link{{Text: "Accessories", Href: "/shop/accessories"}}},
	{Heading: "Shortcuts", Large: nil, Links: []site.Link{
		{Text: "Visit a Store", Href: "/retail"},
		{Text: "Track an Order", Href: "/shop/orders"},
		{Text: "Trade In", Href: "/shop/trade-in"},
	}},
}

var globalNavLaptops = []site.FlyoutGroup{
	{Heading: "Laptops", Large: []site.Link{
		{Text: "All Laptops", Href: "/laptops"},
		{Text: "Cider Book 13", Href: "/cider-book-13"},
		{Text: "Cider Book 15", Href: "/cider-book-15"},
		{Text: "Cider Desk", Href: "/cider-desk"},
	}, Links: []site.Link{{Text: "Compare Models", Href: "/laptops/compare"}}},
	{Heading: "Buy", Large: nil, Links: []site.Link{
		{Text: "Buy a Laptop", Href: "/shop/laptops"},
		{Text: "Laptop Accessories", Href: "/shop/laptops/accessories"},
		{Text: "Financing", Href: "/shop/financing"},
	}},
	{Heading: "More", Large: nil, Links: []site.Link{
		{Text: "Laptop Support", Href: "/support/laptops"},
		{Text: "Cider OS", Href: "/cider-os"},
		{Text: "Sync Across Devices", Href: "/sync"},
	}},
}

func TestGlobalNavRendersTheFlyoutExample(t *testing.T) {
	flyouts := []site.NavLink{
		{Text: "Store", Href: "/store", Flyout: globalNavStore},
		{Text: "Laptops", Href: "/laptops", Flyout: globalNavLaptops},
		{Text: "Tablets", Href: "/tablets", Flyout: nil},
		{Text: "Support", Href: "/support", Flyout: nil},
	}
	example.Check(t, site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: flyouts, Search: site.LinkedSearch{Href: "/search"}, Bag: site.LinkedBag{Href: "/bag"}, Dark: false, Clear: false, Class: ""}, "site/globalnav/examples/flyout.html")
}

func TestGlobalNavRendersTheBagExample(t *testing.T) {
	emptyBag := site.FlyoutBag{
		Message: "Nothing here yet.",
		Line:    &site.BagLine{Text: "Sign in", Href: "/signin", Rest: " for items you saved earlier"},
		Heading: "Account",
		Links: []site.BagLink{
			{Icon: icon.Orders, IconCompact: icon.OrdersCompact, Text: "Orders", Href: "/orders"},
			{Icon: icon.Saves, IconCompact: icon.SavesCompact, Text: "Saved Items", Href: "/saves"},
			{Icon: icon.Account, IconCompact: icon.AccountCompact, Text: "Settings", Href: "/account"},
			{Icon: icon.SignIn, IconCompact: icon.SignInCompact, Text: "Sign in", Href: "/signin"},
		},
	}
	example.Check(t, site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.LinkedSearch{Href: "/search"}, Bag: emptyBag, Dark: false, Clear: false, Class: ""}, "site/globalnav/examples/bag.html")
}

var globalNavQuick = []site.Link{
	{Text: "Visit a Store", Href: "/retail"},
	{Text: "Laptops", Href: "/laptops"},
	{Text: "Tablets", Href: "/tablets"},
	{Text: "Gift Cards", Href: "/gift-cards"},
	{Text: "Support", Href: "/support"},
}

func TestGlobalNavRendersTheSearchExample(t *testing.T) {
	example.Check(t, site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.FlyoutSearch{Action: "/search", Links: globalNavQuick}, Bag: site.LinkedBag{Href: "/bag"}, Dark: false, Clear: false, Class: ""}, "site/globalnav/examples/search.html")
}

// the suggestions' example is the search's as a site's script leaves it, its
// sections in the quick links' place, each row after the package's icons: a
// link's arrow, a suggested search's magnifier
func TestGlobalNavTheSuggestionsExampleIsTheSearchWithItsSections(t *testing.T) {
	search := example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.FlyoutSearch{Action: "/search", Links: globalNavQuick}, Bag: site.LinkedBag{Href: "/bag"}, Dark: false, Clear: false, Class: ""})
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
	for _, want := range []string{
		fmt.Sprintf(`href="/laptops">%s<span><mark>La</mark>ptops</span>`, images(icon.Arrow, icon.ArrowCompact)),
		fmt.Sprintf(`href="/search?q=laptop">%s<span><mark>la</mark>ptop</span>`, images(icon.Suggestion, icon.SuggestionCompact)),
	} {
		if !strings.Contains(suggestions, want) {
			t.Errorf("the suggestions have no %s", want)
		}
	}
}

func TestGlobalNavASearchWithNoQuickLinksHasNoHeading(t *testing.T) {
	got := example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.FlyoutSearch{Action: "/find?a&b", Links: nil}, Bag: nil, Dark: false, Clear: false, Class: ""})
	if !strings.Contains(got, `<form class="cui-globalnav-searchfield" action="/find?a&amp;b" method="get" role="search">`) {
		t.Errorf("got %s", got)
	}
	if strings.Contains(got, "cui-globalnav-searchheading") || strings.Contains(got, "cui-globalnav-searchlinks") {
		t.Errorf("got %s", got)
	}
}

func TestGlobalNavASearchWithNoPagePanics(t *testing.T) {
	mustPanic(t, "needs its page", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: site.FlyoutSearch{Action: "", Links: nil}, Bag: nil, Dark: false, Clear: false, Class: ""})
	})
}

func TestGlobalNavDarkClearAndThePagesUtilities(t *testing.T) {
	got := example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: nil, Dark: true, Clear: true, Class: "sticky top-0"})
	if !strings.HasPrefix(got, `<nav class="cui-globalnav cui-dark sticky top-0" data-variant="clear" aria-label="Global">`) {
		t.Errorf("got %s", got)
	}
	if strings.Contains(got, "cui-globalnav-icon") {
		t.Errorf("got %s", got)
	}
}

func TestGlobalNavEscapesItsText(t *testing.T) {
	got := example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/?a&b", LogoAriaLabel: `A "B"`, Links: []site.NavLink{{Text: "<b>", Href: "/x", Flyout: nil}}, Search: nil, Bag: nil, Dark: false, Clear: false, Class: ""})
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
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: nil, Search: nil, Bag: nil, Dark: false, Clear: false, Class: ""})
	})
}

func TestGlobalNavAnEmptyHomePanics(t *testing.T) {
	mustPanic(t, "GlobalNav needs a Home", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: nil, Dark: false, Clear: false, Class: ""})
	})
}

func TestGlobalNavAnEmptyLogoAriaLabelPanics(t *testing.T) {
	mustPanic(t, "GlobalNav needs a LogoAriaLabel", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "", Links: globalNavLinks, Search: nil, Bag: nil, Dark: false, Clear: false, Class: ""})
	})
}

func TestGlobalNavALinkWithNoTextPanics(t *testing.T) {
	mustPanic(t, "GlobalNav's Links need their Text and Href", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Text: "", Href: "/x", Flyout: nil}}, Search: nil, Bag: nil, Dark: false, Clear: false, Class: ""})
	})
}

func TestGlobalNavAFlyoutGroupWithNoHeadingPanics(t *testing.T) {
	flyout := []site.FlyoutGroup{{Heading: "", Large: nil, Links: nil}}
	mustPanic(t, "FlyoutGroups need their Heading", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Text: "Store", Href: "/store", Flyout: flyout}}, Search: nil, Bag: nil, Dark: false, Clear: false, Class: ""})
	})
}

func TestGlobalNavABagWithNoMessagePanics(t *testing.T) {
	bag := site.FlyoutBag{Message: "", Line: nil, Heading: "", Links: nil}
	mustPanic(t, "bag needs its Message", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: bag, Dark: false, Clear: false, Class: ""})
	})
}

func TestGlobalNavABagLineWithNoTextPanics(t *testing.T) {
	bag := site.FlyoutBag{Message: "Nothing here yet.", Line: &site.BagLine{Text: "", Href: "/signin", Rest: ""}, Heading: "", Links: nil}
	mustPanic(t, "Line needs its link's Text and Href", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: bag, Dark: false, Clear: false, Class: ""})
	})
}

func TestGlobalNavBagLinksWithNoHeadingPanics(t *testing.T) {
	bag := site.FlyoutBag{Message: "Nothing here yet.", Line: nil, Heading: "", Links: []site.BagLink{{Icon: icon.Icon{}, IconCompact: icon.Icon{}, Text: "Orders", Href: "/orders"}}}
	mustPanic(t, "bag's Links need their Heading", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: bag, Dark: false, Clear: false, Class: ""})
	})
}

func TestGlobalNavABagLinkWithNoTextPanics(t *testing.T) {
	bag := site.FlyoutBag{Message: "Nothing here yet.", Line: nil, Heading: "Account", Links: []site.BagLink{{Icon: icon.Icon{}, IconCompact: icon.Icon{}, Text: "", Href: "/orders"}}}
	mustPanic(t, "bag's Links need their Text and Href", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: globalNavLinks, Search: nil, Bag: bag, Dark: false, Clear: false, Class: ""})
	})
}

func TestGlobalNavLargeLinksAfterTheFirstGroupPanics(t *testing.T) {
	flyout := []site.FlyoutGroup{
		{Heading: "Shop", Large: nil, Links: nil},
		{Heading: "More", Large: []site.Link{{Text: "x", Href: "/x"}}, Links: nil},
	}
	mustPanic(t, "site.GlobalNav's Large links are only in a flyout's first group", func() {
		example.Render(site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Text: "Store", Href: "/store", Flyout: flyout}}, Search: nil, Bag: nil, Dark: false, Clear: false, Class: ""})
	})
}
