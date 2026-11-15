package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

// @external: The owner's feedback, already read: opening an unread item marks it read.
var item = func() signedin.Spec {
	sides := []string{"border-left-width", "border-left-style", "border-left-color", "border-right-width", "border-right-style", "border-right-color"}
	button := slices.Concat(signedin.Type, signedin.Box, sides, []string{"height", "cursor", "opacity"})
	attachment := `a[class*="AttachmentsSection__AttachmentGlyphLink"]`
	return signedin.Spec{
		Path: "/feedback/<app>",
		Probes: []signedin.Probe{
			{Name: "item.current", Select: `[class*="ListItemWrapper"] a[href^="/feedback/"] [class^="Card-sc-"]`, Props: slices.Concat([]string{"background-color"}, sides, []string{"border-bottom-width", "border-bottom-color"}), Box: true},
			{Name: "byline", Select: `div:has(> [class^="Monogram-sc-"])`, Props: slices.Concat(signedin.Box, []string{"height"}), Box: true},
			{Name: "byline.monogram", Select: ".sb-avatar__text", Props: slices.Concat(signedin.Type, []string{"width", "height", "background-color", "border-top-left-radius"}), Box: true},
			{Name: "byline.name", Select: `[class^="Name-sc-"]`, Props: signedin.Type, Box: true},
			{Name: "byline.date", Select: `[class^="Date-sc-"]`, Props: signedin.Type, Box: true},
			{Name: "title.row", Select: `div:has(> p[class*="ItemDetail__FeedbackTitle"])`, Props: slices.Concat(signedin.Box, []string{"height"}), Box: true},
			{Name: "title", Select: `p[class*="ItemDetail__FeedbackTitle"]`, Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "actions", Select: `button[class*="ActionsMenuStyles__CustomPrimaryButton"]`, Props: button, Box: true},
			{Name: "status", Select: `[class^="Status-sc-"]`, Props: slices.Concat(signedin.Box, sides), Box: true},
			{Name: "status.label", Select: `[class*="Status__StatusLabel"]`, Props: slices.Concat(signedin.Type, []string{"width", "min-width"}), Box: true},
			{Name: "status.value", Select: `[class*="Status__StatusValue"]`, Props: signedin.Type, Box: true},
			{Name: "section", Select: `[class^="Section-sc-"]`, Props: signedin.Box, Box: true},
			{Name: "section.lead", Select: `[class^="Description-sc-"]`, Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "section.label", Select: `[class^="Section-sc-"] [class^="Header-sc-"]`, Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "section.text", Select: `[class^="Section-sc-"] [class^="Body-sc-"]`, Props: slices.Concat(signedin.Type, []string{"white-space", "margin-bottom"}), Box: true},
			{Name: "attachment.glyph", Select: attachment + " svg", Props: []string{"width", "height", "color", "fill"}, Box: true},
			{Name: "attachment.link", Select: attachment, Nth: 1, Props: slices.Concat(signedin.Type, signedin.Link), Box: true},
			{Name: "attachment.line", Select: "p.size-time", Props: signedin.Type, Box: true},
			{Name: "reply", Select: `[class^="FeedbackResponse-sc-"]`, Props: slices.Concat(signedin.Box, sides, []string{"position", "bottom"}), Box: true},
			{Name: "reply.field", Select: `textarea[class*="FormTextArea"]`, Props: slices.Concat(signedin.Type, signedin.Box, sides, []string{"height", "resize"}), Box: true},
			{Name: "reply.attach", Select: `button[class*="feedbackResponseForm__AttachmentButton"] svg`, Props: []string{"width", "height", "color", "fill"}, Box: true},
			{Name: "reply.send.disabled", Select: `[class^="ActionRow-sc-"] button[class*="PrimaryButton__Button"]:disabled`, Props: button, Box: true},
		},
	}
}()
