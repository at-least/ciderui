package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: A product page, iPhone 18 Pro's, a dark page: apple.com's base type scale
// (the classes its stylesheets define, added to the page), the global nav's dark side,
// the content's width, a section, its header, a dark and an alternate section, a
// rounded tile, a gallery (its scroller, items, paddles and dots), the "more" link, a
// stat, a footnote, and the floating local nav, which shows once the page is scrolled
// past its hero.
func product() pagespec.Spec {
	TYPOGRAPHY := []string{
		"headline-super", "headline-elevated", "headline", "headline-reduced", "eyebrow-super",
		"intro-elevated", "intro", "body", "body-reduced", "body-reduced-tight", "callout-super", "callout-elevated",
		"callout-base", "callout-cta", "caption-elevated", "caption", "caption-cta", "caption-sosumi",
		"caption-elevated-sosumi", "label", "section-header-eyebrow", "section-header-headline", "section-header-link",
	}
	var typography []pagespec.Probe
	for _, c := range TYPOGRAPHY {
		typography = append(typography, pagespec.Probe{Name: "type." + c, Inject: `<p class="typography-` + c + `">Sample text</p>`, Into: "main", Props: TYPE})
	}
	return pagespec.Spec{
		URL: "https://www.apple.com/iphone-18-pro/",
		Probes: slices.Concat([]pagespec.Probe{
			// @external: a class the page does not define reads as this paragraph does
			{Name: "type.none", Inject: "<p>Sample text</p>", Into: "main", Props: TYPE},
		}, typography, []pagespec.Probe{
			{Name: "page.body", Select: "body", Props: PAGE},
			{Name: "page.strong", Select: "main strong, main b, .rs-page-content strong, body strong", Props: TYPE},
			// @external: the page is dark (body.globalheader-dark): the global nav's dark side
			{Name: "nav", Select: "#globalnav", Props: slices.Concat([]string{"position", "background-color"}, BLUR)},
			{Name: "nav.link", Select: "#globalnav .globalnav-item:not(.globalnav-item-apple) .globalnav-link", Props: TYPE, States: []string{"hover"}},
			{Name: "content", Select: "main .viewport-content", Props: []string{"width", "max-width", "margin-left", "margin-right", "padding-left", "padding-right"}, Box: true},
			{Name: "section", Select: "main section.section:not(.theme-dark):not(.background-alt)", Props: slices.Concat(BOX, []string{"padding-top", "padding-bottom"}), Box: true},
			{Name: "section.dark", Select: "main section.section.theme-dark", Props: slices.Concat(TYPE, []string{"background-color"})},
			{Name: "section.alt", Select: "main .background-alt", Props: slices.Concat(TYPE, []string{"background-color"})},
			{Name: "section.header", Select: "main .section-header", Props: slices.Concat(BOX, FLEX), Box: true},
			{Name: "section.header-links", Select: "main .section-header-cta-list", Props: slices.Concat(TYPE, BOX, FLEX, []string{"list-style-type"})},
			// @external: a section's padding: the highlights' (dark, alternate), the battery's (dark)
			{Name: "section.highlights", Select: "main section.section-highlights", Props: BOX},
			{Name: "section.battery", Select: "main section.section-battery", Props: BOX},
			{Name: "section.headline", Select: "main .section-header-headline", Props: slices.Concat(TYPE, []string{"margin-bottom"})},
			{Name: "section.header-link", Select: "main .section-header-cta-list a", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			{Name: "tile.rounded", Select: "main .tile.tile-rounded", Props: slices.Concat(BOX, []string{"overflow"}), Box: true},
			{Name: "tile.rounded-light", Select: "main .theme-light .tile.tile-rounded, main .tile.tile-rounded:not(.theme-dark *)", Props: BOX, Box: true},
			{Name: "paddle", Select: "main .paddlenav-arrow", Props: slices.Concat(BOX, []string{"width", "height", "color", "opacity"}, BLUR, MOTION), Box: true, States: []string{"hover", "active"}},
			{Name: "paddle.icon", Select: "main .paddlenav-arrow svg", Props: []string{"width", "height", "fill", "color"}, Box: true},
			// @external: the footer on the dark page: light as everywhere (the footer tokens'
			// dark side)
			{Name: "footer", Select: "#ac-globalfooter", Props: slices.Concat(TYPE, []string{"background-color"})},
			{Name: "footer.title", Select: ".ac-gf-directory-column-section-title-text", Props: TYPE},
			{Name: "footer.link", Select: ".ac-gf-directory-column-section-link", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			{Name: "footer.shop", Select: ".ac-gf-footer-shop", Props: []string{"border-bottom-color"}},
			// @external: a rule between copy blocks on the dark page: the divider's dark side
			{Name: "divider", Select: "main .callout-subsection-copy-block.keyline", Props: []string{"border-top-width", "border-top-style", "border-top-color"}},
			// @external: a gallery (the first scroll gallery, the camera system's): its scroller,
			// its items' inset and gap, an item's snapping, its paddles' row below
			// it and an enabled paddle (the next: the first, the previous, is
			// disabled); a media card gallery's dots, 8px, the current a pill (its
			// timer's fill drawn in it)
			{Name: "gallery.scroller", Select: "main .scroll-gallery .scroll-container", Props: []string{"overflow-x", "overflow-y", "scroll-snap-type", "scroll-padding-left", "overscroll-behavior-x", "padding-top", "padding-bottom", "margin-top", "margin-bottom"}, Box: true},
			{Name: "gallery.items", Select: "main .scroll-gallery .item-container", Props: []string{"padding-left", "padding-right"}, Box: true},
			{Name: "gallery.set", Select: "main .scroll-gallery .item-container > ul", Props: []string{"display", "column-gap"}, Box: true},
			{Name: "gallery.item", Select: "main .scroll-gallery .item-container > ul > li", Props: []string{"scroll-snap-align", "scroll-snap-stop"}, Box: true, Position: true},
			{Name: "gallery.nav", Select: "main .scroll-gallery .scroll-gallery-paddlenav", Props: []string{"display", "justify-content", "padding-top"}, Box: true, Position: true},
			{Name: "gallery.nav-list", Select: "main .scroll-gallery .scroll-gallery-paddlenav ul", Props: []string{"display", "column-gap", "margin-top"}, Box: true},
			{Name: "paddle.next", Select: "main .scroll-gallery .paddlenav-arrow-next", Props: slices.Concat(BOX, []string{"width", "height", "color", "opacity", "cursor"}, MOTION), Box: true, Position: true, States: []string{"hover", "active"}},
			{Name: "gallery.dots", Select: "main .dotnav-items", Props: FLEX, Box: true},
			{Name: "gallery.dot-item", Select: "main .dotnav-item:not(.current)", Props: []string{"margin-left", "margin-right", "padding-left", "padding-right"}, Box: true},
			{Name: "gallery.dot", Select: "main .dotnav-link:not(.current)", Props: slices.Concat([]string{"width", "height", "background-color", "border-top-left-radius", "padding-top"}, MOTION), Box: true, States: []string{"hover"}},
			{Name: "gallery.dot-current", Select: "main .dotnav-link.current", Props: []string{"width", "height", "background-color", "border-top-left-radius", "padding-top"}, Box: true},
			{Name: "gallery.dot-fill", Select: "main .dotnav-link.current", Pseudo: "::after", Props: []string{"width", "height", "background-color", "border-top-left-radius", "position", "content"}},
			// @external: a link in running text in a dark section: apple.com's own `a`, on markup added to
			// it
			{Name: "link.inline", Inject: "<p>Text with <a href=\"#cui\">a link</a> in it.</p>", Into: "main section.section.theme-dark", Select: "a", Props: slices.Concat(TYPE, LINK), States: []string{"hover", "active"}},
			{Name: "link.more", Select: "main a.icon-wrapper", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			{Name: "link.more-icon", Select: "main a.icon-wrapper .icon-after", Props: []string{"font-family", "font-size", "color", "padding-left", "margin-left", "top", "position", "content"}, Pseudo: "::after"},
			{Name: "stat.content", Select: "main .stat .stat-content", Props: []string{"padding-top", "padding-bottom", "margin-top", "margin-bottom"}},
			{Name: "footnote.link", Select: "main sup.footnote a", Props: []string{"font-feature-settings", "color", "text-decoration-line", "font-size", "vertical-align"}},
			{Name: "stat.value", Select: "main .stat-value", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom", "display"})},
			{Name: "stat.caption", Select: "main .stat-caption", Props: slices.Concat(TYPE, []string{"display"})},
			{Name: "footnote", Select: "main sup.footnote", Props: slices.Concat([]string{"font-feature-settings"}, TYPE, []string{"vertical-align", "top", "position"})},
			{Name: "localnav", Select: "#localnav-container", Scroll: 2500, Props: slices.Concat(BOX, []string{"position", "top", "bottom", "left", "width", "height", "z-index"}, BLUR), Box: true},
			{Name: "localnav.wrap", Select: "#localnav-container .localnav-wrap", Scroll: 2500, Props: slices.Concat(BOX, []string{"height"}, BLUR), Box: true},
			{Name: "localnav.product", Select: "#localnav-container .product-name", Scroll: 2500, Props: TYPE},
			{Name: "localnav.explore", Select: "#localnav-container .cta.explore", Scroll: 2500, Props: slices.Concat(TYPE, BOX), Box: true, States: []string{"hover"}},
			{Name: "localnav.buy", Select: "#localnav-container .cta.buy", Scroll: 2500, Props: slices.Concat(TYPE, BOX), Box: true, States: []string{"hover"}},
			// @invariant: the global nav's icons on a dark page, and its phone menu opened: last
			{Name: "nav.search", Select: "#globalnav .globalnav-link-search", Props: []string{"color"}, States: []string{"hover"}},
			{Name: "nav.menutrigger", Select: "#globalnav .globalnav-menutrigger-button", Props: []string{"color", "opacity"}, Widths: []int{900, 390}, States: []string{"hover"}},
			{Name: "menu.open", Click: "#globalnav-menutrigger-button", Select: "#globalnav", Props: slices.Concat([]string{"background-color"}, BLUR), Widths: []int{390}},
			{Name: "menu.content", Select: "#globalnav .globalnav-content", Props: slices.Concat([]string{"background-color"}, BLUR), Widths: []int{390}},
			{Name: "menu.link", Select: "#globalnav .globalnav-item-mac .globalnav-link", Props: TYPE, Widths: []int{390}, States: []string{"hover"}},
			{Name: "menu.trigger", Select: "#globalnav .globalnav-menutrigger-button", Props: []string{"color", "opacity"}, Widths: []int{390}},
		}),
	}
}
