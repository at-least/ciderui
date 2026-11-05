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

func storiesOf(story site.Story) site.Stories {
	return site.Stories{Heading: "x", Class: "", Items: []site.Story{story}}
}

func TestStoryNoCategoryNoDateNoPicture(t *testing.T) {
	got := example.Render(storiesOf(site.Story{Size: site.StorySizeSmall, Href: "/x/", Media: nil, Category: "", Date: nil, Headline: "x", Class: ""}))
	want := `<li><a class="cui-story" href="/x/"><div class="cui-story-content"><div class="cui-story-head"><h3 class="cui-story-headline">x</h3></div></div></a></li>`
	if !strings.Contains(got, want) {
		t.Errorf("got %s", got)
	}
}

func TestStoryADateWithNoDatetimePanics(t *testing.T) {
	mustPanic(t, "site.Date needs a Datetime", func() {
		example.Render(storiesOf(site.Story{Size: site.StorySizeSmall, Href: "/x/", Media: nil, Category: "", Date: &site.Date{Datetime: "", Text: "Today"}, Headline: "x", Class: ""}))
	})
}

func TestStoriesAnEmptyHeadingPanics(t *testing.T) {
	mustPanic(t, "Stories needs a Heading", func() {
		example.Render(site.Stories{Heading: "", Class: "", Items: nil})
	})
}

func TestStoryTakesThePagesUtilities(t *testing.T) {
	got := example.Render(storiesOf(site.Story{Size: site.StorySizeSmall, Href: "/x/", Media: nil, Category: "", Date: nil, Headline: "x", Class: "hidden"}))
	if !strings.Contains(got, `<a class="cui-story hidden"`) {
		t.Errorf("got %s", got)
	}
}

func TestStoryAnEmptyHrefPanics(t *testing.T) {
	mustPanic(t, "Story needs an Href", func() {
		example.Render(storiesOf(site.Story{Size: site.StorySizeSmall, Href: "", Media: nil, Category: "", Date: nil, Headline: "x", Class: ""}))
	})
}

func TestStoryAnEmptyHeadlinePanics(t *testing.T) {
	mustPanic(t, "Story needs a Headline", func() {
		example.Render(storiesOf(site.Story{Size: site.StorySizeSmall, Href: "/x/", Media: nil, Category: "", Date: nil, Headline: "", Class: ""}))
	})
}
