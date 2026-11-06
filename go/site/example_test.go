package site_test

import (
	"os"

	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func ExampleField() {
	page := h.Form(site.Field{ID: "email", Type: "email", Name: "email", Label: "Email", Autocomplete: "email", Value: "", Help: "", Error: "", Prefix: "", Attrs: nil, Class: ""})
	_ = page.Render(os.Stdout)
	// Output: <form><div class="cui-field"><input class="cui-field-input" id="email" type="email" name="email" placeholder=" " autocomplete="email"><label class="cui-field-label" for="email">Email</label></div></form>
}

func ExampleAccordion() {
	page := site.Accordion{Large: false, Class: "", Items: []site.AccordionItem{{Question: "How long does delivery take?", Open: false, Class: "", Children: h.P(g.Text("Add your address at checkout to see a delivery date."))}}}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-accordion"><details class="cui-accordion-item"><summary class="cui-accordion-question">How long does delivery take?<svg aria-hidden="true" width="17" height="17" viewBox="0 0 17 8.85"><path d="M15 1.13 8.5 7.72 2 1.13" fill="none" stroke="currentColor" stroke-width="2.25" stroke-linecap="round" stroke-linejoin="round"/></svg></summary><div class="cui-accordion-answer"><p>Add your address at checkout to see a delivery date.</p></div></details></div>
}

func ExampleFooter() {
	directory := [][]site.FooterSection{{{ID: "f-shop", Title: "Shop", Links: []site.Link{{Label: "Store", Href: "/store/"}}}}}
	page := site.Footer{Notes: nil, NotesAriaLabel: "", Directory: directory, DirectoryAriaLabel: "Directory", Shop: nil, Copyright: g.Text("Copyright © 2026 Example Inc."), Links: []site.Link{{Label: "Privacy Policy", Href: "/privacy/"}}, Locale: nil, Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <footer class="cui-footer"><div class="cui-footer-content"><nav class="cui-footer-directory" aria-label="Directory"><div class="cui-footer-column"><div class="cui-footer-section"><input class="cui-footer-toggle" type="checkbox" id="f-shop"><h3 class="cui-footer-title"><label for="f-shop">Shop<svg aria-hidden="true" width="11" height="6" viewBox="0 0 11 6"><path d="M10.075.675 5.5 5.323.925.675" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round"/></svg></label></h3><ul class="cui-footer-list"><li><a href="/store/">Store</a></li></ul></div></div></nav><div class="cui-footer-legal"><div class="cui-footer-row"><p class="cui-footer-copyright">Copyright © 2026 Example Inc.</p><ul class="cui-footer-links"><li><a href="/privacy/">Privacy Policy</a></li></ul></div></div></div></footer>
}

func ExampleGlobalNav_flyouts() {
	laptops := []site.FlyoutGroup{{Heading: "Laptops", Large: []site.Link{{Label: "All Laptops", Href: "/laptops"}}, Links: []site.Link{{Label: "Compare Models", Href: "/laptops/compare"}}}}
	_ = []site.NavLink{
		{Label: "Laptops", Href: "/laptops", Action: site.NavFlyout{Groups: laptops, AriaLabel: "Laptops menu"}},
		{Label: "Sign Out", Href: "/signout", Action: site.NavPost{}},
	}
	_ = site.FlyoutSearch{Action: "/search", AriaLabel: "Search", Placeholder: "Search", ClearAriaLabel: "Clear search", Heading: "", Links: nil}
	_ = site.FlyoutBag{AriaLabel: "Bag", Message: "Nothing here yet.", Line: nil, Heading: "Account", Links: []site.BagLink{{Icon: site.BagLinkIconOrders, Label: "Orders", Href: "/orders"}}}
}

func ExampleGlobalNav() {
	logo := g.Raw(`<img src="/logo.svg" alt="" width="14" height="44">`)
	page := site.GlobalNav{Logo: logo, Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Label: "Store", Href: "/store", Action: nil}}, Search: nil, Bag: nil, Text: site.GlobalNavText{AriaLabel: "Global", MenuAriaLabel: "Menu", CloseAriaLabel: "Close", MainMenuAriaLabel: ""}, Clear: false, Class: "sticky top-0"}
	_ = page.Render(os.Stdout)
	// Output: <nav class="cui-globalnav sticky top-0" aria-label="Global"><div class="cui-globalnav-content"><ul class="cui-globalnav-list" id="cui-globalnav-menu" popover><li><a class="cui-globalnav-logo" href="/" aria-label="Cider"><img src="/logo.svg" alt="" width="14" height="44"></a></li><li><button class="cui-globalnav-closebutton" popovertarget="cui-globalnav-menu" popovertargetaction="hide" aria-label="Close" autofocus><svg aria-hidden="true" width="18" height="18" viewBox="0 0 18 18"><path d="M2 5H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/><path d="M2 12H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/></svg></button></li><li><a class="cui-globalnav-link" href="/store">Store</a></li></ul><button class="cui-globalnav-menubutton" popovertarget="cui-globalnav-menu" aria-label="Menu"><svg aria-hidden="true" width="18" height="18" viewBox="0 0 18 18"><path d="M2 5H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/><path d="M2 12H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/></svg></button></div></nav>
}

func ExampleMoreLink() {
	page := site.MoreLink{Href: "#learn", Label: "Learn more", Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <a class="cui-link" data-variant="more" href="#learn"><span class="cui-link-text">Learn more</span><svg class="cui-link-icon" aria-hidden="true" width="0.442em" height="0.54em" viewBox="0 0 442 540"><path d="M107.7 47.7 332.3 272.5 107.7 497.3" fill="none" stroke="currentColor" stroke-width="67.9" stroke-linecap="round" stroke-linejoin="round"/></svg></a>
}

func ExampleOptionTile() {
	page := site.OptionTile{ID: "s-256", Name: "storage", Value: "256", Title: "256GB", Detail: "", Price: "From $1099", Checked: false, Attrs: nil, Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-option"><input class="cui-option-input" type="radio" id="s-256" name="storage" value="256"><label class="cui-option-label" for="s-256"><span class="cui-option-title">256GB</span><span class="cui-option-price">From $1099</span></label></div>
}

func ExampleRibbon() {
	page := site.Ribbon{Class: "", Children: h.P(g.Text("New this week. "), site.MoreLink{Href: "/new", Label: "See what's new", Class: ""})}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-ribbon"><div class="cui-ribbon-content"><p>New this week. <a class="cui-link" data-variant="more" href="/new"><span class="cui-link-text">See what&#39;s new</span><svg class="cui-link-icon" aria-hidden="true" width="0.442em" height="0.54em" viewBox="0 0 442 540"><path d="M107.7 47.7 332.3 272.5 107.7 497.3" fill="none" stroke="currentColor" stroke-width="67.9" stroke-linecap="round" stroke-linejoin="round"/></svg></a></p></div></div>
}
