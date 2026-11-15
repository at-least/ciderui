package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

func releases() pagespec.Spec {
	ENTRY := ".article-list-wrapper article.article"
	return pagespec.Spec{
		URL:  "https://developer.apple.com/news/releases/",
		Wait: 3000,
		Probes: []pagespec.Probe{
			{Name: "page", Select: "body", Props: PAGE},
			{Name: "entry", Select: ENTRY, Props: []string{"padding-top", "padding-bottom", "margin-top", "margin-bottom"}, Box: true},
			{Name: "entry.picture", Select: ENTRY + " .article-image-container", Props: []string{"width", "padding-top", "padding-right", "padding-bottom", "padding-left"}, Box: true},
			{Name: "entry.image", Select: ENTRY + " img.article-image", Props: []string{"width", "height", "border-top-left-radius"}, Box: true},
			{Name: "entry.content", Select: ENTRY + " .article-content-container", Props: []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "border-top-width", "border-top-style", "border-top-color"}, Box: true},
			{Name: "entry.title.link", Select: ENTRY + " a.article-title", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			{Name: "entry.title", Select: ENTRY + " a.article-title h2", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "entry.date", Select: ENTRY + " p.article-date", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
			{Name: "entry.text", Select: ENTRY + " .article-text", Props: slices.Concat(TYPE, []string{"margin-top"})},
			{Name: "entry.link", Select: ENTRY + " .article-text a", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			{Name: "entry.share", Select: ENTRY + " button.social-icon", Props: slices.Concat(TYPE, []string{"width", "height", "margin-right"}), Box: true},
			{Name: "next", Select: ENTRY + " ~ article.article .article-content-container", Props: []string{"padding-top", "margin-top", "border-top-width", "border-top-style", "border-top-color"}},
			{Name: "rss", Select: "a.rss-link", Props: slices.Concat(TYPE, []string{"width", "height"}), Box: true},
		},
	}
}
