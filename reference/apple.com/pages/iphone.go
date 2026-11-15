package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The iPhone page, a light one: its links (the "more" link, its copy and its
// chevron; a section header's link). No ribbon is read here: 800ms in, it
// was mid-motion (2026-10-22); ribbon.go reads one once its motion has run.
func iphone() pagespec.Spec {
	return pagespec.Spec{
		URL: "https://www.apple.com/iphone/",
		Probes: []pagespec.Probe{
			{Name: "page.body", Select: "body", Props: PAGE},
			{Name: "link.more", Select: "main a.icon-wrapper:not(.section-header-cta-link)", Props: slices.Concat(TYPE, LINK, TEXT), States: []string{"hover", "focus-visible"}},
			{Name: "link.more-copy", Select: "main a.icon-wrapper:not(.section-header-cta-link) .icon-copy", StateOn: "main a.icon-wrapper:not(.section-header-cta-link)", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			{Name: "link.more-icon", Select: "main a.icon-wrapper:not(.section-header-cta-link) .icon-after", Props: []string{"display", "font-size", "line-height", "padding-left", "margin-left", "vertical-align", "position", "top"}, Box: true},
			{Name: "link.more-glyph", Select: "main a.icon-wrapper:not(.section-header-cta-link) .icon-after", Pseudo: "::after", Props: []string{"content", "font-family", "font-size", "line-height", "font-weight", "display", "padding-left", "margin-left", "top", "position", "vertical-align", "color", "text-decoration-line"}},
			// @external: a link in running text: apple.com's own `a`, on markup added to the page
			{Name: "link.inline", Inject: "<p>Text with <a href=\"#cui\">a link</a> in it.</p>", Into: "main", Select: "a", Props: slices.Concat(TYPE, LINK), States: []string{"hover", "focus-visible", "active"}},
			{Name: "link.section-header", Select: "main a.section-header-cta-link", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			// @external: the content's width; a section's padding (light, alternate), its header
			{Name: "content", Select: "main .viewport-content", Props: []string{"width", "max-width", "margin-left", "margin-right", "padding-left", "padding-right"}, Box: true},
			{Name: "section.incentive", Select: "main section.section-incentive", Props: BOX},
			{Name: "section.image-accordion", Select: "main section.section-image-accordion", Props: BOX},
			{Name: "section.header", Select: "main .section-header", Props: slices.Concat(BOX, FLEX), Box: true},
			{Name: "section.headline", Select: "main .section-header-headline", Props: slices.Concat(TYPE, []string{"margin-bottom"})},
			// @external: a light page's gallery paddles (the first enabled one, and a disabled
			// one) and a media card gallery's dots
			{Name: "paddle", Select: "main .paddlenav-arrow-next:not([disabled])", Props: slices.Concat(BOX, []string{"width", "height", "color", "opacity", "cursor"}, MOTION), Box: true, States: []string{"hover", "active", "focus-visible"}},
			{Name: "paddle.icon", Select: "main .paddlenav-arrow-next:not([disabled]) svg", Props: []string{"width", "height", "fill", "color"}, Box: true},
			{Name: "paddle.disabled", Select: "main .paddlenav-arrow[disabled]", Props: []string{"opacity", "background-color", "color", "cursor", "pointer-events"}},
			{Name: "gallery.dot", Select: "main .dotnav-link:not(.current)", Props: slices.Concat([]string{"width", "height", "background-color", "border-top-left-radius", "padding-top"}, MOTION), Box: true, States: []string{"hover"}},
			{Name: "gallery.dot-current", Select: "main .dotnav-link.current", Props: []string{"width", "height", "background-color", "border-top-left-radius", "padding-top"}, Box: true},
			{Name: "gallery.dot-fill", Select: "main .dotnav-link.current", Pseudo: "::after", Props: []string{"width", "height", "background-color", "border-top-left-radius", "position", "content"}},
		},
	}
}
