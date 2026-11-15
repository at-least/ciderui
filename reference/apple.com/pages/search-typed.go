package pages

import (
	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The global nav's search typed into, read 2026-10-21 and again 2026-10-23
// (the results' run) with the owner's leave to type on apple.com: the home page's
// search opened as search.go opens it, then one key at a time into its field (`type`),
// nothing sent: "i" (one character, the quick links), "p" (its service asked: one
// section, Suggested Searches), "h" (two sections), "o"; Backspace back to "iph", "ip",
// "i" and the field empty; "x" and "q" (two characters no name of its starts with); and
// the clear button clicked. At each step, every animation and transition its script
// starts in the nav (`animations`, each first seen in ms from the key: one seen within
// 20ms is still held from a step before, a filled animation staying among the
// document's), the search's classes (its shift-<sections before>-<after> and
// with-results), the flyout's height, the results now (each section by its analytics
// region, and the number of its rows), and a row's rest once they have run, with its
// animation's keyframes. At 1600px: the rules of this motion have no width of their own
// (globalheader.css), its slide's distance aside. Read only.
func searchTyped() pagespec.Spec {
	FLY := "#globalnav-submenu-search"
	STEPS := []pagespec.Probe{
		{Name: "i", Type: []string{"i"}},
		{Name: "ip", Type: []string{"p"}},
		{Name: "iph", Type: []string{"h"}},
		{Name: "ipho", Type: []string{"o"}},
		{Name: "iph.back", Type: []string{"Backspace"}},
		{Name: "ip.back", Type: []string{"Backspace"}},
		{Name: "i.back", Type: []string{"Backspace"}},
		{Name: "empty", Type: []string{"Backspace"}},
		{Name: "x", Type: []string{"x"}},
		{Name: "xq", Type: []string{"q"}},
		{Name: "cleared", Click: FLY + " .globalnav-searchfield-reset", Pointer: pagespec.At(200, 430)},
	}
	s := search()
	s.Widths = []int{1600}
	probes := pick(s.Probes, "nav")
	for _, act := range STEPS {
		name := act.Name
		step := act
		step.Animations, step.Select, step.Props = "#globalnav", "#globalnav .globalnav-search", []string{"display"}
		probes = append(probes, step,
			pagespec.Probe{Name: name + ".flyout", Select: FLY, Props: []string{"height"}, Vars: []string{"--r-globalnav-flyout-height", "--r-globalnav-suggested-links-height"}, Box: true},
			pagespec.Probe{Name: name + ".results", Select: FLY + " .globalnav-searchresults", Props: []string{"height"}, Sections: []string{".globalnav-searchresults-container", "data-analytics-region", ".globalnav-searchresults-list-item"}, Box: true},
		)
		// @external: a row at rest once the motion has run (none to read where nothing
		// may be suggested)
		if name != "xq" {
			probes = append(probes, pagespec.Probe{Name: name + ".row", Select: FLY + " .globalnav-searchresults-current .globalnav-searchresults-list-item", Props: []string{"opacity", "transform", "animation-name"}, Keyframes: true})
		}
	}
	s.Probes = probes
	return s
}
