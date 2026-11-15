package pages

import (
	"fmt"
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @invariant: The property lists the page specs read. A list built from others is
// slices.Concat's fresh slice, never an append onto a shared one.
var (
	TYPE   = []string{"font-family", "font-size", "line-height", "font-weight", "letter-spacing", "color"}
	BOX    = []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "margin-top", "margin-bottom", "border-top-width", "border-top-style", "border-top-color", "border-bottom-width", "border-bottom-style", "border-bottom-color", "border-top-left-radius", "background-color", "box-shadow", "min-width", "min-height", "max-width"}
	MOTION = []string{"transition-property", "transition-duration", "transition-timing-function"}
	BLUR   = []string{"backdrop-filter", "-webkit-backdrop-filter"}
	LINK   = []string{"text-decoration-line", "text-decoration-color", "text-underline-offset"}
	FOCUS  = []string{"outline-style", "outline-color", "outline-width", "outline-offset"}
	// @invariant: the page's own: its text, its ground, how its text is drawn
	PAGE = slices.Concat(TYPE, []string{"background-color", "-webkit-font-smoothing", "font-synthesis", "text-rendering"})
	// @invariant: a control's text: how it sits, wraps and shows a link
	TEXT = []string{"display", "text-align", "white-space", "text-decoration-line", "cursor", "vertical-align", "opacity"}
	// @invariant: a flex or grid box's own layout
	FLEX      = []string{"display", "flex-direction", "flex-wrap", "align-items", "justify-content", "row-gap", "column-gap"}
	ANIMATION = []string{"animation-name", "animation-duration", "animation-delay", "animation-timing-function", "animation-fill-mode", "animation-iteration-count"}
)

// @invariant: a spec read again on a touch screen (`touch`): its probes of these names,
// in its order (a probe that clicks is named with the ones after it), each
// with states forcing these alone (:hover; nil is hover)
func touched(spec pagespec.Spec, names []string, states []string) pagespec.Spec {
	if states == nil {
		states = []string{"hover"}
	}
	var probes []pagespec.Probe
	for _, p := range spec.Probes {
		if !slices.Contains(names, p.Name) {
			continue
		}
		if p.States != nil {
			kept := []string{}
			for _, s := range p.States {
				if slices.Contains(states, s) {
					kept = append(kept, s)
				}
			}
			p.States = kept
		}
		probes = append(probes, p)
	}
	var missing []string
	for _, n := range names {
		if !slices.ContainsFunc(probes, func(p pagespec.Probe) bool { return p.Name == n }) {
			missing = append(missing, n)
		}
	}
	if len(missing) > 0 {
		panic(fmt.Sprintf("no probe %v in %s's spec", missing, spec.URL))
	}
	spec.Touch = true
	spec.Probes = probes
	return spec
}

// @invariant: the probes of these names, in the probes' order, not the names'
// (probes.filter(p => names.includes(p.name)))
func pick(probes []pagespec.Probe, names ...string) []pagespec.Probe {
	var out []pagespec.Probe
	for _, p := range probes {
		if slices.Contains(names, p.Name) {
			out = append(out, p)
		}
	}
	return out
}
