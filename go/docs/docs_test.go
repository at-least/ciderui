package docs_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/docs"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func mustPanic(t *testing.T, want string, render func()) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil || !strings.Contains(fmt.Sprint(r), want) {
			t.Errorf("panicked with %v, want %q", r, want)
		}
	}()
	render()
}

func TestDocsAsideRendersTheExamples(t *testing.T) {
	example.Check(t, docs.DocsAside{Variant: docs.DocsAsideVariantNote, Class: "", Children: h.P(g.Text("A report from a beta build carries its beta identifier; one from the App Store does not."))}, "docs/docs-aside/examples/note.html")
	type kind struct {
		variant docs.DocsAsideVariant
		text    string
	}
	kinds := []kind{
		{variant: docs.DocsAsideVariantImportant, text: "Upload the symbol files of every build you ship."},
		{variant: docs.DocsAsideVariantWarning, text: "Deleting a project deletes its reports."},
		{variant: docs.DocsAsideVariantTip, text: "Mark a deploy to compare its crashes with the release before it."},
		{variant: docs.DocsAsideVariantDeprecated, text: "Use the store’s key in its place."},
		{variant: docs.DocsAsideVariantExperiment, text: "Try grouping by the crashed frame alone."},
	}
	example.Check(t, h.Div(h.Class("grid gap-4"), g.Map(kinds, func(k kind) g.Node {
		return docs.DocsAside{Variant: k.variant, Class: "", Children: h.P(g.Text(k.text))}
	})), "docs/docs-aside/examples/kinds.html")
}

func TestDocsAsideTakesThePagesUtilities(t *testing.T) {
	got := example.Render(docs.DocsAside{Variant: docs.DocsAsideVariantNote, Class: "mt-8", Children: h.P(g.Text("A."))})
	if !strings.HasPrefix(got, `<aside class="cui-docs-aside mt-8" aria-label="Note">`) {
		t.Errorf("got %s", got)
	}
}

func TestDocsBarRendersTheExample(t *testing.T) {
	example.Check(t, docs.DocsBar{Title: "Documentation", Href: "#documentation", Navigator: "docs-navigator", Class: ""}, "docs/docs-bar/examples/bar.html")
}

