package console

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleHint is an "i" button beside a word, opening a popover that says
// what the word means.
type ConsoleHint struct {
	// ID is the popover's id, unique on the page.
	ID string
	// AriaLabel: "What A means".
	AriaLabel string
	Hint      string
}

func (t ConsoleHint) Render(w io.Writer) error {
	if t.ID == "" {
		panic("console.ConsoleHint needs an ID")
	}
	if t.AriaLabel == "" {
		panic("console.ConsoleHint needs an AriaLabel")
	}
	if t.Hint == "" {
		panic("console.ConsoleHint needs a Hint")
	}
	// the popover a <span>, not a <div>: a hint sits in a heading or a <p>,
	// where a <div> would close the <p>
	return g.Group{
		h.Button(h.Type("button"), h.Class("cui-console-hint"), h.PopoverTarget(t.ID), h.Aria("label", t.AriaLabel), g.Text("i")),
		h.Span(h.Class("cui-console-popover"), h.ID(t.ID), h.Popover(), g.Text(t.Hint)),
	}.Render(w)
}
