package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

const (
	feedbackBOXES = ".form-textbox:has(> .form-textarea)"
	feedbackWRAP  = feedbackBOXES + " > .form-textarea"
	feedbackAREA  = feedbackWRAP + " textarea"
	feedbackLABEL = feedbackBOXES + " > .form-textbox-label"
)

var (
	feedbackEDGE  = []string{"border-top-width", "border-top-style", "border-top-color", "border-bottom-width", "border-top-left-radius"}
	feedbackSPACE = []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "margin-top", "margin-bottom"}
)

// @external: The iPhone feedback form, for apple.com's multi-line field (its
// form-textarea: a bordered box, its textarea inside under its label), at
// rest; its text field's label beside it, for the page's colors. Read only:
// nothing typed, nothing sent. Its states are feedback-states.go.
func feedback() pagespec.Spec {
	return pagespec.Spec{
		URL:  "https://www.apple.com/feedback/iphone/",
		Wait: 2500,
		Probes: []pagespec.Probe{
			{Name: "textarea.box", Select: feedbackBOXES, Props: []string{"position", "height", "margin-bottom"}, Box: true},
			{Name: "textarea", Select: feedbackWRAP, Props: slices.Concat(TYPE, feedbackSPACE, feedbackEDGE, []string{"display", "height", "box-sizing", "background-color", "outline-style"}, MOTION), Box: true, Position: true},
			{Name: "textarea.input", Select: feedbackAREA, Props: slices.Concat(TYPE, feedbackSPACE, feedbackEDGE, []string{"height", "min-height", "width", "box-sizing", "background-color", "resize", "outline-style", "overflow-y"}), Box: true, Position: true},
			{Name: "textarea.label", Select: feedbackLABEL, Props: slices.Concat(TYPE, []string{"position", "top", "left", "pointer-events", "white-space"}, MOTION), Box: true, Position: true},
			{Name: "field.label", Select: ".form-textbox:has(> .form-textbox-input) > .form-textbox-label", Props: TYPE},
			{Name: "field.input", Select: ".form-textbox > .form-textbox-input", Props: slices.Concat(TYPE, []string{"border-top-color", "background-color"})},
		},
	}
}
