package pages

import (
	"slices"
	"strings"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The store's home: the store's own type scale (its t- classes, added to the
// page), the buttons of its shared button styles (common-css), the header, the
// quick links, the product nav, a shelf's header, its cards (a light one with
// a tag, a dark one, one with a plain eyebrow; their parts and hover), the
// shelf's paddles.
func store() pagespec.Spec {
	T := []string{
		"headline-super", "headline-elevated", "headline", "headline-reduced", "headline-standalone", "eyebrow-super",
		"eyebrow-elevated", "eyebrow", "eyebrow-reduced", "intro-elevated", "intro", "body", "body-tight", "body-reduced",
		"body-reduced-tight", "caption", "label", "callout", "tout", "quote", "quote-reduced", "manifesto",
	}
	BUTTONS := []string{
		"", "button-secondary", "button-neutral", "button-secondary-neutral", "button-super", "button-elevated",
		"button-reduced", "button-block", "button-light", "button-dark",
	}
	STATES := []string{"hover", "active", "focus-visible"}
	// @invariant: a button's probe name: button, or button.<its class without button->
	buttonName := func(c string) string {
		if c == "" {
			return "button"
		}
		return "button." + strings.Replace(c, "button-", "", 1)
	}
	var types, buttons, disabled []pagespec.Probe
	for _, c := range T {
		types = append(types, pagespec.Probe{Name: "type.t-" + c, Inject: `<p class="t-` + c + `">Sample text</p>`, Into: ".rs-page-content, main, body", Props: TYPE})
	}
	for _, c := range BUTTONS {
		buttons = append(buttons, pagespec.Probe{Name: buttonName(c), Inject: `<a href="#" class="button ` + c + `">Buy</a>`, Into: ".rs-page-content, main, body", Props: slices.Concat(TYPE, BOX, TEXT, []string{"width"}, FOCUS, MOTION), Box: true, States: STATES})
	}
	for _, c := range []string{"", "button-secondary", "button-neutral"} {
		disabled = append(disabled, pagespec.Probe{Name: buttonName(c) + ".disabled", Inject: `<button class="button ` + c + `" disabled>Buy</button>`, Into: ".rs-page-content, main, body", Props: slices.Concat(TYPE, BOX, TEXT), Box: true, States: []string{"hover"}})
	}
	// @external: a card by its classes: light with a tag and its picture, dark with a tag,
	// light with a plain eyebrow
	LIGHT := ".rf-ccard-40.rf-card-msgtag-orange:not(.rf-ccard-darkbg):not(.rf-ccard-withvideo)"
	DARK := ".rf-ccard-40.rf-ccard-darkbg.rf-card-msgtag-orange:not(.rf-ccard-withvideo)"
	PLAIN := ".rf-ccard:not([class*=\"msgtag\"]):not(.rf-ccard-darkbg):not(.rf-ccard-withvideo)"
	// @external: the shelf of those cards
	SHELF := ".rs-cardsshelf:has(.rf-ccard-40)"
	return pagespec.Spec{
		URL: "https://www.apple.com/store",
		Probes: slices.Concat([]pagespec.Probe{
			{Name: "type.none", Inject: "<p>Sample text</p>", Into: ".rs-page-content, main, body", Props: TYPE},
		}, types, buttons,
			// @external: disabled: a <button> with the attribute, the shared styles' own
			disabled, []pagespec.Probe{
				{Name: "page.body", Select: "body", Props: PAGE},
				{Name: "page.strong", Select: "main strong, main b, .rs-page-content strong, body strong", Props: TYPE},
				{Name: "header", Select: "h1.rs-shop-header", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
				{Name: "header.tagline", Select: ".rs-shop-header + *, .rs-shop-subheader", Props: TYPE},
				{Name: "quicklink", Select: ".rs-quicklinks-button", Props: slices.Concat(TYPE, BOX), Box: true, States: []string{"hover"}},
				// @external: the quick links: their place (the shelf's inset), title, list and items
				{Name: "quicklinks", Select: ".rs-quicklinks", Props: []string{"margin-left", "margin-right", "padding-top", "padding-bottom"}, Box: true, Position: true},
				{Name: "quicklinks.title", Select: ".rs-quicklinks-title", Props: slices.Concat(TYPE, []string{"margin-left", "margin-right", "padding-bottom", "margin-top", "margin-bottom"}), Box: true},
				{Name: "quicklinks.list", Select: ".rs-quicklinks-list", Props: []string{"display", "flex-wrap", "padding-top", "padding-right", "padding-bottom", "padding-left", "margin-top"}, Box: true},
				{Name: "quicklinks.item", Select: ".rs-quicklinks-item", Props: []string{"display", "padding-top", "padding-right", "padding-bottom", "padding-left"}, Box: true},
				{Name: "quicklinks.item-second", Select: ".rs-quicklinks-item:nth-child(2)", Props: []string{"display"}, Box: true, Position: true},
				{Name: "productnav.title", Select: ".rf-productnav-card-title", Props: slices.Concat(TYPE, []string{"display", "text-align"})},
				{Name: "productnav.image", Select: ".rf-productnav-card-image img", Props: []string{"width", "height"}, Box: true},
				// @external: the product nav: a shelf of small cards (the store's own shelf: its
				// inset, its snapping), each its image above its name
				{Name: "productnav", Select: ".rs-productnav-cardsshelf", Props: []string{"padding-top", "padding-bottom"}, Box: true},
				{Name: "productnav.crop", Select: ".rs-productnav-cardsshelf .rf-cards-scroller-crop", Props: []string{"height", "padding-bottom"}, Box: true},
				{Name: "productnav.card", Select: ".rs-productnav-cardsshelf .rf-cards-scroller-itemview", Props: []string{"margin-right", "transform"}, Box: true, Position: true},
				{Name: "productnav.card-second", Select: ".rs-productnav-cardsshelf [data-core-scroller-item]:nth-child(2) .rf-cards-scroller-itemview", Props: []string{"transform"}, Box: true, Position: true},
				{Name: "productnav.content", Select: ".rs-productnav-cardsshelf .rf-productnav-card-content", Props: []string{"padding-top", "padding-right", "padding-bottom", "padding-left"}, Box: true},
				{Name: "productnav.image-box", Select: ".rs-productnav-cardsshelf .rf-productnav-card-image", Props: []string{"padding-bottom", "text-align"}, Box: true},
				// @external: a shelf of cards, the one holding the 400px cards: its header (its
				// headline and the secondary text after it), its scroller (in a crop that
				// hides its scrollbar), its items, the room around them
				{Name: "shelf", Select: SHELF, Props: []string{"padding-top", "padding-bottom", "position", "background-color"}, Box: true},
				{Name: "shelf.header", Select: SHELF + " .rs-cards-shelf-header", Props: []string{"margin-left", "margin-right", "padding-top", "padding-bottom"}, Box: true, Position: true},
				{Name: "shelf.mainheader", Select: SHELF + " .rs-cards-shelf-mainheader", Props: slices.Concat(TYPE, []string{"display", "margin-top", "margin-bottom"})},
				{Name: "shelf.subheader", Select: SHELF + " .rs-cards-shelf-secondaryheader", Props: slices.Concat(TYPE, []string{"display"})},
				{Name: "shelf.crop", Select: SHELF + " .rf-cards-scroller-crop", Props: []string{"overflow-x", "overflow-y", "padding-bottom", "height"}, Box: true},
				{Name: "shelf.scroller", Select: SHELF + " .rf-cards-scroller-content", Props: []string{"overflow-x", "overflow-y", "scroll-snap-type", "overscroll-behavior-x"}, Box: true},
				{Name: "shelf.platter", Select: SHELF + " .rf-cards-scroller-platter", Props: []string{"display", "padding-top", "padding-bottom"}, Box: true},
				{Name: "shelf.item", Select: SHELF + " [data-core-scroller-item]", Props: []string{"scroll-snap-align", "scroll-snap-stop"}, Box: true},
				{Name: "shelf.card", Select: SHELF + " .rf-cards-scroller-itemview", Props: []string{"margin-right", "transform"}, Box: true, Position: true},
				{Name: "shelf.card-second", Select: SHELF + " [data-core-scroller-item]:nth-child(2) .rf-cards-scroller-itemview", Props: []string{"transform"}, Box: true, Position: true},
				{Name: "shelf.card-last", Select: SHELF + " [data-core-scroller-item]:last-child .rf-cards-scroller-itemview", Props: []string{"margin-right", "padding-right", "transform"}, Box: true},
				// @external: the cards: a light one with a tag (rf-card-msgtag-orange: "NEW"), its
				// picture filling it, its info over the picture; a dark one
				// (rf-ccard-darkbg); a light one with a plain eyebrow ("APPLE
				// SPECIALIST"); a header first in its info (no eyebrow above it); a dark
				// card's plain eyebrow on a copy the probe adds, apple.com's classes (no
				// dark card on the page has one)
				{Name: "card", Select: LIGHT + " .rf-ccard-content", Props: slices.Concat(BOX, []string{"transform", "overflow", "display", "flex-direction", "position", "color", "cursor"}, MOTION), Box: true, States: []string{"hover", "focus-within"}},
				{Name: "card.dark", Select: DARK + " .rf-ccard-content", Props: slices.Concat(BOX, []string{"transform", "overflow"}), Box: true, States: []string{"hover"}},
				{Name: "card.media", Select: LIGHT + " .rf-ccard-img-full", Props: []string{"display", "object-fit", "position"}, Box: true, Widths: []int{1600, 1200, 900}},
				{Name: "card.info", Select: LIGHT + " .rf-ccard-content-info", Props: []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "position"}, Box: true, Position: true},
				{Name: "card.tag", Select: LIGHT + " .rf-ccard-content-header-eyebrow", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom", "padding-top", "padding-bottom", "text-transform"})},
				{Name: "card.eyebrow", Select: PLAIN + " .rf-ccard-content-header-eyebrow", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom", "padding-top", "padding-bottom", "text-transform"})},
				{Name: "card.header", Select: LIGHT + " .rf-ccard-content-header", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom", "padding-top", "padding-bottom"})},
				{Name: "card.header-first", Select: ".rf-ccard-content-info > .rf-ccard-content-header:first-child", Props: slices.Concat(TYPE, []string{"padding-top"})},
				{Name: "card.link", Select: LIGHT + " .rf-ccard-content-headerlink", Props: slices.Concat(TYPE, []string{"display"}, LINK, []string{"cursor"}), States: []string{"hover"}},
				{Name: "card.desc", Select: LIGHT + " .rf-ccard-content-desc", Props: slices.Concat(TYPE, []string{"margin-top", "padding-top", "padding-bottom"})},
				{Name: "card.description", Select: LIGHT + " .rf-ccard-content-desccontent", Props: slices.Concat(TYPE, []string{"display", "padding-top", "padding-bottom"})},
				{Name: "card.price", Select: LIGHT + " .rf-ccard-content-descprice", Props: TYPE},
				{Name: "card.dark-tag", Select: DARK + " .rf-ccard-content-header-eyebrow", Props: []string{"color"}},
				{Name: "card.dark-eyebrow", Inject: "<div class=\"rf-ccard rf-ccard-40 rf-ccard-darkbg rf-card-bgblack\"><div class=\"rf-ccard-content\"><div class=\"rf-ccard-content-info\"><h3 class=\"rf-ccard-content-header-eyebrow\">Apple Specialist</h3></div></div></div>", Into: ".rs-page-content, main, body", Select: ".rf-ccard-content-header-eyebrow", Props: []string{"color"}},
				{Name: "card.dark-header", Select: DARK + " .rf-ccard-content-header", Props: []string{"color"}},
				{Name: "card.dark-link", Select: DARK + " .rf-ccard-content-headerlink", Props: []string{"color"}},
				{Name: "card.dark-desc", Select: DARK + " .rf-ccard-content-desc", Props: []string{"color"}},
				{Name: "card.dark-price", Select: DARK + " .rf-ccard-content-descprice", Props: []string{"color"}},
				// @external: the shelf's paddles, at its sides over its scroller: the next one (the
				// previous disabled at the start), its place and motion; the previous
				// hidden
				{Name: "shelf.paddle", Select: SHELF + " .paddlenav-arrow-next", Props: slices.Concat(BOX, []string{"width", "height", "color", "opacity", "transform", "position", "top", "right"}, MOTION), Box: true, Position: true, States: []string{"hover", "active"}},
				{Name: "shelf.paddle-previous", Select: SHELF + " .paddlenav-arrow-previous", Props: slices.Concat([]string{"opacity", "transform", "position", "left", "top", "margin-top"}, MOTION), Box: true, Position: true},
				{Name: "shelf.paddles", Select: SHELF + " .rf-cards-scroller", Props: []string{"position"}, Box: true, Position: true},
			}),
	}
}
