package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: apple.com's home page: the type scale (every typography class its stylesheets
// define, added to the page), the page, the global nav, the hero and its
// buttons, a light and a dark promo tile, the gallery's dots, the footer.
func home() pagespec.Spec {
	TYPOGRAPHY := []string{"headline-super", "headline-elevated", "headline", "headline-reduced", "intro-elevated", "intro",
		"body", "body-reduced", "callout", "caption", "cta-medium", "tile-headline", "tile-subhead", "tile-callout",
		"promo-headline", "promo-subhead", "promo-callout"}
	var typography []pagespec.Probe
	for _, c := range TYPOGRAPHY {
		typography = append(typography, pagespec.Probe{Name: "type." + c, Inject: `<p class="typography-` + c + `">Sample text</p>`, Into: "main", Props: TYPE})
	}
	STATES := []string{"hover", "active", "focus-visible"}
	return pagespec.Spec{
		URL: "https://www.apple.com/",
		Probes: slices.Concat([]pagespec.Probe{
			// @external: a class the page does not define reads as this paragraph does
			{Name: "type.none", Inject: "<p>Sample text</p>", Into: "main", Props: TYPE},
		}, typography, []pagespec.Probe{
			{Name: "page.html", Select: "html", Props: []string{"background-color", "color-scheme"}},
			{Name: "page.body", Select: "body", Props: PAGE},
			{Name: "nav", Select: "#globalnav", Props: slices.Concat([]string{"height", "position", "background-color"}, BLUR), Vars: []string{"--globalnav-background", "--globalnav-backdrop-filter"}, Box: true},
			{Name: "nav.curtain", Select: "#globalnav-curtain", Props: slices.Concat([]string{"background-color"}, BLUR)},
			{Name: "nav.content", Select: "#globalnav .globalnav-content", Props: slices.Concat([]string{"height", "max-width", "padding-left", "padding-right", "background-color"}, BLUR), Box: true},
			{Name: "nav.link", Select: "#globalnav .globalnav-item:not(.globalnav-item-apple) .globalnav-link", Props: slices.Concat(TYPE, []string{"opacity"}, MOTION), States: []string{"hover"}},
			{Name: "nav.link-text", Select: "#globalnav .globalnav-link-text", Props: TYPE},
			// @external: the bar's parts: its list, a link's box, the logo's, the search's and the bag's
			// and their images
			{Name: "nav.list", Select: "#globalnav .globalnav-list", Props: slices.Concat(BOX, FLEX, []string{"margin-left", "margin-right", "padding-left", "padding-right", "height"}), Box: true},
			{Name: "nav.link-box", Select: "#globalnav .globalnav-item-mac .globalnav-link", Props: slices.Concat(BOX, FLEX, []string{"padding-left", "padding-right", "height", "text-decoration-line"}), Box: true},
			{Name: "nav.logo", Select: "#globalnav .globalnav-link-apple", Props: slices.Concat(BOX, FLEX, []string{"padding-left", "padding-right", "height", "color"}, MOTION), Box: true, Position: true, States: []string{"hover"}},
			{Name: "nav.logo-image", Select: "#globalnav .globalnav-link-apple .globalnav-link-image:not([style*=\"none\"]) svg", Props: []string{"width", "height", "fill"}, Box: true},
			{Name: "nav.search", Select: "#globalnav .globalnav-link-search", Props: slices.Concat(BOX, FLEX, []string{"padding-left", "padding-right", "height", "color"}, MOTION), Box: true, Position: true, States: []string{"hover"}},
			{Name: "nav.search-image", Select: "#globalnav .globalnav-link-search svg", Props: []string{"width", "height", "fill"}, Box: true},
			{Name: "nav.bag", Select: "#globalnav .globalnav-link-bag", Props: slices.Concat(BOX, FLEX, []string{"padding-left", "padding-right", "height", "color"}), Box: true, Position: true, States: []string{"hover"}},
			{Name: "nav.bag-image", Select: "#globalnav .globalnav-link-bag svg", Props: []string{"width", "height", "fill"}, Box: true},
			// @external: the phone's images: the compact logo, search and bag (17 by 48), and the menu's
			// button
			{Name: "nav.logo-compact", Select: "#globalnav .globalnav-link-apple .globalnav-image-compact svg", Props: []string{"width", "height"}, Box: true, Widths: []int{390}},
			{Name: "nav.search-compact", Select: "#globalnav .globalnav-link-search .globalnav-image-compact svg", Props: []string{"width", "height"}, Box: true, Widths: []int{390}},
			{Name: "nav.bag-compact", Select: "#globalnav .globalnav-link-bag .globalnav-image-compact svg", Props: []string{"width", "height"}, Box: true, Widths: []int{390}},
			{Name: "nav.menutrigger", Select: "#globalnav .globalnav-menutrigger-button", Props: slices.Concat(BOX, FLEX, []string{"width", "height", "color", "opacity", "cursor"}, MOTION), Box: true, Position: true, Widths: []int{900, 390}, States: []string{"hover"}},
			// @external: its two lines are an SVG's polylines (reference/icons/ reads their points); their
			// box
			{Name: "nav.menutrigger-bread", Select: "#globalnav .globalnav-menutrigger-bread-top", Props: []string{"stroke-width"}, Box: true, Widths: []int{390}},
			{Name: "hero.section", Select: ".section-hero", Props: BOX, Box: true},
			{Name: "hero.tile", Select: ".section-hero .tile-wrapper", Props: BOX, Box: true},
			{Name: "hero.content", Select: ".section-hero .tile-content", Props: slices.Concat(BOX, []string{"text-align"}, FLEX), Box: true},
			{Name: "hero.headline", Select: ".section-hero .tile-headline", Props: slices.Concat(TYPE, []string{"margin-bottom"})},
			{Name: "hero.subhead", Select: ".section-hero .tile-subhead", Props: slices.Concat(TYPE, []string{"margin-top"})},
			{Name: "hero.callout", Select: ".section-hero .tile-callout", Props: slices.Concat(TYPE, []string{"margin-top"})},
			{Name: "hero.ctas", Select: ".section-hero .tile-ctas", Props: []string{"display", "gap", "column-gap", "justify-content", "margin-top"}},
			{Name: "button", Select: ".section-hero .tile-ctas .button:not(.button-secondary)", Props: slices.Concat(TYPE, BOX, []string{"display", "cursor"}, FOCUS, MOTION), Box: true, States: STATES},
			{Name: "button.secondary", Select: ".section-hero .tile-ctas .button-secondary", Props: slices.Concat(TYPE, BOX, []string{"display"}, FOCUS, MOTION), Box: true, States: STATES},
			{Name: "promo.section", Select: ".section-promo", Props: slices.Concat(BOX, []string{"display", "gap", "grid-template-columns"}), Box: true},
			{Name: "tile.light", Select: ".section-promo .tile-wrapper:not(.theme-dark)", Props: slices.Concat(BOX, FLEX, []string{"overflow", "text-align", "position"}), Box: true},
			{Name: "tile.light-headline", Select: ".section-promo .tile-wrapper:not(.theme-dark) .tile-headline", Props: TYPE},
			{Name: "tile.light-subhead", Select: ".section-promo .tile-wrapper:not(.theme-dark) .tile-subhead", Props: slices.Concat(TYPE, []string{"margin-top"})},
			// @external: a promo tile's spacing: its first with a text headline and buttons
			{Name: "tile.promo-content", Select: ".section-promo .tile-wrapper:has(h3.tile-headline:not(.logo-image)) .tile-content", Props: FLEX, Box: true},
			{Name: "tile.promo-headline", Select: ".section-promo .tile-wrapper:has(h3.tile-headline:not(.logo-image)) .tile-headline", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
			{Name: "tile.promo-subhead", Select: ".section-promo .tile-wrapper:has(h3.tile-headline:not(.logo-image)) .tile-subhead", Props: slices.Concat(TYPE, []string{"margin-top"})},
			{Name: "tile.promo-ctas", Select: ".section-promo .tile-wrapper .tile-ctas", Props: []string{"display", "gap", "column-gap", "justify-content", "margin-top", "grid-auto-flow"}},
			{Name: "tile.promo-button", Select: ".section-promo .tile-wrapper .tile-ctas .button", Props: slices.Concat(TYPE, []string{"padding-top", "padding-left", "min-width"}), Box: true},
			{Name: "tile.media", Select: ".section-promo .tile-wrapper .tile-image-wrapper", Props: []string{"position", "top", "left", "width", "height", "z-index"}, Box: true},
			{Name: "tile.media-image", Select: ".section-promo .tile-wrapper .tile-image-wrapper img", Props: []string{"position", "object-fit", "height", "z-index"}, Box: true},
			{Name: "tile.link", Select: ".section-promo .tile-wrapper .tile-link", Props: []string{"position", "z-index"}, Box: true},
			{Name: "tile.copy", Select: ".section-promo .tile-wrapper .tile-copy-wrapper", Props: []string{"position", "z-index"}},
			{Name: "tile.dark", Select: ".tile-wrapper.theme-dark", Props: BOX, Box: true},
			{Name: "tile.dark-headline", Select: ".tile-wrapper.theme-dark .tile-headline", Props: TYPE},
			{Name: "tile.dark-subhead", Select: ".tile-wrapper.theme-dark .tile-subhead", Props: TYPE},
			{Name: "tile.dark-button", Select: ".tile-wrapper.theme-dark .button:not(.button-secondary)", Props: slices.Concat(TYPE, BOX), States: []string{"hover", "active"}},
			{Name: "tile.dark-button-secondary", Select: ".tile-wrapper.theme-dark .button-secondary", Props: slices.Concat(TYPE, BOX), States: []string{"hover", "active"}},
			{Name: "gallery.dot", Select: ".dotnav-item:not(.current) .dotnav-link", Props: slices.Concat(BOX, []string{"width", "height", "opacity"}), Box: true, States: []string{"hover"}},
			{Name: "gallery.dot-current", Select: ".dotnav-item.current .dotnav-link", Props: slices.Concat(BOX, []string{"width", "height", "opacity"}), Box: true},
			// @external: a neutral button on a dark gallery item (li.media-gallery-item.theme-dark)
			{Name: "gallery.button-neutral", Select: ".button.button-neutral", Props: slices.Concat(TYPE, BOX), Box: true, States: []string{"hover", "active"}},
			{Name: "footer", Select: "#ac-globalfooter", Props: slices.Concat(TYPE, []string{"background-color"})},
			{Name: "footer.content", Select: "#ac-globalfooter .ac-gf-content", Props: []string{"max-width", "padding-left", "padding-right", "margin-left", "margin-right"}, Box: true},
			{Name: "footer.sosumi", Select: "#ac-globalfooter .ac-gf-sosumi", Props: slices.Concat(TYPE, BOX)},
			// @external: a link in running text: a footnote's
			{Name: "footer.sosumi-link", Select: "#ac-globalfooter .ac-gf-sosumi a", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			{Name: "footer.title", Select: ".ac-gf-directory-column-section-title-text", Props: TYPE},
			{Name: "footer.link", Select: ".ac-gf-directory-column-section-link", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			{Name: "footer.shop", Select: ".ac-gf-footer-shop", Props: slices.Concat(TYPE, BOX)},
			// @external: a link in the shop line's sentence
			{Name: "footer.shop-link", Select: ".ac-gf-footer-shop a", Props: slices.Concat(TYPE, LINK)},
			{Name: "footer.legal", Select: ".ac-gf-footer-legal", Props: slices.Concat(TYPE, BOX)},
			{Name: "footer.legal-link", Select: ".ac-gf-footer-legal-link", Props: slices.Concat(TYPE, LINK, []string{"border-right-color", "border-right-width", "padding-right", "margin-right"}), States: []string{"hover"}},
			{Name: "footer.copyright", Select: ".ac-gf-footer-legal-copyright", Props: slices.Concat(TYPE, []string{"margin-top", "margin-right"})},
			// @external: the footer's layout: its directory (a row of columns, a column's
			// sections, a title, a list, an item), the rows below it (the shop line,
			// the legal row: its links, their separators, the locale); on a phone its
			// sections' title rows and rules (opened in pages/home-footer.go)
			{Name: "footer.directory", Select: "#ac-globalfooter .ac-gf-directory", Props: slices.Concat(FLEX, []string{"padding-top"}), Box: true},
			{Name: "footer.column", Select: "#ac-globalfooter .ac-gf-directory-column", Props: []string{"flex-grow", "flex-shrink", "flex-basis"}, Box: true},
			{Name: "footer.section-second", Select: "#ac-globalfooter .ac-gf-directory-column-section + .ac-gf-directory-column-section", Props: []string{"padding-top"}},
			{Name: "footer.title-text", Select: "#ac-globalfooter .ac-gf-directory-column-section-title > .ac-gf-directory-column-section-title-text", Props: []string{"margin-right", "margin-bottom", "display"}},
			{Name: "footer.list", Select: "#ac-globalfooter .ac-gf-directory-column-section-list", Props: []string{"margin-right", "padding-top", "padding-bottom"}, Widths: []int{1600, 1200, 900}},
			{Name: "footer.item", Select: "#ac-globalfooter .ac-gf-directory-column-section-item", Props: []string{"margin-bottom"}, Widths: []int{1600, 1200, 900}},
			{Name: "footer.item-last", Select: "#ac-globalfooter .ac-gf-directory-column-section-item:last-child", Props: []string{"margin-bottom"}, Widths: []int{1600, 1200, 900}},
			{Name: "footer.footer", Select: "#ac-globalfooter .ac-gf-footer", Props: []string{"padding-top", "padding-bottom"}},
			{Name: "footer.end", Select: "#ac-globalfooter .ac-gf-footer-end", Props: []string{"display"}},
			{Name: "footer.legal-links", Select: "#ac-globalfooter .ac-gf-footer-legal-links", Props: []string{"margin-top", "margin-right"}},
			{Name: "footer.legal-item", Select: "#ac-globalfooter .ac-gf-footer-legal-links-item", Props: []string{"display", "margin-top", "margin-right"}},
			{Name: "footer.legal-rule", Select: "#ac-globalfooter .ac-gf-footer-legal-links-item", Pseudo: "::after", Props: []string{"content", "width", "height", "background-color", "margin-top", "position", "right"}},
			{Name: "footer.legal-link-box", Select: "#ac-globalfooter .ac-gf-footer-legal-link", Props: []string{"padding-left", "padding-right", "margin-left"}},
			{Name: "footer.locale", Select: "#ac-globalfooter .ac-gf-footer-locale", Props: slices.Concat(TYPE, []string{"margin-top", "margin-left", "float", "display"}), Box: true, Position: true},
			{Name: "footer.locale-link", Select: "#ac-globalfooter .ac-gf-footer-locale a, #ac-globalfooter a.ac-gf-footer-locale-link", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			{Name: "footer.section-phone", Select: "#ac-globalfooter .ac-gf-directory-column-section", Props: []string{"padding-left", "padding-right", "padding-bottom", "margin-left", "position"}, Box: true, Widths: []int{390}},
			{Name: "footer.section-rule", Select: "#ac-globalfooter .ac-gf-directory-column-section", Pseudo: "::after", Props: []string{"content", "position", "bottom", "left", "right", "height", "background-color"}, Widths: []int{390}},
			{Name: "footer.title-button", Select: "#ac-globalfooter .ac-gf-directory-column-section-title-button", Props: slices.Concat(TYPE, []string{"display", "padding-top", "padding-bottom", "width", "text-align", "cursor"}), Box: true, Widths: []int{390}},
			{Name: "footer.title-icon", Select: "#ac-globalfooter .ac-gf-directory-column-section-title-icon", Props: []string{"color", "margin-top", "display", "position", "right", "top"}, Box: true, Widths: []int{390}},
			{Name: "footer.list-closed", Select: "#ac-globalfooter .ac-gf-directory-column-section-list", Props: []string{"visibility", "height", "overflow"}, Box: true, Widths: []int{390}},
			// @invariant: the phone's menu, opened: last, as the page stays open after the click
			{Name: "menu.open", Click: "#globalnav-menutrigger-button", Animations: "#globalnav", Select: "#globalnav", Props: slices.Concat([]string{"background-color", "height"}, BLUR), Box: true, Widths: []int{390}},
			{Name: "menu.content", Select: "#globalnav .globalnav-content", Props: slices.Concat(BOX, []string{"height"}, BLUR, MOTION), Box: true, Position: true, Widths: []int{390}},
			// @external: the page held, the logo and the search faded out, the rows come in (each 20ms
			// after the one above)
			{Name: "menu.root", Select: "html", Props: []string{"overflow-y"}, Widths: []int{390}},
			{Name: "menu.body", Select: "body", Props: []string{"overflow-y"}, Widths: []int{390}},
			{Name: "menu.logo", Select: "#globalnav .globalnav-link-apple", Props: slices.Concat([]string{"opacity", "visibility"}, MOTION, []string{"transition-delay"}), Widths: []int{390}},
			{Name: "menu.search", Select: "#globalnav .globalnav-link-search", Props: slices.Concat([]string{"opacity", "visibility"}, MOTION, []string{"transition-delay"}), Widths: []int{390}},
			{Name: "menu.row", Select: "#globalnav .globalnav-item-store .globalnav-submenu-trigger-group", Props: slices.Concat([]string{"opacity", "transform"}, MOTION, []string{"transition-delay"}), Widths: []int{390}},
			{Name: "menu.row-second", Select: "#globalnav .globalnav-item-mac .globalnav-submenu-trigger-group", Props: slices.Concat([]string{"opacity", "transform"}, MOTION, []string{"transition-delay"}), Widths: []int{390}},
			// @external: the rows' places in the open menu, which fills the window: the list, its first row
			// and its second
			{Name: "menu.list", Select: "#globalnav .globalnav-menu-list", Props: slices.Concat(BOX, []string{"padding-left", "padding-right"}), Box: true, Position: true, Widths: []int{390}},
			{Name: "menu.link", Select: "#globalnav .globalnav-item-mac .globalnav-link", Props: slices.Concat(TYPE, BOX, []string{"padding-left", "padding-right", "height", "display"}), Box: true, Position: true, Widths: []int{390}, States: []string{"hover"}},
			{Name: "menu.link-first", Select: "#globalnav .globalnav-item-store .globalnav-link", Props: []string{"padding-top", "margin-bottom"}, Box: true, Position: true, Widths: []int{390}},
			{Name: "menu.item", Select: "#globalnav .globalnav-item-mac", Props: slices.Concat(BOX, []string{"height"}), Box: true, Position: true, Widths: []int{390}},
			{Name: "menu.trigger", Select: "#globalnav .globalnav-menutrigger-button", Props: []string{"color", "opacity"}, Box: true, Position: true, Widths: []int{390}},
			{Name: "menu.trigger-bread-top", Select: "#globalnav .globalnav-menutrigger-bread-top", Props: []string{"stroke-width"}, Box: true, Widths: []int{390}},
			{Name: "menu.trigger-bread-bottom", Select: "#globalnav .globalnav-menutrigger-bread-bottom", Props: []string{"stroke-width"}, Box: true, Widths: []int{390}},
		}),
	}
}
