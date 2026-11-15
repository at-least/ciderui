package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: apple.com's Newsroom, a press release (Apple TV 4K, 2026-10-13: a hero
// image, an inline image, subheads, small print with a list, footnotes, the
// related stories below): the article's column and rhythm, its header, its
// copy and what markdown writes in it (a paragraph, a link, a footnote's
// mark, a subhead, a list, a figure and its caption), its footnotes, its
// related rows, and the local nav it has (Newsroom's, ac-ln), at the steps
// it takes (735 and 1069) and the kit's widths. What is a press release's
// alone (the dateline, the share sheet, the downloads, the press contacts,
// About Apple) is not read.
func newsroomArticle() pagespec.Spec {
	COLUMN := []string{"width", "max-width", "margin-left", "margin-right", "padding-left", "padding-right"}
	SPACE := []string{"margin-top", "margin-bottom", "padding-top", "padding-bottom"}
	// @external: how a word too long for its line breaks, if it does
	WRAP := []string{"overflow-wrap", "word-break", "hyphens", "white-space"}

	return pagespec.Spec{
		URL:    "https://www.apple.com/newsroom/2026/10/apple-unveils-the-new-apple-tv-4k/",
		Widths: []int{1600, 1200, 1069, 1068, 900, 735, 734, 390},
		Probes: []pagespec.Probe{
			{Name: "page.body", Select: "body", Props: PAGE},
			// @external: the column: each component's content box, the copy's and a figure's
			{Name: "column.header", Select: ".article-header .component-content", Props: COLUMN, Box: true},
			{Name: "column.copy", Select: ".pagebody .component-content", Props: COLUMN, Box: true},
			// @external: the text itself: below 735px the copy's content box is the figure's
			// (414px at most) and each paragraph takes its own margins
			{Name: "column.text", Select: ".pagebody .component-content > .pagebody-copy", Props: COLUMN, Box: true},
			{Name: "column.figure", Select: ".pagebody", Props: COLUMN, Box: true},
			// @external: the local nav (Newsroom's own, not a product page's)
			{Name: "localnav", Select: "#ac-localnav", Props: []string{"height", "position", "top", "z-index"}, Box: true},
			{Name: "localnav.background", Select: "#ac-localnav .ac-ln-background", Props: slices.Concat([]string{"background-color"}, BLUR)},
			{Name: "localnav.content", Select: "#ac-localnav .ac-ln-content", Props: []string{"max-width", "padding-left", "padding-right"}, Box: true},
			{Name: "localnav.title", Select: "#ac-localnav .ac-ln-title a", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			// @external: the header: its parts' room, the eyebrow (category, date), the
			// headline, the subhead
			{Name: "header", Select: ".article-header", Props: SPACE, Box: true},
			{Name: "header.eyebrow", Select: ".article-header .category.component", Props: SPACE, Box: true},
			{Name: "header.headline", Select: ".article-header .pagetitle.component", Props: SPACE, Box: true},
			{Name: "header.subhead", Select: ".article-header .article-subhead.component", Props: SPACE, Box: true},
			{Name: "eyebrow", Select: ".category-eyebrow", Props: slices.Concat(BOX, FLEX)},
			{Name: "eyebrow.category", Select: ".category-eyebrow__category", Props: slices.Concat(TYPE, TEXT, []string{"text-transform"})},
			{Name: "eyebrow.date", Select: ".category-eyebrow__date", Props: slices.Concat(TYPE, TEXT, []string{"margin-top"})},
			{Name: "headline", Select: "h1.hero-headline", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom", "text-align"}, WRAP)},
			{Name: "subhead", Select: ".article-subhead .component-content", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}, WRAP)},
			// @external: the hero figure, its picture and its caption
			{Name: "figure", Select: "figure.image-fullbleed", Props: slices.Concat(COLUMN, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "figure.img", Select: "figure.image-fullbleed img.picture-image", Props: []string{"max-width", "border-top-left-radius"}, Box: true},
			{Name: "figure.caption", Select: "figure.image-fullbleed .image-caption", Props: slices.Concat(TYPE, SPACE, []string{"text-align"}, WRAP)},
			// @external: the caption's room: the text's column, under the picture
			{Name: "figure.description", Select: "figure.image-fullbleed .image-description", Props: slices.Concat(COLUMN, []string{"margin-top"}), Box: true},
			{Name: "figure.inline", Select: "figure.image-inline", Props: slices.Concat(COLUMN, []string{"margin-top", "margin-bottom"}), Box: true},
			// @external: the copy: a paragraph of it, a paragraph a page adds there, a link, a
			// footnote's mark, bold, a subhead and its strong
			{Name: "copy", Select: ".pagebody .component-content > .pagebody-copy", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}, WRAP)},
			{Name: "copy.p", Inject: "<p>Sample text</p>", Into: ".pagebody .component-content", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
			{Name: "copy.link", Select: ".pagebody-copy a", Props: slices.Concat(TYPE, LINK, WRAP), States: []string{"hover"}},
			{Name: "copy.sup", Select: ".pagebody-copy sup", Props: slices.Concat(TYPE, []string{"font-feature-settings", "vertical-align", "position", "top"})},
			{Name: "copy.strong", Select: ".pagebody-small .pagebody-copy li strong", Props: TYPE},
			{Name: "h2", Select: "h2.pagebody-header", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}, WRAP)},
			{Name: "h2.strong", Select: "h2.pagebody-header strong", Props: TYPE},
			// @external: the small print's list
			{Name: "list", Select: ".pagebody-small .pagebody-copy ul", Props: []string{"margin-top", "margin-bottom", "padding-left", "list-style-type", "list-style-position", "margin-left"}},
			{Name: "list.li", Select: ".pagebody-small .pagebody-copy li", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
			{Name: "list.link", Select: ".pagebody-small .pagebody-copy li a", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			// @external: the footnotes
			{Name: "footnotes", Select: ".sosumi", Props: slices.Concat(SPACE, []string{"border-top-width", "border-top-style", "border-top-color"})},
			{Name: "footnotes.ol", Select: ".sosumi ol", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom", "padding-left", "list-style-type", "margin-left"})},
			{Name: "footnotes.li", Select: ".sosumi li", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}, WRAP)},
			{Name: "footnotes.link", Select: ".sosumi li a", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			// @external: the related stories: a heading over ruled rows, each a link with a
			// picture, a category, a title and a date
			{Name: "related.heading", Select: ".article-list__heading", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom", "text-align"})},
			{Name: "related.list", Select: ".article-list", Props: slices.Concat(SPACE, COLUMN), Box: true},
			{Name: "related.item", Select: ".article-list__item", Props: slices.Concat(BOX, []string{"border-bottom-width"}), Box: true},
			{Name: "related.link", Select: ".article-list__item__link", Props: slices.Concat(FLEX, LINK, []string{"cursor"}), Box: true, States: []string{"hover"}},
			{Name: "related.img", Select: ".article-list__item__img-wrapper", Props: slices.Concat(BOX, []string{"width", "height", "flex-shrink"}), Box: true},
			{Name: "related.category", Select: ".article-list__item__category", Props: slices.Concat(TYPE, TEXT, []string{"text-transform", "margin-bottom"})},
			{Name: "related.title", Select: ".article-list__item__title", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}, LINK), States: []string{"hover"}},
			{Name: "related.date", Select: ".article-list__item__date", Props: slices.Concat(TYPE, []string{"margin-top"})},
			// @external: the room under the related stories, before the footer's band (the
			// footer tile, Newsroom's own, not read)
			{Name: "footertile", Select: ".footertile", Props: []string{"margin-top", "margin-bottom"}},
		},
	}
}
