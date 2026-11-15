package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: A buy page, iPhone 18 Pro's: the title, an option (a form selector, its
// parts) at rest and, as a copy added beside it, checked, a color swatch the same way, an
// accordion's question, a price, a button link.
func buy() pagespec.Spec {
	OPTION := "<div class=\"rc-dimension-selector-row form-selector\"><input class=\"form-selector-input\" id=\"cui-option\" type=\"radio\" name=\"cui-option\" checked><label for=\"cui-option\" class=\"form-selector-label\"><span class=\"row\"><span class=\"column form-selector-left-col rf-bfe-selector-left-col\"><span class=\"form-selector-title\">iPhone 18 Pro <span class=\"form-label-small\">6.3-inch display</span></span></span></span></label></div>"
	SWATCH := "<li class=\"colornav-item\"><input type=\"radio\" class=\"colornav-value rc-dimension-colornav-input rf-bfe-product-dimension-colornav-input\" id=\"cui-swatch\" name=\"cui-swatch\" checked><label for=\"cui-swatch\" class=\"colornav-link rc-dimension-colornav-link rf-bfe-product-dimension-colornav-label\"><img alt=\"\" class=\"colornav-swatch\" src=\"data:image/gif;base64,R0lGODlhAQABAAAAACw=\"><span class=\"colornav-label\">Burgundy</span></label></li>"
	return pagespec.Spec{
		URL:  "https://www.apple.com/shop/buy-iphone/iphone-18-pro",
		Wait: 2500,
		Probes: []pagespec.Probe{
			{Name: "page.body", Select: "body", Props: PAGE},
			{Name: "title", Select: "h1", Props: TYPE},
			{Name: "option", Select: ".form-selector:not([data-cui-injected]) .form-selector-label", Props: slices.Concat(TYPE, BOX, MOTION), Box: true, States: []string{"hover"}},
			{Name: "option.focus", Select: ".form-selector:not([data-cui-injected]) .form-selector-label", StateOn: ".form-selector:not([data-cui-injected]) .form-selector-input", Props: slices.Concat(BOX, []string{"outline-color", "outline-style", "outline-width", "outline-offset"}), States: []string{"focus-visible", "focus"}},
			{Name: "option.title", Select: ".form-selector-title", Props: TYPE},
			{Name: "option.small", Select: ".form-label-small", Props: TYPE},
			// @external: its parts: the radio over the tile, the row, its columns (the title's
			// centered, the price's at the top), the title's detail, the price
			{Name: "option.box", Select: ".form-selector:not([data-cui-injected])", Props: []string{"display", "position", "margin-top"}, Box: true},
			{Name: "option.input", Select: ".form-selector:not([data-cui-injected]) .form-selector-input", Props: []string{"position", "top", "left", "opacity", "margin-top", "margin-left", "display"}, Box: true},
			{Name: "option.label", Select: ".form-selector:not([data-cui-injected]) .form-selector-label", Props: slices.Concat(FLEX, []string{"text-align"}), Box: true},
			{Name: "option.left", Select: ".form-selector:not([data-cui-injected]) .form-selector-left-col", Props: []string{"align-self", "padding-right", "text-align"}, Box: true, Position: true},
			{Name: "option.detail", Select: ".form-selector:not([data-cui-injected]) .form-selector-left-col .form-label-small", Props: slices.Concat(TYPE, []string{"display", "margin-top"})},
			{Name: "option.right", Select: ".form-selector:not([data-cui-injected]) .form-selector-right-col", Props: []string{"align-self", "text-align"}, Box: true, Position: true},
			{Name: "option.price", Select: ".form-selector:not([data-cui-injected]) .form-selector-right-col .form-label-small", Props: slices.Concat(TYPE, []string{"display", "margin-top", "text-align"}), Box: true, Position: true},
			{Name: "option.checked", Inject: OPTION, Into: ".rc-dimension-selector-group", Select: ".form-selector-label", Props: slices.Concat(TYPE, BOX), Box: true},
			{Name: "swatch", Select: ".colornav-link:not([data-cui-injected] *) .colornav-swatch", Props: slices.Concat(BOX, []string{"width", "height"}), Box: true},
			{Name: "swatch.link", Select: ".colornav-link", Props: slices.Concat(BOX, []string{"width", "height"}), Box: true, States: []string{"hover"}},
			{Name: "swatch.checked", Inject: SWATCH, Into: ".colornav-items, .rc-dimension-colornav ul, .rc-dimension-colornav", Select: ".colornav-swatch", Props: slices.Concat(BOX, []string{"width", "height", "outline-color", "outline-width", "outline-offset"}), Box: true},
			{Name: "swatch.checked-link", Inject: SWATCH, Into: ".colornav-items, .rc-dimension-colornav ul, .rc-dimension-colornav", Select: ".colornav-link", Props: slices.Concat(BOX, []string{"width", "height"}), Box: true},
			{Name: "swatch.label", Select: ".colornav-label", Props: slices.Concat(TYPE, []string{"display"})},
			// @external: the swatches' list and an item's parts: the radio over the ring, the
			// color's name (hidden but read), the ring focused (on the copy)
			{Name: "swatch.list", Select: ".colornav-items", Props: []string{"display", "flex-wrap", "padding-top", "padding-bottom", "margin-left", "margin-right"}, Box: true},
			{Name: "swatch.item", Select: ".colornav-item:not([data-cui-injected])", Props: []string{"margin-top", "margin-right", "margin-bottom", "margin-left"}, Box: true},
			{Name: "swatch.input", Select: ".colornav-item:not([data-cui-injected]) .colornav-value", Props: []string{"position", "top", "left", "width", "height", "opacity", "margin-top"}, Box: true},
			{Name: "swatch.name", Select: ".colornav-item:not([data-cui-injected]) .colornav-label", Props: []string{"position", "width", "height", "clip", "clip-path", "overflow"}},
			{Name: "swatch.focus", Inject: SWATCH, Into: ".colornav-items, .rc-dimension-colornav ul, .rc-dimension-colornav", Select: ".colornav-link", StateOn: "#cui-swatch", Props: []string{"outline-style", "outline-width", "outline-color", "outline-offset", "border-top-color"}, States: []string{"focus-visible"}},
			{Name: "accordion.item", Select: ".rc-accordion-item", Props: BOX, Box: true},
			{Name: "accordion.button", Select: ".rc-accordion-button", Props: slices.Concat(TYPE, BOX, []string{"width"}), Box: true, States: []string{"hover", "focus-visible"}},
			{Name: "accordion.title", Select: ".rc-accordion-title", Props: TYPE},
			{Name: "accordion.icon", Select: ".rc-accordion-button .icon-plus, .rc-accordion-button .as-accordion-plusicon", Props: []string{"font-family", "font-size", "color", "width", "height"}, Pseudo: "::after"},
			{Name: "accordion.content", Select: ".rc-accordion-box-content", Props: slices.Concat(TYPE, BOX)},
			{Name: "price", Select: ".price-point-fullPrice", Props: TYPE},
			// @external: the page's own price, under its title: its amounts kept whole
			{Name: "price.header", Select: ".rf-bfe-header-price .rc-prices-currentprice", Props: slices.Concat(TYPE, []string{"display"})},
			{Name: "price.amount", Select: ".rf-bfe-header-price .nowrap", Props: slices.Concat(TYPE, []string{"white-space", "display"})},
			{Name: "buttonlink", Select: ".as-buttonlink", Props: slices.Concat(TYPE, BOX), States: []string{"hover"}},
		},
	}
}
