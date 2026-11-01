package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestPostsRendersTheExamples(t *testing.T) {
	const picture = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMjgwIiBoZWlnaHQ9IjcyMCIvPg=="
	type row struct{ href, category, datetime, date, headline string }
	archive := []row{
		{"/news/zola/", "Update", "2026-10-19", "October 19, 2026", "ciderui comes to Zola"},
		{"/news/article/", "Update", "2026-10-18", "October 18, 2026", "An article in the Newsroom's look"},
		{"/news/stories/", "Feature", "2026-10-17", "October 17, 2026", "A grid of stories"},
	}
	related := []row{
		{"/news/archive/", "Update", "2026-10-16", "October 16, 2026", "The archive's rows"},
		{"/news/fonts/", "Update", "2026-10-15", "October 15, 2026", "Inter where SF is not"},
	}
	for _, c := range []struct {
		size    site.PostSize
		heading string
		rows    []row
		file    string
	}{
		{site.PostSize{}, "October 2026", archive, "site/posts/examples/archive.html"},
		{site.PostSizeSmall, "More from the Newsroom", related, "site/posts/examples/related.html"},
	} {
		example.Check(t, site.Posts{Size: c.size, Heading: c.heading, Class: "", Children: g.Map(c.rows, func(r row) g.Node {
			media := h.Img(h.Class("cui-post-media"), h.Src(picture), h.Alt(""), h.Width("1280"), h.Height("720"))
			return h.Li(site.Post{Href: r.href, Media: media, Category: r.category, Datetime: r.datetime, Date: r.date, Headline: r.headline, Class: ""})
		})}, c.file)
	}
}

func TestPostADateWithNoDatetimePanics(t *testing.T) {
	mustPanic(t, "Date needs its Datetime", func() {
		example.Render(site.Post{Href: "/x/", Media: nil, Category: "", Datetime: "", Date: "Today", Headline: "x", Class: ""})
	})
}

func TestPostsAnEmptyHeadingPanics(t *testing.T) {
	mustPanic(t, "Posts needs a Heading", func() {
		example.Render(site.Posts{Size: site.PostSize{}, Heading: "", Class: "", Children: g.Text("")})
	})
}

func TestPostTakesThePagesUtilities(t *testing.T) {
	got := example.Render(site.Post{Href: "/x/", Media: nil, Category: "", Datetime: "", Date: "", Headline: "x", Class: "hidden"})
	if !strings.HasPrefix(got, `<a class="cui-post hidden"`) {
		t.Errorf("got %s", got)
	}
}
