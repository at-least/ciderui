package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	h "maragu.dev/gomponents/html"
)

func TestPostsRendersTheExamples(t *testing.T) {
	const picture = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMjgwIiBoZWlnaHQ9IjcyMCIvPg=="
	post := func(href, category, datetime, date, headline string) site.Post {
		media := h.Img(h.Class("cui-post-media"), h.Src(picture), h.Alt(""), h.Width("1280"), h.Height("720"))
		return site.Post{Href: href, Media: media, Category: category, Date: &site.Date{Datetime: datetime, Text: date}, Headline: headline, Class: ""}
	}
	example.Check(t, site.Posts{Small: false, Heading: "October 2026", Class: "", Items: []site.Post{
		post("/news/zola/", "Update", "2026-10-19", "October 19, 2026", "ciderui comes to Zola"),
		post("/news/article/", "Update", "2026-10-18", "October 18, 2026", "An article in the Newsroom's look"),
		post("/news/stories/", "Feature", "2026-10-17", "October 17, 2026", "A grid of stories"),
	}}, "site/posts/examples/archive.html")
	example.Check(t, site.Posts{Small: true, Heading: "More from the Newsroom", Class: "", Items: []site.Post{
		post("/news/archive/", "Update", "2026-10-16", "October 16, 2026", "The archive's rows"),
		post("/news/fonts/", "Update", "2026-10-15", "October 15, 2026", "Inter where SF is not"),
	}}, "site/posts/examples/related.html")
}

func postsOf(post site.Post) site.Posts {
	return site.Posts{Small: false, Heading: "x", Class: "", Items: []site.Post{post}}
}

func TestPostADateWithNoDatetimePanics(t *testing.T) {
	mustPanic(t, "site.Date needs a Datetime", func() {
		example.Render(postsOf(site.Post{Href: "/x/", Media: nil, Category: "", Date: &site.Date{Datetime: "", Text: "Today"}, Headline: "x", Class: ""}))
	})
}

func TestPostsAnEmptyHeadingPanics(t *testing.T) {
	mustPanic(t, "Posts needs a Heading", func() {
		example.Render(site.Posts{Small: false, Heading: "", Class: "", Items: nil})
	})
}

func TestPostTakesThePagesUtilities(t *testing.T) {
	got := example.Render(postsOf(site.Post{Href: "/x/", Media: nil, Category: "", Date: nil, Headline: "x", Class: "hidden"}))
	if !strings.Contains(got, `<li><a class="cui-post hidden"`) {
		t.Errorf("got %s", got)
	}
}

func TestPostAnEmptyHrefPanics(t *testing.T) {
	mustPanic(t, "Post needs an Href", func() {
		example.Render(postsOf(site.Post{Href: "", Media: nil, Category: "", Date: nil, Headline: "x", Class: ""}))
	})
}

func TestPostAnEmptyHeadlinePanics(t *testing.T) {
	mustPanic(t, "Post needs a Headline", func() {
		example.Render(postsOf(site.Post{Href: "/x/", Media: nil, Category: "", Date: nil, Headline: "", Class: ""}))
	})
}
