package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

func bag() pagespec.Spec {
	STATES := []string{"hover", "active", "focus-visible"}
	return pagespec.Spec{
		URL:  "https://www.apple.com/shop/bag",
		Wait: 2500,
		Probes: []pagespec.Probe{
			{Name: "page.body", Select: "body", Props: PAGE},
			{Name: "header", Select: ".rs-bag-header", Props: slices.Concat(TYPE, []string{"margin-bottom"})},
			{Name: "message", Select: ".rs-bagempty-message", Props: TYPE},
			{Name: "form-button", Select: ".rs-bagempty .form-button:not(.form-button-secondary)", Props: slices.Concat(TYPE, BOX, TEXT, []string{"width"}, FOCUS, MOTION), Box: true, States: STATES},
			{Name: "form-button.disabled", Inject: "<button class=\"form-button\" disabled>Sign In</button>", Into: ".rs-bagempty-actions", Props: slices.Concat(TYPE, BOX, TEXT), Box: true, States: []string{"hover"}},
			{Name: "form-button.secondary", Select: ".rs-bagempty .form-button-secondary", Props: slices.Concat(TYPE, BOX, []string{"width", "display"}, FOCUS, MOTION), Box: true, States: STATES},
			{Name: "accordion.button", Select: ".rc-accordion-button", Props: slices.Concat(TYPE, BOX, []string{"width"}), Box: true, States: []string{"hover"}},
			{Name: "accordion.item", Select: ".rc-accordion-item", Props: BOX, Box: true},
			{Name: "ribbon", Select: ".ribbon-content-wrapper", Props: slices.Concat(TYPE, BOX), Box: true},
		},
	}
}
