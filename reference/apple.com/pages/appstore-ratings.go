package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: the App Store's web page for an app (Pages): its ratings' bars, each a
// track with a filled part as wide as its share (an inline width).
func appstoreRatings() pagespec.Spec {
	ROW := ".numbers__star-graph__row"
	return pagespec.Spec{
		URL:  "https://apps.apple.com/us/app/pages/id361309726",
		Wait: 3000,
		Probes: []pagespec.Probe{
			{Name: "graph", Select: ".numbers__star-graph", Props: slices.Concat(FLEX, []string{"grid-template-columns", "width"}), Box: true},
			{Name: "row", Select: ROW, Props: slices.Concat(BOX, FLEX, []string{"grid-template-columns", "height"}), Box: true},
			{Name: "track", Select: ROW + "__bar", Props: slices.Concat(BOX, []string{"height", "width", "border-bottom-left-radius", "overflow", "margin-left"}), Box: true},
			{Name: "fill", Select: ROW + "__bar__foreground", Props: []string{"background-color", "height", "width", "border-top-left-radius", "border-top-right-radius"}, Box: true},
		},
	}
}
