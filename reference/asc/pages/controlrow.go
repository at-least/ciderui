package pages

import "github.com/at-least/ciderui/internal/signedin"

// @external: The newer look's toolkit's control row (--tb-component-ControlRow-*: the
// room after a control in a row of them, a section's after a group), read
// off the root as the page computes it: tokens.go names no ControlRow.
var controlrow = func() signedin.Spec {
	var props []string
	for _, s := range []string{"sm", "md", "lg", "xl"} {
		props = append(props,
			"--tb-component-ControlRow-anatomy-Button-variant-controlSize-"+s+"-marginRight",
			"--tb-component-ControlRow-anatomy-Button-variant-controlSize-"+s+"-plainBackedMarginRight",
			"--tb-component-ControlRow-anatomy-Section-variant-controlSize-"+s+"-marginRight",
		)
	}
	return signedin.Spec{
		Path:   "/apps/<app>/analytics/overview",
		Probes: []signedin.Probe{{Name: "root", Select: "html", Props: props}},
	}
}()
