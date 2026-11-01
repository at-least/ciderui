package console

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleHint is cui-console-hint (components/console/console-hint): App Store
// Connect's newer look's info button, a small "i" beside a word that opens a
// popover saying what the word means. The popover is the button's
// popovertarget: a click or a tap opens it, Escape or a click elsewhere
// closes it, no script.
type ConsoleHint struct {
	// ID is the popover's id, unique on the page. Required: an empty one
	// panics.
	ID string
	// Label is the button's accessible name (aria-label), what it is about.
	// Required: an empty one panics.
	Label string
	// Hint is what the word means: the popover's text. Required: an empty
	// one panics.
	Hint string
}

// Render writes the button and its popover.
func (t ConsoleHint) Render(w io.Writer) error {
	if t.ID == "" {
		panic("console.ConsoleHint needs an ID")
	}
	if t.Label == "" {
		panic("console.ConsoleHint needs a Label")
	}
	if t.Hint == "" {
		panic("console.ConsoleHint needs a Hint")
	}
	// the popover a <span>: a hint sits in a heading or a <p>, where a <div>
	// would close the <p>
	return g.Group{
		h.Button(h.Type("button"), h.Class("cui-console-hint"), h.PopoverTarget(t.ID), h.Aria("label", t.Label), g.Text("i")),
		h.Span(h.Class("cui-console-popover"), h.ID(t.ID), h.Popover(), g.Text(t.Hint)),
	}.Render(w)
}
