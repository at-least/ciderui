package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The iPhone feedback form's multi-line field focused (its textarea clicked
// once, at the first width: the class apple.com's script sets on its box),
// and left empty (the focus taken off: its script marks it in error, with
// its message). Read only: nothing typed, nothing sent.
func feedbackStates() pagespec.Spec {
	return pagespec.Spec{
		URL:    "https://www.apple.com/feedback/iphone/",
		Widths: []int{1600, 390},
		Wait:   2500,
		Probes: []pagespec.Probe{
			{Name: "focused", Click: feedbackAREA, Select: feedbackWRAP, Props: slices.Concat(feedbackSPACE, feedbackEDGE, []string{"background-color", "outline-style", "height"}), Box: true, Widths: []int{1600}},
			{Name: "focused.input", Select: feedbackAREA, Props: slices.Concat(feedbackSPACE, []string{"height"}), Box: true, Position: true, Widths: []int{1600}},
			{Name: "focused.label", Select: feedbackLABEL, Props: slices.Concat(TYPE, []string{"top", "left"}), Box: true, Position: true, Widths: []int{1600}},
			{Name: "error", Blur: true, Select: feedbackWRAP, Props: slices.Concat(feedbackEDGE, []string{"background-color", "margin-bottom", "height"}), Box: true},
			{Name: "error.label", Select: feedbackLABEL, Props: slices.Concat(TYPE, []string{"top", "left"})},
			{Name: "error.wrapper", Select: feedbackBOXES + " > .form-message-wrapper", Props: slices.Concat(BOX, []string{"display"}), Box: true, Position: true},
			{Name: "error.message", Select: feedbackBOXES + " .form-message", Props: TYPE},
		},
	}
}
