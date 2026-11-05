package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	h "maragu.dev/gomponents/html"
)

func TestStoriesRendersTheExample(t *testing.T) {
	const picture = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMjgwIiBoZWlnaHQ9IjcyMCIvPg=="
	story := func(size site.StorySize, href, category, datetime, date, headline string) site.Story {
		media := h.Img(h.Class("cui-story-media"), h.Src(picture), h.Alt(""), h.Width("1280"), h.Height("720"))
		return site.Story{Size: size, Href: href, Media: media, Category: category, Date: &site.Date{Datetime: datetime, Text: date}, Headline: headline, Class: ""}
	}
	example.Check(t, site.Stories{Heading: "Latest News", Class: "", Items: []site.Story{
		story(site.StorySizeLarge, "/news/zola/", "Update", "2026-10-19", "October 19, 2026", "ciderui comes to Zola"),
		story(site.StorySizeMedium, "/news/article/", "Update", "2026-10-18", "October 18, 2026", "An article in the Newsroom's look"),
		story(site.StorySizeMedium, "/news/stories/", "Feature", "2026-10-17", "October 17, 2026", "A grid of stories"),
		story(site.StorySizeSmall, "/news/archive/", "Update", "2026-10-16", "October 16, 2026", "The archive's rows"),
		story(site.StorySizeSmall, "/news/fonts/", "Update", "2026-10-15", "October 15, 2026", "Inter where SF is not"),
		story(site.StorySizeSmall, "/news/tokens/", "Update", "2026-10-14", "October 14, 2026", "A fixed palette"),
	}}, "site/stories/examples/stories.html")
}

func validStories() site.Stories {
	return site.Stories{Heading: "x", Class: "", Items: []site.Story{{Size: site.StorySizeSmall, Href: "/x/", Media: nil, Category: "", Date: nil, Headline: "y", Class: ""}}}
}

func TestStoriesPanics(t *testing.T) {
	example.Panics(t, validStories, []example.Break[site.Stories]{
		{Want: "site.Stories needs a Heading", Change: func(p *site.Stories) { p.Heading = "" }},
		{Want: "site.Story needs an Href", Change: func(p *site.Stories) { p.Items[0].Href = "" }},
		{Want: "site.Story needs a Headline", Change: func(p *site.Stories) { p.Items[0].Headline = "" }},
		{Want: "site.Date needs a Datetime", Change: func(p *site.Stories) { p.Items[0].Date = &site.Date{Datetime: "", Text: "Today"} }},
	})
}

func TestStoryNoCategoryNoDateNoPicture(t *testing.T) {
	want := `<li><a class="cui-story" href="/x/"><div class="cui-story-content"><div class="cui-story-head"><h3 class="cui-story-headline">y</h3></div></div></a></li>`
	if got := example.Render(validStories()); !strings.Contains(got, want) {
		t.Errorf("no %s in %s", want, got)
	}
}
