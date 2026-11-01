package docs_test

import (
	"os"

	"github.com/at-least/ciderui/go/docs"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func ExampleDocsAside() {
	_ = docs.DocsAside{Variant: docs.DocsAsideVariantTip, Class: "", Children: h.P(g.Text("Mark a deploy."))}.Render(os.Stdout)
	// Output: <aside class="cui-docs-aside" data-variant="tip" aria-label="Tip"><p class="cui-docs-aside-label">Tip</p><p>Mark a deploy.</p></aside>
}

func ExampleDocsBar() {
	_ = docs.DocsBar{Title: "Documentation", Href: "/", Navigator: "docs-navigator", Class: ""}.Render(os.Stdout)
	// Output: <nav class="cui-docs-bar" aria-label="Documentation"><button class="cui-docs-bar-toggle" type="button" popovertarget="docs-navigator" aria-label="Navigator"><svg aria-hidden="true" width="19" height="19" viewBox="0 0 14 14"><rect x="1.1125" y="2.4875" width="11.775" height="9.025" rx="1.27" fill="none" stroke="currentColor" stroke-width="0.975"/><path d="M5.3 2.4875V11.5125" fill="none" stroke="currentColor" stroke-width="0.95"/><path d="M2.653 4.635H3.797M2.653 6.056H3.797M2.653 7.473H3.797" fill="none" stroke="currentColor" stroke-width="0.656" stroke-linecap="round"/></svg></button><a class="cui-docs-bar-title" href="/">Documentation</a></nav>
}

func ExampleDocsCopy() {
	_ = docs.DocsCopy{Class: "", Children: h.Pre(h.Code(g.Text("cargo build")))}.Render(os.Stdout)
	// Output: <div class="cui-docs-copy"><pre><code>cargo build</code></pre><button class="cui-docs-copy-button" type="button" aria-label="Copy" hidden><svg class="cui-docs-copy-idle" aria-hidden="true" width="25.5" height="25.5" viewBox="0 0 24 24"><path d="M16 2H4A1 1 0 0 0 3 3V17" fill="none" stroke="currentColor" stroke-width="2"/><rect x="7" y="6" width="13" height="16" rx="1" fill="none" stroke="currentColor" stroke-width="2"/></svg><svg class="cui-docs-copy-done" aria-hidden="true" width="25.5" height="25.5" viewBox="0 0 24 24"><path d="M4.1 12.7 9 17.6 20.3 6.3" fill="none" stroke="currentColor" stroke-width="1.98"/></svg></button></div>
}

func ExampleDocsHero() {
	crumbs := []docs.DocsCrumb{{Label: "Components", Href: "/components/"}}
	_ = docs.DocsHero{Crumbs: crumbs, Eyebrow: "Component", Title: "Docs Hero", Abstract: "The top of a page.", Class: ""}.Render(os.Stdout)
	// Output: <header class="cui-docs-hero"><nav aria-label="Breadcrumbs"><ol class="cui-docs-hero-crumbs"><li><a href="/components/">Components</a></li><li aria-current="page">Docs Hero</li></ol></nav><p class="cui-docs-hero-eyebrow">Component</p><h1 class="cui-docs-hero-title">Docs Hero</h1><p class="cui-docs-hero-abstract">The top of a page.</p></header>
}

func ExampleDocsNavigator() {
	items := []docs.DocsNavigatorItem{docs.DocsNavigatorPage{Label: "Get Started", Href: "/get-started/", Current: true, Open: false, Children: nil}}
	_ = docs.DocsNavigator{Title: "ciderui", Href: "/", Items: items, Class: ""}.Render(os.Stdout)
	// Output: <nav class="cui-docs-navigator" aria-label="Navigator"><a class="cui-docs-navigator-title" href="/">ciderui</a><ul class="cui-docs-navigator-list"><li class="cui-docs-navigator-item"><div class="cui-docs-navigator-row"><svg class="cui-docs-navigator-icon" aria-hidden="true" width="17" height="17" viewBox="0 0 14 14"><path d="M7.515 0.525H3.475A1.35 1.35 0 0 0 2.125 1.875V12.125A1.35 1.35 0 0 0 3.475 13.475H10.725A1.35 1.35 0 0 0 12.075 12.125V5.085Z" fill="none" stroke="currentColor" stroke-width="1.05" stroke-linejoin="round"/><path d="M7.09 0.6V4.9A0.69 0.69 0 0 0 7.78 5.59H12" fill="none" stroke="currentColor" stroke-width="0.975"/><path d="M4.53 8.28H9.5M4.53 10.57H9.5" fill="none" stroke="currentColor" stroke-width="0.8" stroke-linecap="round"/></svg><a class="cui-docs-navigator-link" href="/get-started/" aria-current="page">Get Started</a></div></li></ul><p class="cui-docs-navigator-empty" hidden>No results found.</p><div class="cui-docs-navigator-filter" role="search" hidden><label class="cui-docs-navigator-field"><svg class="cui-docs-navigator-filter-icon" aria-hidden="true" width="21" height="21" viewBox="0 0 21 21"><path d="M2.75 6.3H18.3M4.6 10.5H16.4M6.5 14.7H14.7" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round"/></svg><input class="cui-docs-navigator-input" type="search" placeholder="Filter" aria-label="Filter" autocomplete="off" spellcheck="false"></label></div><p class="cui-docs-navigator-status" aria-live="polite"></p></nav>
}

func ExampleDocsPage() {
	_ = docs.DocsPage{Title: "Documentation", Href: "/", Navigator: nil, Class: "", Children: h.Div(h.Class("cui-docs-page-column"), g.Text("…"))}.Render(os.Stdout)
	// Output: <div class="cui-docs-page"><nav class="cui-docs-bar" aria-label="Documentation"><button class="cui-docs-bar-toggle" type="button" popovertarget="docs-navigator" aria-label="Navigator"><svg aria-hidden="true" width="19" height="19" viewBox="0 0 14 14"><rect x="1.1125" y="2.4875" width="11.775" height="9.025" rx="1.27" fill="none" stroke="currentColor" stroke-width="0.975"/><path d="M5.3 2.4875V11.5125" fill="none" stroke="currentColor" stroke-width="0.95"/><path d="M2.653 4.635H3.797M2.653 6.056H3.797M2.653 7.473H3.797" fill="none" stroke="currentColor" stroke-width="0.656" stroke-linecap="round"/></svg></button><a class="cui-docs-bar-title" href="/">Documentation</a></nav><div class="cui-docs-page-sidebar" id="docs-navigator" popover><button class="cui-docs-page-close" type="button" popovertarget="docs-navigator" popovertargetaction="hide" aria-label="Close Navigator"><svg aria-hidden="true" width="19" height="19" viewBox="0 0 14 14"><path d="M1.79 1.79 11.235 11.235M11.235 1.79 1.79 11.235" fill="none" stroke="currentColor" stroke-width="1.58" stroke-linecap="round"/></svg></button><nav class="cui-docs-navigator" aria-label="Navigator"><a class="cui-docs-navigator-title" href="/">Documentation</a><ul class="cui-docs-navigator-list"></ul><p class="cui-docs-navigator-empty" hidden>No results found.</p><div class="cui-docs-navigator-filter" role="search" hidden><label class="cui-docs-navigator-field"><svg class="cui-docs-navigator-filter-icon" aria-hidden="true" width="21" height="21" viewBox="0 0 21 21"><path d="M2.75 6.3H18.3M4.6 10.5H16.4M6.5 14.7H14.7" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round"/></svg><input class="cui-docs-navigator-input" type="search" placeholder="Filter" aria-label="Filter" autocomplete="off" spellcheck="false"></label></div><p class="cui-docs-navigator-status" aria-live="polite"></p></nav></div><main class="cui-docs-page-main"><div class="cui-docs-page-column">…</div></main></div>
}

func ExampleDocsTopics() {
	topics := []docs.DocsTopic{{Label: "Docs Bar", Href: "/components/docs-bar/", Abstract: "", Collection: false}}
	groups := []docs.DocsTopicsGroup{{Heading: "", Topics: topics}}
	_ = docs.DocsTopics{Title: "See Also", Groups: groups, Class: ""}.Render(os.Stdout)
	// Output: <section class="cui-docs-topics"><h2 class="cui-docs-topics-title">See Also</h2><div class="cui-docs-topics-group"><div class="cui-docs-topics-item"><a class="cui-docs-topics-link" href="/components/docs-bar/"><svg class="cui-docs-topics-icon" aria-hidden="true" width="17" height="17" viewBox="0 0 14 14"><path d="M7.515 0.525H3.475A1.35 1.35 0 0 0 2.125 1.875V12.125A1.35 1.35 0 0 0 3.475 13.475H10.725A1.35 1.35 0 0 0 12.075 12.125V5.085Z" fill="none" stroke="currentColor" stroke-width="1.05" stroke-linejoin="round"/><path d="M7.09 0.6V4.9A0.69 0.69 0 0 0 7.78 5.59H12" fill="none" stroke="currentColor" stroke-width="0.975"/><path d="M4.53 8.28H9.5M4.53 10.57H9.5" fill="none" stroke="currentColor" stroke-width="0.8" stroke-linecap="round"/></svg><span>Docs Bar</span></a></div></div></section>
}