func TestDocsBarNoNavigatorNoButton(t *testing.T) {
	got := example.Render(docs.DocsBar{Title: "Documentation", Href: "/", Navigator: "", Class: ""})
	want := `<nav class="cui-docs-bar" aria-label="Documentation"><a class="cui-docs-bar-title" href="/">Documentation</a></nav>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestDocsBarAnEmptyTitlePanics(t *testing.T) {
	mustPanic(t, "needs a Title", func() {
		example.Render(docs.DocsBar{Title: "", Href: "/", Navigator: "", Class: ""})
	})
}

func TestDocsCopyRendersTheExample(t *testing.T) {
	// not g.Text: the example writes the listing's quotes bare, g.Text
	// &#34;, which example.Render makes &quot;
	listing := h.Pre(h.Code(g.Raw(`@import "ciderui/docs.css";`)))
	example.Check(t, h.Div(h.Class("cui-docs-content"), docs.DocsCopy{Class: "", Children: listing}), "docs/docs-copy/examples/copy.html")
}

func TestDocsHeroRendersTheExample(t *testing.T) {
	crumbs := []docs.DocsCrumb{
		{Label: "Crash reports", Href: "#crash-reports"},
		{Label: "Diagnosing crashes", Href: "#diagnosing"},
	}
	example.Check(t, docs.DocsHero{Crumbs: crumbs, Eyebrow: "Article", Title: "Reading a crash report", Abstract: "Find the thread that crashed, the frame that failed and the images it had loaded.", Class: ""}, "docs/docs-hero/examples/hero.html")
}

func TestDocsHeroAPageAtTheTopHasNoCrumbs(t *testing.T) {
	got := example.Render(docs.DocsHero{Crumbs: nil, Eyebrow: "", Title: "Get Started", Abstract: "", Class: ""})
	want := `<header class="cui-docs-hero"><h1 class="cui-docs-hero-title">Get Started</h1></header>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestDocsHeroAnEmptyTitlePanics(t *testing.T) {
	mustPanic(t, "needs a Title", func() {
		example.Render(docs.DocsHero{Crumbs: nil, Eyebrow: "", Title: "", Abstract: "", Class: ""})
	})
}

func page(label, href string, current, open bool, children []docs.DocsNavigatorItem) docs.DocsNavigatorItem {
	return docs.DocsNavigatorPage{Label: label, Href: href, Current: current, Open: open, Children: children}
}

func TestDocsNavigatorRendersTheExample(t *testing.T) {
	reports := []docs.DocsNavigatorItem{
		docs.DocsNavigatorGroup{Title: "Reports"},
		page("Reading a crash report", "#reading", true, false, nil),
		page("Symbolicating a backtrace", "#symbolicating", false, false, nil),
	}
	alerts := []docs.DocsNavigatorItem{page("Routing an alert", "#routing", false, false, nil)}
	items := []docs.DocsNavigatorItem{
		page("Getting started", "#getting-started", false, false, nil),
		page("Crash reports", "#crash-reports", false, true, reports),
		page("Alerts", "#alerts", false, false, alerts),
	}
	example.Check(t, docs.DocsNavigator{Title: "Crashes", Href: "#crashes", Items: items, Class: ""}, "docs/docs-navigator/examples/navigator.html")
}

func TestDocsNavigatorAChevronIsNamedAfterItsPagesPath(t *testing.T) {
	under := []docs.DocsNavigatorItem{page("Button", "https://example.test/components/button/", false, false, nil)}
	for _, href := range []string{"/components/", "https://example.test/components/"} {
		items := []docs.DocsNavigatorItem{page("Components", href, false, false, under)}
		got := example.Render(docs.DocsNavigator{Title: "T", Href: "/", Items: items, Class: ""})
		if !strings.Contains(got, `id="docs-navigator-components"`) {
			t.Errorf("%s: got %s", href, got)
		}
	}
}

func TestDocsNavigatorAnOpenPageWithNoPagesPanics(t *testing.T) {
	mustPanic(t, "is open with no pages under it", func() {
		items := []docs.DocsNavigatorItem{page("Alone", "/alone/", false, true, nil)}
		example.Render(docs.DocsNavigator{Title: "T", Href: "/", Items: items, Class: ""})
	})
}

func TestDocsNavigatorAnEmptyTitlePanics(t *testing.T) {
	mustPanic(t, "needs a Title", func() {
		example.Render(docs.DocsNavigator{Title: "", Href: "/", Items: nil, Class: ""})
	})
}

func TestDocsNavigatorANilItemPanics(t *testing.T) {
	mustPanic(t, "docs.DocsNavigator: a nil item", func() {
		example.Render(docs.DocsNavigator{Title: "T", Href: "/", Items: []docs.DocsNavigatorItem{nil}, Class: ""})
	})
}

func TestDocsNavigatorAPageWithAnEmptyLabelPanics(t *testing.T) {
	mustPanic(t, "a page needs a Label", func() {
		items := []docs.DocsNavigatorItem{page("", "/alone/", false, false, nil)}
		example.Render(docs.DocsNavigator{Title: "T", Href: "/", Items: items, Class: ""})
	})
}

func TestDocsPageRendersTheExample(t *testing.T) {
	reports := []docs.DocsNavigatorItem{
		page("Reading a crash report", "#reading", true, false, nil),
		page("Symbolicating a backtrace", "#symbolicating", false, false, nil),
	}
	navigator := []docs.DocsNavigatorItem{
		page("Getting started", "#getting-started", false, false, nil),
		page("Crash reports", "#crash-reports", false, true, reports),
	}
	crumbs := []docs.DocsCrumb{{Label: "Crash reports", Href: "#crash-reports"}}
	topics := []docs.DocsTopic{{Label: "Symbolicating a backtrace", Href: "#symbolicating", Abstract: "Replace the addresses in a backtrace with the names of the functions they fall in.", Collection: false}}
	groups := []docs.DocsTopicsGroup{{Heading: "", Topics: topics}}
	columns := g.Group{
		h.Div(h.Class("cui-docs-page-column"),
			docs.DocsHero{Crumbs: crumbs, Eyebrow: "Article", Title: "Reading a crash report", Abstract: "Find the thread that crashed, the frame that failed and the images it had loaded.", Class: ""},
		),
		h.Div(h.Class("cui-docs-page-column"),
			h.Div(h.Class("cui-docs-content"),
				h.H2(g.Text("Overview")),
				h.P(g.Text("A crash report records what an app was doing when it stopped.")),
			),
		),
		h.Div(h.Class("cui-docs-page-column"),
			docs.DocsTopics{Title: "See Also", Groups: groups, Class: ""},
		),
	}
	example.Check(t, docs.DocsPage{Title: "Crashes", Href: "#crashes", Navigator: navigator, Class: "", Children: columns}, "docs/docs-page/examples/page.html")
}

func TestDocsTopicsRendersTheExample(t *testing.T) {
	topics := []docs.DocsTopic{
		{Label: "Symbolicating a backtrace", Href: "#symbolicating", Abstract: "Replace the addresses in a backtrace with the names of the functions they fall in.", Collection: false},
		{Label: "Identifying the cause of common crashes", Href: "#common-crashes", Abstract: "Match a crash report to a pattern and find what to fix.", Collection: false},
		{Label: "Alerts", Href: "#alerts", Abstract: "Hear about a new crash where your team already is.", Collection: true},
	}
	groups := []docs.DocsTopicsGroup{{Heading: "Crash reports", Topics: topics}}
	example.Check(t, docs.DocsTopics{Title: "See Also", Groups: groups, Class: ""}, "docs/docs-topics/examples/topics.html")
}

func TestDocsTopicsAnEmptyTitlePanics(t *testing.T) {
	mustPanic(t, "needs a Title", func() {
		example.Render(docs.DocsTopics{Title: "", Groups: nil, Class: ""})
	})
}

func TestDocsTopicsAnEmptyLabelPanics(t *testing.T) {
	mustPanic(t, "a topic needs a Label", func() {
		groups := []docs.DocsTopicsGroup{{Heading: "", Topics: []docs.DocsTopic{{Label: "", Href: "/", Abstract: "", Collection: false}}}}
		example.Render(docs.DocsTopics{Title: "See Also", Groups: groups, Class: ""})
	})
}

func TestDocsNavigatorTwoChevronsOfOneIDPanics(t *testing.T) {
	under := []docs.DocsNavigatorItem{page("Button", "/components/button/", false, false, nil)}
	items := []docs.DocsNavigatorItem{
		page("Components", "/components/", false, false, under),
		page("Components again", "https://example.test/components/", false, false, under),
	}
	mustPanic(t, "docs.DocsNavigator: two pages share the chevron id docs-navigator-components", func() {
		example.Render(docs.DocsNavigator{Title: "T", Href: "/", Items: items, Class: ""})
	})
}
