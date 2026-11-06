package site_test

import (
	"os"

	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func ExampleField() {
	page := h.Form(site.Field{ID: "email", Type: "email", Name: "email", Label: "Email", Autocomplete: "email", Value: "", Error: "", Prefix: "", Attrs: nil, Class: ""})
	_ = page.Render(os.Stdout)
	// Output: <form><div class="cui-field"><input class="cui-field-input" id="email" type="email" name="email" placeholder=" " autocomplete="email"><label class="cui-field-label" for="email">Email</label></div></form>
}

func ExampleAccordion() {
	page := site.Accordion{Items: []site.AccordionItem{{Question: "How long does delivery take?", Open: false, Class: "", Children: h.P(g.Text("Add your address at checkout to see a delivery date."))}}}
	_ = page.Render(os.Stdout)
	// Output: <details class="cui-accordion-item"><summary class="cui-accordion-question">How long does delivery take?<svg aria-hidden="true" width="17" height="17" viewBox="0 0 17 8.85"><path d="M15 1.13 8.5 7.72 2 1.13" fill="none" stroke="currentColor" stroke-width="2.25" stroke-linecap="round" stroke-linejoin="round"/></svg></summary><div class="cui-accordion-answer"><p>Add your address at checkout to see a delivery date.</p></div></details>
}

func ExampleFooter() {
	page := site.Footer{Copyright: g.Text("© 2026 Example Inc."), Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <footer class="cui-footer"><div class="cui-footer-content"><div class="cui-footer-legal"><p class="cui-footer-copyright">© 2026 Example Inc.</p></div></div></footer>
}

func ExampleGlobalNav() {
	logo := g.Raw(`<img src="/logo.svg" alt="" width="14" height="44">`)
	page := site.GlobalNav{Logo: logo, Home: "/", LogoAriaLabel: "Cider", Links: []site.NavLink{{Label: "Store", Href: "/store", Post: false}}, AriaLabel: "Global", MenuAriaLabel: "Menu", CloseAriaLabel: "Close", Class: "sticky top-0"}
	_ = page.Render(os.Stdout)
	// Output: <nav class="cui-globalnav sticky top-0" aria-label="Global"><div class="cui-globalnav-content"><ul class="cui-globalnav-list" id="cui-globalnav-menu" popover><li><a class="cui-globalnav-logo" href="/" aria-label="Cider"><img src="/logo.svg" alt="" width="14" height="44"></a></li><li><button class="cui-globalnav-closebutton" popovertarget="cui-globalnav-menu" popovertargetaction="hide" aria-label="Close" autofocus><svg aria-hidden="true" width="18" height="18" viewBox="0 0 18 18"><path d="M2 5H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/><path d="M2 12H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/></svg></button></li><li><a class="cui-globalnav-link" href="/store">Store</a></li></ul><button class="cui-globalnav-menubutton" popovertarget="cui-globalnav-menu" aria-label="Menu"><svg aria-hidden="true" width="18" height="18" viewBox="0 0 18 18"><path d="M2 5H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/><path d="M2 12H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/></svg></button></div></nav>
}

func ExampleOptionTile() {
	page := site.OptionTile{ID: "s-256", Name: "storage", Value: "256", Title: "256GB", Detail: "", Checked: false, Attrs: nil, Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-option"><input class="cui-option-input" type="radio" id="s-256" name="storage" value="256"><label class="cui-option-label" for="s-256"><span class="cui-option-title">256GB</span></label></div>
}

func ExampleRibbon() {
	page := site.Ribbon{Class: "", Children: h.P(g.Text("New this week. "), h.A(h.Class("cui-link"), h.Href("/new"), g.Text("See what's new")))}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-ribbon"><div class="cui-ribbon-content"><p>New this week. <a class="cui-link" href="/new">See what&#39;s new</a></p></div></div>
}
