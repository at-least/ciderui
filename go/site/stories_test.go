package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestStoriesRendersTheExample(t *testing.T) {
	const picture = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMjgwIiBoZWlnaHQ9IjcyMCIvPg=="
	type row struct {
		size                                     site.StorySize
		href, category, datetime, date, headline string
	}
	stories := []row{
		{site.StorySizeLarge, "/news/zola/", "Update", "2026-10-19", "October 19, 2026", "ciderui comes to Zola"},
		{site.StorySizeMedium, "/news/article/", "Update", "2026-10-18", "October 18, 2026", "An article in the Newsroom's look"},
		{site.StorySizeMedium, "/news/stories/", "Feature", "2026-10-17", "October 17, 2026", "A grid of stories"},
		{site.StorySize{}, "/news/archive/", "Update", "2026-10-16", "October 16, 2026", "The archive's rows"},
		{site.StorySize{}, "/news/fonts/", "Update", "2026-10-15", "October 15, 2026", "Inter where SF is not"},
		{site.StorySize{}, "/news/tokens/", "Update", "2026-10-14", "October 14, 2026", "A fixed palette"},
	}
	example.Check(t, site.Stories{Heading: "Latest News", Class: "", Children: g.Map(stories, func(r row) g.Node {
		media := h.Img(h.Class("cui-story-media"), h.Src(picture), h.Alt(""), h.Width("1280"), h.Height("720"))
		return h.Li(site.Story{Size: r.size, Href: r.href, Media: media, Category: r.category, Datetime: r.datetime, Date: r.date, Headline: r.headline, Class: ""})
	})}, "site/stories/examples/stories.html")
}

func TestStoryNoCategoryNoDateNoPicture(t *testing.T) {
	got := example.Render(site.Story{Size: site.StorySize{}, Href: "/x/", Media: nil, Category: "", Datetime: "", Date: "", Headline: "x", Class: ""})
	want := `<a class="cui-story" href="/x/"><div class="cui-story-content"><div class="cui-story-head"><h3 class="cui-story-headline">x</h3></div></div></a>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestStoryADateWithNoDatetimePanics(t *testing.T) {
	mustPanic(t, "Date needs its Datetime", func() {
		example.Render(site.Story{Size: site.StorySize{}, Href: "/x/", Media: nil, Category: "", Datetime: "", Date: "Today", Headline: "x", Class: ""})
	})
}

func TestStoriesAnEmptyHeadingPanics(t *testing.T) {
	mustPanic(t, "Stories needs a Heading", func() {
		example.Render(site.Stories{Heading: "", Class: "", Children: g.Text("")})
	})
}

func TestStoryTakesThePagesUtilities(t *testing.T) {
	got := example.Render(site.Story{Size: site.StorySize{}, Href: "/x/", Media: nil, Category: "", Datetime: "", Date: "", Headline: "x", Class: "hidden"})
	if !strings.HasPrefix(got, `<a class="cui-story hidden"`) {
		t.Errorf("got %s", got)
	}
}

func TestStoryAnEmptyHrefPanics(t *testing.T) {
	mustPanic(t, "Story needs an Href", func() {
		example.Render(site.Story{Size: site.StorySize{}, Href: "", Media: nil, Category: "", Datetime: "", Date: "", Headline: "x", Class: ""})
	})
}

func TestStoryAnEmptyHeadlinePanics(t *testing.T) {
	mustPanic(t, "Story needs a Headline", func() {
		example.Render(site.Story{Size: site.StorySize{}, Href: "/x/", Media: nil, Category: "", Datetime: "", Date: "", Headline: "", Class: ""})
	})
}
