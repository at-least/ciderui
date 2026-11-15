package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: apple.com's ribbon, the message above a page's content (the AirPods page's,
// which drops in blue and turns white): read once its motion has run, its animations
// as its stylesheets name them. The iPhone page's, read until 2026-11-13, turned gray;
// it is now a campaign's own part with no link (`ac-ribbon__upgrade-2026`), and the
// store's ribbon turns white too.
func ribbon() pagespec.Spec {
	// @external: its colors run from the start to the end of its keyframes, which name them
	COLORS := []string{"--ribbon-background-color-initial", "--ribbon-background-color", "--ribbon-text-color-initial", "--ribbon-text-color", "--ribbon-link-color-initial", "--ribbon-link-color"}
	return pagespec.Spec{
		URL:  "https://www.apple.com/airpods/",
		Wait: 1000,
		// @external: its colors turn 1.6s after it shows, in 0.8s, and a new width runs it again
		Settle: 3500,
		Probes: []pagespec.Probe{
			{Name: "ribbon", Select: ".ribbon", Props: []string{"overflow", "height"}, Box: true},
			{Name: "ribbon.drop", Select: ".ribbon-drop-wrapper", Props: ANIMATION, Keyframes: true},
			{Name: "ribbon.ground", Select: ".ribbon .ribbon-content-wrapper", Props: slices.Concat(TYPE, BOX, []string{"text-align"}, ANIMATION), Vars: COLORS, Box: true, Keyframes: true},
			{Name: "ribbon.content", Select: ".ribbon .ribbon-content", Props: slices.Concat(TYPE, []string{"width", "max-width", "margin-left", "margin-right"}, ANIMATION), Box: true, Keyframes: true},
			{Name: "ribbon.text", Select: ".ribbon .ribbon-content p", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "ribbon.link", Select: ".ribbon .ribbon-link", Props: slices.Concat(TYPE, LINK, ANIMATION), States: []string{"hover"}, Keyframes: true},
			// @invariant: the pointer on it, a hand's hover, against the forced one: last, as the
			// page keeps it hovered
			{Name: "ribbon.link-hovered", Hover: ".ribbon .ribbon-link", Select: ".ribbon .ribbon-link", Props: slices.Concat([]string{"color"}, LINK, ANIMATION), Matches: []string{":hover"}},
		},
	}
}
