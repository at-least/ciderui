package reference

import (
	"fmt"
)

// Page is what was measured on one of apple.com's pages: reference/apple.com/
// results/<Name>.go, which PAGE=<Name> go run ./reference/apple.com writes.
type Page struct {
	Name, URL, FinalURL, Browser, Date string
	// Widths in the order they were read, ascending.
	Widths []Width
}

// Width is the probes read at one window width, in the spec's order.
type Width struct {
	Width  int
	Probes []Probe
}

// Probe is one probe's reading at one width.
type Probe struct {
	Name string
	// Missing: the probe's selector matched nothing (the run failed on it).
	Missing bool
	// HTML is the head of the element's outerHTML, without the page's words.
	HTML   string
	Styles []Style
	// States are the values with a state forced, each in the order it was read.
	States []State
	// Box is nil where the probe did not read it.
	Box *Box
	// Keyframes is nil where the probe did not read them, empty where the element ran
	// none.
	Keyframes []Keyframes
	// Animations is nil where the probe did not watch, empty where none started.
	Animations []Animation
	Matches    []Match
	Sections   []Section
}

// Style is one computed value: Text, or Image where the probe read a picture as what
// it draws.
type Style struct {
	Name, Text string
	Image      *Image
}

// Image is a data: SVG's viewBox and its shapes, in document order.
type Image struct {
	ViewBox string
	Shapes  []Shape
}

// Shape is one of an SVG's shapes and its paint, inherited where it sets none.
type Shape struct {
	Shape, Fill, Stroke string
}

// State is the computed values with one state forced ("focus+focus-visible" for two).
type State struct {
	Name   string
	Styles []Style
}

// Box is the element's bounding box; At is nil where the probe did not read its
// position.
type Box struct {
	Width, Height float64
	At            *Position
}

// Position is a box's top left corner in the window.
type Position struct {
	Top, Left float64
}

// Keyframes is one @keyframes rule the element runs.
type Keyframes struct {
	Name   string
	Frames []Keyframe
}

// Keyframe is one of a rule's frames: its key text and its declarations.
type Keyframe struct {
	At, Style string
}

// Animation is one animation or transition that started on what a click, tap, type or
// hover opened.
type Animation struct {
	Kind, What, On   string
	Duration, Delay  float64
	Easing, From, To string
	// Observed is nil in records read before the probe recorded it.
	Observed *Observed
}

// Observed is an animation's direction and when it was first seen, in ms after the
// watch started.
type Observed struct {
	Direction string
	Seen      int
}

// Match is whether the element matched a selector.
type Match struct {
	Selector string
	Matches  bool
}

// Section is a section's attribute and its count of rows.
type Section struct {
	Name string
	Rows int
}

func (p Page) probe(name string, width int) (Probe, error) {
	for _, w := range p.Widths {
		if w.Width != width {
			continue
		}
		for _, pr := range w.Probes {
			if pr.Name == name {
				return pr, nil
			}
		}
		return Probe{}, fmt.Errorf("reference %s:%s has nothing at %dpx", p.Name, name, width)
	}
	return Probe{}, fmt.Errorf("reference %s has no width %d", p.Name, width)
}

// Value is apple.com's measured value of prop on the page's probe at width, in state
// if one is named.
func (p Page) Value(name string, width int, prop, state string) (string, error) {
	pr, err := p.probe(name, width)
	if err != nil {
		return "", err
	}
	// @external: the probe recorded the box at rest only
	switch prop {
	case "box-height", "box-width", "box-top", "box-left", "box-right":
		s := ""
		if state != "" {
			s = state + " "
		}
		none := fmt.Errorf("reference %s:%s has no %s%s at %dpx", p.Name, name, s, prop, width)
		if state != "" || pr.Box == nil {
			return "", none
		}
		b := pr.Box
		switch prop {
		case "box-height":
			return px(b.Height), nil
		case "box-width":
			return px(b.Width), nil
		}
		if b.At == nil {
			return "", none
		}
		switch prop {
		case "box-top":
			return px(b.At.Top), nil
		case "box-left":
			return px(b.At.Left), nil
		}
		return px(b.At.Left + b.Width), nil
	}
	styles := pr.Styles
	where := ""
	if state != "" {
		styles = nil
		where = ":" + state
		for _, st := range pr.States {
			if st.Name == state {
				styles = st.Styles
			}
		}
	}
	for _, st := range styles {
		if st.Name != prop {
			continue
		}
		if st.Image != nil {
			return "", fmt.Errorf("reference %s:%s%s's %s at %dpx is a picture, not text", p.Name, name, where, prop, width)
		}
		return st.Text, nil
	}
	return "", fmt.Errorf("reference %s:%s%s has no %s at %dpx", p.Name, name, where, prop, width)
}
