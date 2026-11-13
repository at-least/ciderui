package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/attr"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsolePlayButton is a button that plays one audio clip and stops it, written
// at rest: paused. It holds both glyphs and both labels; the page's script
// sets data-state="playing" (and an icon-only one's aria-label) and, while
// playing, --cui-console-play-button-progress (0% to 100%), the ring round the glyph.
type ConsolePlayButton struct {
	// AriaLabel is an icon-only button's name at rest, what pressing it does
	// ("Play"); "" with a Label, whose words name it.
	AriaLabel string
	// Label is the words after the glyph while paused, PlayingLabel while
	// playing; "" both for a round button of the glyph alone.
	Label        string
	PlayingLabel string
	Size         ConsoleBtnSize
	// Attrs: the root's other attributes, for the page's script (data-*, an
	// htmx attribute).
	Attrs g.Node
}

func (p ConsolePlayButton) Render(w io.Writer) error {
	labelled := p.Label != ""
	if labelled != (p.PlayingLabel != "") {
		panic("console.ConsolePlayButton needs a Label and a PlayingLabel, or neither")
	}
	if labelled == (p.AriaLabel != "") {
		panic("console.ConsolePlayButton needs an AriaLabel without a Label, and none with one")
	}
	root := "cui-console-btn cui-console-btn-icon cui-console-play-button"
	if labelled {
		root = "cui-console-btn cui-console-play-button"
	}
	return h.Button(h.Class(root),
		attr.Optional("data-size", p.Size.value), h.Data("shape", "capsule"),
		h.Type("button"), attr.Optional("aria-label", p.AriaLabel), p.Attrs,
		h.Span(h.Class("cui-console-play-button-glyph"), lucide.Play.Class("cui-console-play-button-paused"), lucide.Square.Class("cui-console-play-button-playing")),
		g.If(labelled, g.Group{
			h.Span(h.Class("cui-console-play-button-paused"), g.Text(p.Label)),
			h.Span(h.Class("cui-console-play-button-playing"), g.Text(p.PlayingLabel)),
		}),
	).Render(w)
}
