package pages

import "github.com/at-least/ciderui/internal/pagespec"

// @external: Where apple.com's breakpoints fall: a headline (the content's type) and the
// global nav (its own breakpoint) stepped across each boundary its stylesheets
// name (734/735, 1068/1069, 1440/1441; the nav's 833/834).
func breakpoints() pagespec.Spec {
	TYPE := []string{"font-size", "line-height", "letter-spacing", "font-weight", "font-family"}
	return pagespec.Spec{
		URL:    "https://www.apple.com/",
		Widths: []int{733, 734, 735, 736, 832, 833, 834, 835, 1067, 1068, 1069, 1070, 1439, 1440, 1441, 1442},
		Probes: []pagespec.Probe{
			{Name: "type.headline", Inject: "<p class=\"typography-headline\">Headline</p>", Props: TYPE},
			{Name: "globalnav", Select: "#globalnav", Props: []string{"height"}, Vars: []string{"--r-globalnav-height"}, Box: true},
		},
	}
}
