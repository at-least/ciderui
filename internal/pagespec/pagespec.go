// Package pagespec is what reference/apple.com reads off a live page: a Spec (the page,
// its widths, how it is loaded) and its Probes (an element, what is read of it, and
// what is done to the page first). It is a package under internal/ so that
// ciderui-check, which holds every literal of the module's other types to every field,
// leaves a probe to name only the fields it uses.
package pagespec

import "encoding/json"

// Spec is one page read at each of its widths.
type Spec struct {
	URL string `json:"url"`
	// Browser is internal/probe's name for the browser; "" is chromium.
	Browser string `json:"browser,omitempty"`
	// ColorScheme is "dark" or "light": the page read in that scheme
	// (prefers-color-scheme), from its load.
	ColorScheme string `json:"colorScheme,omitempty"`
	// Touch reads the page on a touch screen, from its load.
	Touch bool `json:"touch,omitempty"`
	// Widths are the window's widths, in this order; nil is DefaultWidths.
	Widths []int `json:"widths,omitempty"`
	// Wait is milliseconds waited once, after the load.
	Wait int `json:"wait,omitempty"`
	// Settle is milliseconds waited after each resize; 0 is 800.
	Settle int     `json:"settle,omitempty"`
	Probes []Probe `json:"probes"`
}

// DefaultWidths is one width inside each of apple.com's bands.
var DefaultWidths = []int{1600, 1200, 900, 390}

// ReadWidths is s.Widths, or DefaultWidths where it is nil.
func (s Spec) ReadWidths() []int {
	if s.Widths == nil {
		return DefaultWidths
	}
	return s.Widths
}

// Probe is one element read at each width. Its JSON is what the page is handed
// (measure's in-page functions read these keys); a Point is not handed.
type Probe struct {
	Name string `json:"name"`
	// Select is a selector; " >>> " steps into a shadow root.
	Select string `json:"select,omitempty"`
	// Inject is markup appended to Into (or body); Select then picks inside it.
	Inject string `json:"inject,omitempty"`
	Into   string `json:"into,omitempty"`
	// StateOn is the element States are forced on, if not the probe's own.
	StateOn string `json:"stateOn,omitempty"`
	// Frame picks the frame whose URL holds it.
	Frame  string   `json:"frame,omitempty"`
	Pseudo string   `json:"pseudo,omitempty"`
	Props  []string `json:"props,omitempty"`
	// Vars are custom properties, read trimmed.
	Vars []string `json:"vars,omitempty"`
	// Image are properties whose value is a base64 SVG, read as what it draws.
	Image     []string `json:"image,omitempty"`
	Keyframes bool     `json:"keyframes,omitempty"`
	// Sections is [section, attribute, row]: each section's attribute and its rows' count.
	Sections []string `json:"sections,omitempty"`
	Matches  []string `json:"matches,omitempty"`
	Box      bool     `json:"box,omitempty"`
	Position bool     `json:"position,omitempty"`
	// States are pseudo-classes forced one at a time; "focus+focus-visible" is two at once.
	States []string `json:"states,omitempty"`
	// Scroll is the window's scroll before the read (and before a hover).
	Scroll int   `json:"scroll,omitempty"`
	Blur   bool  `json:"blur,omitempty"`
	Set    []Set `json:"set,omitempty"`
	// Hover is a selector the pointer moves to the middle of, at every width; HoverAt
	// a point instead.
	Hover   string `json:"hover,omitempty"`
	HoverAt Point  `json:"-"`
	// Hovered is milliseconds waited after the hover; 0 is 1500.
	Hovered int    `json:"hovered,omitempty"`
	Pointer Point  `json:"-"`
	Tap     string `json:"tap,omitempty"`
	// Type is keys pressed one at a time.
	Type []string `json:"type,omitempty"`
	// Typed is milliseconds waited after the keys; 0 is 2500.
	Typed int `json:"typed,omitempty"`
	// Click is a selector clicked once; ClickAt a point of the window instead.
	Click   string `json:"click,omitempty"`
	ClickAt Point  `json:"-"`
	// Clicked is milliseconds waited after the click; 0 is two frames and 900ms.
	Clicked int `json:"clicked,omitempty"`
	// Animations is a selector: every animation the action starts inside it is recorded.
	Animations string `json:"animations,omitempty"`
	// Widths are the spec's widths the probe is read at; nil is all of them.
	Widths []int `json:"widths,omitempty"`
}

// Point is a point of the window, given the window's width.
type Point func(width int) [2]float64

// At is the point x, y at every width.
func At(x, y float64) Point { return func(int) [2]float64 { return [2]float64{x, y} } }

// Set is one change a probe makes to the page before it reads: an attribute set or
// removed; the name "value" sets an input's value, "classList" adds a class.
type Set struct {
	selector, name string
	value          *string
}

// Attr sets name to value on the element selector picks.
func Attr(selector, name, value string) Set { return Set{selector, name, &value} }

// NoAttr removes name from the element selector picks.
func NoAttr(selector, name string) Set { return Set{selector, name, nil} }

// MarshalJSON is [selector, name, value], value null for NoAttr.
func (s Set) MarshalJSON() ([]byte, error) {
	return json.Marshal([]any{s.selector, s.name, s.value})
}
