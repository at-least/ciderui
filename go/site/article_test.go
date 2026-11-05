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
	example.Check(t, site.Article{Category: "Update", Date: &site.Date{Datetime: "2026-10-19", Text: "October 19, 2026"}, Headline: "ciderui comes to Zola", Subhead: "The kit's look, now for a static site's pages and its news.", Body: g.Raw(want[start:end]), Class: ""}, "site/article/examples/article.html")
}

func TestArticleNoEyebrowAndNoSubheadWriteNothing(t *testing.T) {
	got := example.Render(site.Article{Category: "", Date: nil, Headline: "Privacy", Subhead: "", Body: g.Raw("<p>x</p>"), Class: "pb-24"})
	want := `<article class="cui-article pb-24"><header class="cui-article-header"><h1 class="cui-article-headline">Privacy</h1></header><div class="cui-article-body"><p>x</p></div></article>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestArticlePanics(t *testing.T) {
	example.Panics(t, func() site.Article {
		return site.Article{Category: "", Date: nil, Headline: "x", Subhead: "", Body: nil, Class: ""}
	}, []example.Break[site.Article]{
		{Want: "site.Article needs a Headline", Change: func(p *site.Article) { p.Headline = "" }},
		{Want: "site.Date needs a Datetime", Change: func(p *site.Article) { p.Date = &site.Date{Datetime: "", Text: "Today"} }},
		{Want: "site.Date needs a Text", Change: func(p *site.Article) { p.Date = &site.Date{Datetime: "2026-10-19", Text: ""} }},
	})
}

func TestArticleEyebrowOfACategoryOrADateAlone(t *testing.T) {
	date := &site.Date{Datetime: "2026-10-19", Text: "Today"}
	for _, c := range []struct {
		article site.Article
		want    string
	}{
		{site.Article{Category: "Update", Date: nil, Headline: "x", Subhead: "", Body: nil, Class: ""}, `<p class="cui-article-eyebrow"><span class="cui-article-category">Update</span></p>`},
		{site.Article{Category: "", Date: date, Headline: "x", Subhead: "", Body: nil, Class: ""}, `<p class="cui-article-eyebrow"><time class="cui-article-date" datetime="2026-10-19">Today</time></p>`},
	} {
		if got := example.Render(c.article); !strings.Contains(got, c.want) {
			t.Errorf("no %s in %s", c.want, got)
		}
	}
}
