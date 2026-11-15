package pages

import "github.com/at-least/ciderui/internal/pagespec"

// @external: System Status's other kinds, which its stylesheet draws and the page shows
// only while a service has one: the legend's mark with its class set to
// light-issue, light-outage and light-maintenance in turn (its available
// one is system-status's), and, at 390px, the toggle's glyph opened
// (`active`).
func systemStatusKinds() pagespec.Spec {
	var kinds []pagespec.Probe
	for _, k := range []string{"issue", "outage", "maintenance"} {
		kinds = append(kinds, pagespec.Probe{Name: k, Widths: []int{1600}, Set: []pagespec.Set{pagespec.Attr(".section-legend li > div", "class", "light-"+k)}, Select: ".section-legend li > div", Props: systemStatusMARK, Image: []string{"background-image"}, Box: true})
	}
	return pagespec.Spec{
		URL:    "https://developer.apple.com/system-status/",
		Wait:   3000,
		Widths: []int{1600, 390},
		Probes: append(kinds, pagespec.Probe{Name: "toggle.active", Set: []pagespec.Set{pagespec.Attr(".section-lights .light-toggle-icon", "classList", "active")}, Select: ".section-lights .light-toggle-icon", Props: systemStatusMARK, Image: []string{"background-image"}, Box: true, Widths: []int{390}}),
	}
}
