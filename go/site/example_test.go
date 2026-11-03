package site_test

import (
	"fmt"
	"os"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func ExampleField() {
	page := h.Form(site.Field{ID: "email", Type: "email", Name: "email", Label: "Email", Autocomplete: "email", Value: "", Error: "", Disabled: false, Prefix: "", Class: ""})
	_ = page.Render(os.Stdout)
	// Output: <form><div class="cui-field"><input class="cui-field-input" id="email" type="email" name="email" placeholder=" " autocomplete="email"><label class="cui-field-label" for="email">Email</label></div></form>
}

func ExampleAccordion() {
	page := site.Accordion{Large: false, Class: "", Children: site.AccordionItem{Question: "How long does delivery take?", Open: false, Class: "", Children: h.P(g.Text("Add your address at checkout to see a delivery date."))}}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-accordion"><details class="cui-accordion-item"><summary class="cui-accordion-question">How long does delivery take?<svg aria-hidden="true" width="17" height="17" viewBox="0 0 17 8.85"><path d="M15 1.13 8.5 7.72 2 1.13" fill="none" stroke="currentColor" stroke-width="2.25" stroke-linecap="round" stroke-linejoin="round"/></svg></summary><div class="cui-accordion-answer"><p>Add your address at checkout to see a delivery date.</p></div></details></div>
}

func ExampleArticle() {
	page := site.Article{Category: "Update", Datetime: "2026-10-19", Date: "October 19, 2026", Headline: "ciderui comes to Zola", Subhead: "", Body: g.Raw("<p>The copy.</p>"), Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <article class="cui-article"><header class="cui-article-header"><p class="cui-article-eyebrow"><span class="cui-article-category">Update</span><time class="cui-article-date" datetime="2026-10-19">October 19, 2026</time></p><h1 class="cui-article-headline">ciderui comes to Zola</h1></header><div class="cui-article-body"><p>The copy.</p></div></article>
}

func ExampleCard() {
	page := site.Card{Dark: true, Href: "/phone/", Media: nil, Eyebrow: "NEW", Tag: true, Headline: "Cider Phone", Description: "", Price: "From $799", Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-card cui-dark"><div class="cui-card-content"><p class="cui-card-eyebrow" data-variant="tag">NEW</p><h3 class="cui-card-headline"><a class="cui-card-link" href="/phone/">Cider Phone</a></h3><p class="cui-card-price">From $799</p></div></div>
}

func ExampleCheckbox() {
	page := site.Checkbox{ID: "remember", Name: "remember", Label: "Remember me", Checked: false, Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-checkbox"><input class="cui-checkbox-input" type="checkbox" id="remember" name="remember"><label class="cui-checkbox-label" for="remember"><span class="cui-checkbox-indicator" aria-hidden="true"><svg aria-hidden="true" width="14" height="14" viewBox="0 0 14 14"><path d="M3.3 7.65 6.25 11.35 11.15 3.7" fill="none" stroke="currentColor" stroke-width="1.35" stroke-linecap="round" stroke-linejoin="round"/></svg></span>Remember me</label></div>
}

func ExampleFooter() {
	directory := [][]site.FooterSection{{{ID: "f-shop", Title: "Shop", Links: []site.Link{{Text: "Store", Href: "/store/"}}}}}
	page := site.Footer{Notes: nil, Directory: directory, Shop: nil, Copyright: "Copyright © 2026 Example Inc.", Links: []site.Link{{Text: "Privacy Policy", Href: "/privacy/"}}, Locale: nil, Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <footer class="cui-footer"><div class="cui-footer-content"><nav class="cui-footer-directory" aria-label="Directory"><div class="cui-footer-column"><div class="cui-footer-section"><input class="cui-footer-toggle" type="checkbox" id="f-shop"><h3 class="cui-footer-title"><label for="f-shop">Shop<svg aria-hidden="true" width="11" height="6" viewBox="0 0 11 6"><path d="M10.075.675 5.5 5.323.925.675" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round"/></svg></label></h3><ul class="cui-footer-list"><li><a href="/store/">Store</a></li></ul></div></div></nav><div class="cui-footer-legal"><div class="cui-footer-row"><p class="cui-footer-copyright">Copyright © 2026 Example Inc.</p><ul class="cui-footer-links"><li><a href="/privacy/">Privacy Policy</a></li></ul></div></div></div></footer>
}

func ExampleFootnote() {
	page := h.P(g.Text("A full day"), site.Footnote{Href: "#footnote-1", Mark: "1", Class: ""})
	_ = page.Render(os.Stdout)
	// Output: <p>A full day<sup class="cui-footnote"><a href="#footnote-1" aria-label="Footnote 1">1</a></sup></p>
}

func ExampleGallery() {
	page := site.Gallery{ID: "camera", Paddles: false, Dots: 2, Current: 1, Class: "", Children: g.Group{
		h.Li(h.ID("camera-1"), g.Text("One")),
		h.Li(h.ID("camera-2"), g.Text("Two")),
	}}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-gallery"><ul class="cui-gallery-scroller" id="camera"><li id="camera-1">One</li><li id="camera-2">Two</li></ul><ul class="cui-gallery-dots"><li><a class="cui-gallery-dot" href="#camera-1" aria-label="Item 1" aria-current="true"></a></li><li><a class="cui-gallery-dot" href="#camera-2" aria-label="Item 2"></a></li></ul></div>
}

func ExampleNavLink() {
	laptops := []site.FlyoutGroup{{Heading: "Laptops", Large: []site.Link{{Text: "All Laptops", Href: "/laptops"}}, Links: []site.Link{{Text: "Compare Models", Href: "/laptops/compare"}}}}
	links := []site.NavLink{{Text: "Store", Href: "/store", Flyout: nil}, {Text: "Laptops", Href: "/laptops", Flyout: laptops}}
	fmt.Println(len(links[1].Flyout))
	// Output: 1
}

func ExampleBagLink() {
	orders := site.BagLink{Icon: icon.Orders, IconCompact: icon.OrdersCompact, Text: "Orders", Href: "/orders"}
	fmt.Println(orders.Text)
	// Output: Orders
}

func ExampleGlobalNav() {
	logo := g.Raw(`<img src="/logo.svg" alt="" width="14" height="44">`)
	page := site.GlobalNav{Logo: logo, Home: "/", Name: "Cider", Links: []site.NavLink{{Text: "Store", Href: "/store", Flyout: nil}}, Search: nil, Bag: nil, Dark: false, Clear: false, Class: "sticky top-0"}
	_ = page.Render(os.Stdout)
	// Output: <nav class="cui-globalnav sticky top-0" aria-label="Global"><div class="cui-globalnav-content"><ul class="cui-globalnav-list" id="cui-globalnav-menu" popover><li><a class="cui-globalnav-logo" href="/" aria-label="Cider"><img src="/logo.svg" alt="" width="14" height="44"></a></li><li><button class="cui-globalnav-closebutton" popovertarget="cui-globalnav-menu" popovertargetaction="hide" aria-label="Close" autofocus><svg aria-hidden="true" width="18" height="18" viewBox="0 0 18 18"><path d="M2 5H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/><path d="M2 12H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/></svg></button></li><li><a class="cui-globalnav-link" href="/store">Store</a></li></ul><button class="cui-globalnav-menubutton" popovertarget="cui-globalnav-menu" aria-label="Menu"><svg aria-hidden="true" width="18" height="18" viewBox="0 0 18 18"><path d="M2 5H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/><path d="M2 12H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/></svg></button></div></nav>
}

func ExampleMoreLink() {
	page := site.MoreLink{Href: "#learn", Text: "Learn more", Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <a class="cui-link" data-variant="more" href="#learn"><span class="cui-link-text">Learn more</span><svg class="cui-link-icon" aria-hidden="true" width="0.442em" height="0.54em" viewBox="0 0 442 540"><path d="M107.7 47.7 332.3 272.5 107.7 497.3" fill="none" stroke="currentColor" stroke-width="67.9" stroke-linecap="round" stroke-linejoin="round"/></svg></a>
}

func ExampleLocalNav() {
	page := site.LocalNav{Title: "Cider Book 13", Href: "/cider-book/", Links: []site.Link{{Text: "Overview", Href: "/cider-book/"}}, Current: "/cider-book/", Button: &site.Link{Text: "Buy", Href: "/buy/"}, Dark: false, Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <nav class="cui-localnav" aria-label="Local"><div class="cui-localnav-content"><a class="cui-localnav-title" href="/cider-book/">Cider Book 13</a><input class="cui-localnav-toggle" type="checkbox" id="cui-localnav-menu" aria-label="Menu"><div class="cui-localnav-menu"><ul class="cui-localnav-list"><li><a class="cui-localnav-link" href="/cider-book/" aria-current="page">Overview</a></li></ul></div><div class="cui-localnav-actions"><label class="cui-localnav-menubutton" for="cui-localnav-menu" aria-hidden="true"><svg aria-hidden="true" width="16" height="9" viewBox="0 0 16 9"><path d="M15.265.835 8 8.167.735.835" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"/></svg></label><a class="cui-btn" data-size="sm" href="/buy/">Buy</a></div></div></nav>
}

func ExampleModal() {
	page := g.Group{
		site.ModalButton{Modal: "models", Class: "cui-btn", Children: g.Text("Show more")},
		site.Modal{ID: "models", Labelledby: "models-title", Page: false, Dark: false, Class: "", Children: h.H2(h.ID("models-title"), g.Text("Not sure which size to get?"))},
	}
	_ = page.Render(os.Stdout)
	// Output: <button class="cui-btn" type="button" commandfor="models" command="show-modal">Show more</button><dialog class="cui-modal" id="models" aria-labelledby="models-title"><div class="cui-modal-panel"><form method="dialog"><button class="cui-modal-close" aria-label="Close"><svg aria-hidden="true" width="20" height="20" viewBox="0 0 20 20"><path d="M4.88 4.88 15.12 15.12M15.12 4.88 4.88 15.12" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round"/></svg></button></form><h2 id="models-title">Not sure which size to get?</h2></div></dialog>
}

func ExampleOptionTile() {
	page := site.OptionTile{ID: "s-256", Name: "storage", Value: "256", Title: "256GB", Detail: "", Price: "From $1099", Checked: false, Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-option"><input class="cui-option-input" type="radio" id="s-256" name="storage" value="256"><label class="cui-option-label" for="s-256"><span class="cui-option-title">256GB</span><span class="cui-option-price">From $1099</span></label></div>
}

func ExamplePosts() {
	page := site.Posts{Size: site.PostSize{}, Heading: "October 2026", Class: "", Children: h.Li(g.Text("…"))}
	_ = page.Render(os.Stdout)
	// Output: <section class="cui-posts"><h2 class="cui-posts-heading">October 2026</h2><ul class="cui-posts-list"><li>…</li></ul></section>
}

func ExamplePost() {
	page := site.Post{Href: "/news/zola/", Media: nil, Category: "Update", Datetime: "2026-10-19", Date: "October 19, 2026", Headline: "ciderui comes to Zola", Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <a class="cui-post" href="/news/zola/"><div class="cui-post-content"><p class="cui-post-category">Update</p><h3 class="cui-post-headline">ciderui comes to Zola</h3><time class="cui-post-date" datetime="2026-10-19">October 19, 2026</time></div></a>
}

func ExampleProductNav() {
	page := site.ProductNav{ID: "products", Label: "Products", Paddles: false, Class: "", Children: site.ProductNavItem{Href: "/shop/laptops", Name: "Laptops", Class: "",
		Children: h.Img(h.Class("cui-productnav-image"), h.Src("/laptops.png"), h.Width("200"), h.Height("130"), h.Alt("")),
	}}
	_ = page.Render(os.Stdout)
	// Output: <nav class="cui-shelf cui-productnav" aria-label="Products"><div class="cui-shelf-row"><ul class="cui-shelf-scroller" id="products"><li><a class="cui-productnav-item" href="/shop/laptops"><img class="cui-productnav-image" src="/laptops.png" width="200" height="130" alt=""><span class="cui-productnav-name">Laptops</span></a></li></ul></div></nav>
}

func ExampleQuickLinks() {
	page := site.QuickLinks{ID: "quicklinks", Title: "Shortcuts", Links: []site.Link{{Text: "Visit a Store", Href: "/retail/"}}, Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <nav class="cui-quicklinks" aria-labelledby="quicklinks-title"><h2 class="cui-quicklinks-title" id="quicklinks-title">Shortcuts</h2><ul class="cui-quicklinks-list"><li><a class="cui-btn" data-variant="secondary-neutral" href="/retail/">Visit a Store</a></li></ul></nav>
}

func ExampleRibbon() {
	page := site.Ribbon{Class: "", Children: h.P(g.Text("New this week. "), site.MoreLink{Href: "/new", Text: "See what's new", Class: ""})}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-ribbon"><div class="cui-ribbon-content"><p>New this week. <a class="cui-link" data-variant="more" href="/new"><span class="cui-link-text">See what&#39;s new</span><svg class="cui-link-icon" aria-hidden="true" width="0.442em" height="0.54em" viewBox="0 0 442 540"><path d="M107.7 47.7 332.3 272.5 107.7 497.3" fill="none" stroke="currentColor" stroke-width="67.9" stroke-linecap="round" stroke-linejoin="round"/></svg></a></p></div></div>
}

func ExampleSection() {
	page := site.Section{Ground: site.GroundAlt, Rhythm: site.RhythmLarge, Dark: true, Class: "", Children: h.P(g.Text("Content"))}
	_ = page.Render(os.Stdout)
	// Output: <section class="cui-section cui-dark" data-variant="alt" data-size="lg"><div class="cui-section-content"><p>Content</p></div></section>
}

func ExampleSectionHeader() {
	page := site.SectionHeader{Headline: "See it in action.", Links: []site.Link{{Text: "Watch the video", Href: "/film"}}, Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <header class="cui-section-header"><h2 class="cui-section-headline">See it in action.</h2><ul class="cui-section-links"><li><a class="cui-link" data-variant="more" href="/film"><span class="cui-link-text">Watch the video</span><svg class="cui-link-icon" aria-hidden="true" width="0.442em" height="0.54em" viewBox="0 0 442 540"><path d="M107.7 47.7 332.3 272.5 107.7 497.3" fill="none" stroke="currentColor" stroke-width="67.9" stroke-linecap="round" stroke-linejoin="round"/></svg></a></li></ul></header>
}

func ExampleShelf() {
	page := site.Shelf{ID: "latest", Headline: "New arrivals.", Subhead: "See what just came in.", Paddles: false, Class: "",
		Children: h.Li(site.Card{Dark: false, Href: "/phone/", Media: nil, Eyebrow: "", Tag: false, Headline: "Cider Phone", Description: "", Price: "From $799", Class: ""}),
	}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-shelf"><div class="cui-shelf-header"><h2 class="cui-shelf-headline">New arrivals.</h2> <span class="cui-shelf-subhead">See what just came in.</span></div><div class="cui-shelf-row"><ul class="cui-shelf-scroller" id="latest"><li><div class="cui-card"><div class="cui-card-content"><h3 class="cui-card-headline"><a class="cui-card-link" href="/phone/">Cider Phone</a></h3><p class="cui-card-price">From $799</p></div></div></li></ul></div></div>
}

func ExampleStat() {
	page := site.Stat{Large: false, Before: "Up to", Value: "2 days", After: "on a single charge", Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <p class="cui-stat"><span class="cui-stat-caption">Up to</span><span class="cui-stat-value">2 days</span><span class="cui-stat-caption">on a single charge</span></p>
}

func ExampleStories() {
	page := site.Stories{Heading: "Latest News", Class: "", Children: h.Li(g.Text("…"))}
	_ = page.Render(os.Stdout)
	// Output: <section class="cui-stories"><div class="cui-stories-content"><h2 class="cui-stories-heading">Latest News</h2><ul class="cui-stories-list"><li>…</li></ul></div></section>
}

func ExampleStory() {
	media := h.Img(h.Class("cui-story-media"), h.Src("/zola.jpg"), h.Alt(""))
	page := site.Story{Size: site.StorySizeLarge, Href: "/news/zola/", Media: media, Category: "Update", Datetime: "2026-10-19", Date: "October 19, 2026", Headline: "ciderui comes to Zola", Class: ""}
	_ = page.Render(os.Stdout)
	// Output: <a class="cui-story" data-size="lg" href="/news/zola/"><img class="cui-story-media" src="/zola.jpg" alt=""><div class="cui-story-content"><div class="cui-story-head"><p class="cui-story-category">Update</p><h3 class="cui-story-headline">ciderui comes to Zola</h3></div><time class="cui-story-date" datetime="2026-10-19">October 19, 2026</time></div></a>
}

func ExampleSwatch() {
	page := h.Ul(h.Class("cui-swatches"),
		site.Swatch{ID: "c-silver", Name: "color", Value: "silver", Label: "Silver", Checked: true, Class: "", Children: h.Img(h.Class("cui-swatch-color"), h.Src("/silver.png"), h.Alt(""))},
	)
	_ = page.Render(os.Stdout)
	// Output: <ul class="cui-swatches"><li class="cui-swatch"><input class="cui-swatch-input" type="radio" id="c-silver" name="color" value="silver" checked><label class="cui-swatch-ring" for="c-silver"><img class="cui-swatch-color" src="/silver.png" alt=""><span class="cui-swatch-name">Silver</span></label></li></ul>
}

func ExampleTile() {
	page := site.Tile{Size: site.TileSizeLarge, Rounded: false, Dark: true, Link: &site.Link{Text: "Cider Phone", Href: "/phone/"}, Media: nil, Headline: "Cider Phone", Subhead: "Built for every day.", Callout: "", Class: "",
		Children: h.A(h.Class("cui-btn"), h.Data("size", "cta"), h.Href("/phone/"), g.Text("Learn more")),
	}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-tile cui-dark" data-size="lg"><a class="cui-tile-link" href="/phone/" aria-label="Cider Phone"></a><div class="cui-tile-content"><h2 class="cui-tile-headline">Cider Phone</h2><p class="cui-tile-subhead">Built for every day.</p><div class="cui-tile-ctas"><a class="cui-btn" data-size="cta" href="/phone/">Learn more</a></div></div></div>
}
