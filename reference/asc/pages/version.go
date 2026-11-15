package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

var version = func() signedin.Spec {
	tab := []string{"height", "padding-bottom", "border-bottom-width", "border-bottom-style", "border-bottom-color"}
	button := slices.Concat(signedin.Type, signedin.Box, []string{"height", "cursor", "opacity"})
	choice := []string{"width", "height", "appearance", "accent-color", "margin-top", "margin-right", "margin-left", "vertical-align"}
	return signedin.Spec{
		Path: "/apps/<app>/distribution/ios/version/inflight",
		Probes: []signedin.Probe{
			{Name: "tab.current", Select: "[role=tablist] a[role=tab][aria-selected=true]", Props: slices.Concat(signedin.Type, signedin.Link, tab), Box: true},
			{Name: "tab", Select: "[role=tablist] a[role=tab][aria-selected=false]", Props: slices.Concat(signedin.Type, signedin.Link, tab), Box: true},
			{Name: "tab.button.current", Select: "[role=tablist] button[role=tab][aria-selected=true]", Props: tab, Box: true},
			{Name: "tab.button.current.text", Select: "[role=tablist] button[role=tab][aria-selected=true] > div", Props: signedin.Type, Box: true},
			{Name: "tab.button", Select: "[role=tablist] button[role=tab][aria-selected=false]", Props: tab, Box: true},
			{Name: "tab.button.text", Select: "[role=tablist] button[role=tab][aria-selected=false] > div", Props: signedin.Type, Box: true},
			{Name: "button.primary", Select: `button:not([type])[class*="Text__BaseText"]`, Props: button, Box: true},
			{Name: "button.neutral", Select: `button:not([type])[class*="Text__BaseText"]`, Nth: 1, Props: button, Box: true},
			{Name: "button.disabled", Select: `button[type=button][class*="Text__BaseText"]:disabled`, Props: button, Box: true},
			{Name: "help", Select: `button[role=button][type=button][class*="Text__BaseText"]`, Props: slices.Concat(signedin.Type, []string{"width", "height", "background-color", "border-top-left-radius"}), Box: true},
			{Name: "field", Select: `input[class*="Input__BaseInput"]:not([type]):enabled`, Props: slices.Concat(signedin.Type, signedin.Box, []string{"height"}), Box: true},
			{Name: "field.disabled", Select: `input[class*="Input__BaseInput"]:disabled`, Props: slices.Concat(signedin.Type, signedin.Box, []string{"height"}), Box: true},
			{Name: "field.label", Select: `label[class*="Text__BaseText"]`, Props: slices.Concat(signedin.Type, []string{"display", "margin-bottom"}), Box: true},
			{Name: "textarea", Select: `textarea[class*="Input__BaseInput"]`, Props: slices.Concat(signedin.Type, signedin.Box, []string{"height", "resize"}), Box: true},
			{Name: "field.count", Select: "p[role=status]", Props: slices.Concat(signedin.Type, []string{"margin-top", "text-align"}), Box: true},
			{Name: "checkbox", Select: "input[type=checkbox]", Props: choice, Box: true},
			{Name: "radio", Select: "input[type=radio]", Props: choice, Box: true},
		},
	}
}()
