package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

// @invariant: Its emails card and the team's address are not read.
var overview = func() signedin.Spec {
	pad := []string{"padding-top", "padding-right", "padding-bottom", "padding-left"}
	row := `li[class*="TeamMembershipDetails-module__TeamMembershipDetailsRow"]`
	return signedin.Spec{
		Path: "/account",
		Probes: []signedin.Probe{
			{Name: "page", Select: "body", Props: signedin.Page},
			{Name: "ground", Select: "div.App.background", Props: []string{"background-color"}, Box: true},
			{Name: "tab", Select: `a[class*="DeveloperTabNav-module__TabNavItemLink"]`, Props: slices.Concat(signedin.Type, []string{"width", "height"}), Box: true},
			{Name: "tab.glyph", Select: `figure[class*="DeveloperTabNav-module__TabNavItemIcon"]`, Props: []string{"width", "height", "margin-bottom", "background-color"}, Box: true},
			{Name: "paddle.disabled", Select: `button[class*="DeveloperTabNav-module__TabNavPaddle--"]:disabled`, Props: []string{"width", "height", "background-color", "color", "border-top-left-radius", "opacity"}, Box: true},
			{Name: "card", Select: `section[class*="Card-module__Card--"]`, Props: slices.Concat(pad, []string{"margin-bottom", "background-color", "border-top-left-radius", "box-shadow", "border-top-width"}), Box: true},
			{Name: "card.title", Select: ".ProgramResourcesCard h4", Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "section.picture", Select: ".CardSection img", Props: []string{"width", "height", "margin-bottom"}, Box: true},
			{Name: "section.title", Select: ".CardSection h5", Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "section.text", Select: ".CardSection p", Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "section.link", Select: `li[class*="Card-module__CardSectionLink"] a`, Props: slices.Concat(signedin.Type, signedin.Link), Box: true},
			{Name: "row", Select: row, Props: slices.Concat(pad, []string{"margin-bottom", "border-bottom-width", "border-bottom-style", "border-bottom-color"}), Box: true},
			{Name: "row.label", Select: row + ` > div[class*="TeamMembershipDetailsColumn"]`, Props: signedin.Type, Box: true},
			{Name: "row.label.word", Select: row + ` > div[class*="TeamMembershipDetailsColumn"] span.lighter`, Props: signedin.Type},
			{Name: "row.value", Select: row + ` > div[class*="TeamMembershipDetailsColumn"].column`, Props: signedin.Type, Box: true},
			{Name: "info", Select: "button.form-tooltip-button", Props: []string{"width", "height", "color", "background-color"}, Box: true},
			{Name: "add", Select: "a.icon-pluscircle", Props: slices.Concat(signedin.Type, signedin.Link), Box: true},
		},
	}
}()
