//go:build browser

// @invariant: What a site links: dist/<set>.css as it ships, and its own utilities'
// stylesheet after it.
package assets

import (
	"strconv"
	"testing"

	"github.com/at-least/ciderui/internal/browser"
	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/check"
	"github.com/mxschmitt/playwright-go"
)

func open(t *testing.T, b *browser.Browser, set, html string) *browser.Page {
	t.Helper()
	css, err := build.Shipped(set)
	if err != nil {
		t.Fatal(err)
	}
	p, err := b.Open(css, html, browser.Options{})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { p.Close() })
	return p
}

func one(t *testing.T, p *browser.Page, selector, prop string) string {
	t.Helper()
	got, err := p.Read(selector, []string{prop}, browser.ReadOptions{})
	if err != nil {
		t.Fatal(err)
	}
	return got[prop]
}

func TestAssets(t *testing.T) {
	own, err := build.OwnSheet("utilities.css", "px-8 md:px-8")
	if err != nil {
		t.Fatal(err)
	}
	b := browser.Launch(t)

	site := open(t, b, "site", `<h1 id="headline" class="cui-headline">a</h1><a id="plain" class="cui-btn">a</a><a id="own" class="cui-btn px-8">a</a><a id="fixed" class="cui-btn md:px-8">a</a>`)
	if _, err := site.Page.AddStyleTag(playwright.PageAddStyleTagOptions{Content: playwright.String(own)}); err != nil {
		t.Fatal(err)
	}
	check.Equal(t, "every class of the type ships: cui-headline", one(t, site, "#headline", "font-size"), "48px")
	check.Equal(t, "a part ships: cui-btn", one(t, site, "#plain", "padding-left"), "15px")
	check.Equal(t, "a utility of the second stylesheet beats a part", one(t, site, "#own", "padding-left"), "32px")
	check.Equal(t, "a breakpoint of the second stylesheet is the kit's", one(t, site, "#fixed", "padding-left"), "32px")
	if err := site.Resize(734); err != nil {
		t.Fatal(err)
	}
	check.Equal(t, "and steps at the kit's width", one(t, site, "#fixed", "padding-left"), "15px")

	console := open(t, b, "console", `<h2 id="h" class="cui-console-h2">a</h2><a id="btn" class="cui-console-btn">a</a>`)
	check.Equal(t, "a part's own type ships: cui-console-h2", one(t, console, "#h", "font-size"), "26px")
	check.Equal(t, "a part ships: cui-console-btn", strconv.FormatBool(one(t, console, "#btn", "display") != "inline"), "true")
}
