package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
)

func TestArticleRendersTheExample(t *testing.T) {
	const body = `<div class="cui-article-body">`
	want := example.Read("site/article/examples/article.html")
	// the example's copy, as the page's markdown would give it
	start := strings.Index(want, body) + len(body)
	end := strings.LastIndex(want, "</div></article>")
	example.Check(t, site.Article{Category: "Update", Datetime: "2026-10-19", Date: "October 19, 2026", Headline: "ciderui comes to Zola", Subhead: "The kit's look, now for a static site's pages and its news.", Body: g.Raw(want[start:end]), Class: ""}, "site/article/examples/article.html")
}

func TestArticleNoEyebrowAndNoSubheadWriteNothing(t *testing.T) {
	got := example.Render(site.Article{Category: "", Datetime: "", Date: "", Headline: "Privacy", Subhead: "", Body: g.Raw("<p>x</p>"), Class: "pb-24"})
	want := `<article class="cui-article pb-24"><header class="cui-article-header"><h1 class="cui-article-headline">Privacy</h1></header><div class="cui-article-body"><p>x</p></div></article>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestArticleAnEmptyHeadlinePanics(t *testing.T) {
	mustPanic(t, "needs a Headline", func() {
		example.Render(site.Article{Category: "", Datetime: "", Date: "", Headline: "", Subhead: "", Body: nil, Class: ""})
	})
}

func TestArticleADateWithNoDatetimePanics(t *testing.T) {
	mustPanic(t, "Date needs its Datetime", func() {
		example.Render(site.Article{Category: "", Datetime: "", Date: "Today", Headline: "x", Subhead: "", Body: nil, Class: ""})
	})
}
