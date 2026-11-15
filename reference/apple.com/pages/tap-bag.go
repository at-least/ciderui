package pages

import "github.com/at-least/ciderui/internal/pagespec"

// @external: The bag's accordion on a touch screen at 390px, its outer item tapped open,
// then its first compact item (`tap`, the DevTools protocol's touch events: a
// touch screen's tap, not a forced :hover): what the tap leaves hovered, its
// chevron's color against bag-accordion:compact.open (opened by a click, the
// pointer gone: #86868b) and compact.icon-hover (its button's hover forced:
// #6e6e73). Read only.
func tapBag() pagespec.Spec {
	OUTER := ".rc-accordion-item > h2 .rc-accordion-button"
	FIRST := ".rc-accordion-compact > .rc-accordion-item:first-child"

	return pagespec.Spec{
		URL:    bagAccordion().URL,
		Wait:   bagAccordion().Wait,
		Touch:  true,
		Widths: []int{390},
		Probes: []pagespec.Probe{
			{Name: "accordion", Tap: OUTER, Select: OUTER, Props: []string{"cursor"}, Matches: []string{":hover", ":focus-visible"}},
			{Name: "compact", Tap: FIRST + " .rc-accordion-button", Select: FIRST + " .rc-accordion-button", Props: []string{"cursor"}, Matches: []string{":hover", ":active", ":focus", ":focus-visible"}},
			{Name: "compact.icon", Select: FIRST + " .rc-accordion-chevrondown", Props: []string{"color"}},
			// @external: the item the tap opened: its answer's panel
			{Name: "compact.panel", Select: FIRST + " > div", Props: []string{"height"}, Box: true},
		},
	}
}
