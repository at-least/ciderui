package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsolePlay is a button that plays one audio clip and stops it, written
// at rest: paused. It holds both glyphs and both labels; the page's script
// sets data-state="playing" (and an icon-only one's aria-label) and, while
// playing, --cui-console-play-progress (0% to 100%), the ring round the glyph.
type ConsolePlay struct {
	// AriaLabel is an icon-only button's name at rest, what pressing it does
	// ("Play"); "" with a Label, whose words name it.
	AriaLabel string
	// Label is the words after the glyph while paused, PlayingLabel while
	// playing; "" both for a round button of the glyph alone.
	Label        string
	PlayingLabel string
	Size         ConsoleBtnSize
	// Attrs: the root's other attributes, for the page's script (data-*, an
	// htmx attribute); nil for none.
	Attrs g.Node
	Class string
}

func (p ConsolePlay) Render(w io.Writer) error {
	labelled := p.Label != ""
	if labelled != (p.PlayingLabel != "") {
		panic("console.ConsolePlay needs a Label and a PlayingLabel, or neither")
	}
	if labelled == (p.AriaLabel != "") {
		panic("console.ConsolePlay needs an AriaLabel without a Label, and none with one")
	}
	root := "cui-console-btn cui-console-btn-icon cui-console-play"
	if labelled {
		root = "cui-console-btn cui-console-play"
	}
	return h.Button(class.Root(root, p.Class),
		optionalAttr("data-size", p.Size.value), h.Data("shape", "capsule"),
		h.Type("button"), optionalAttr("aria-label", p.AriaLabel), p.Attrs,
		h.Span(h.Class("cui-console-play-glyph"), lucide.Play.Class("cui-console-play-paused"), lucide.Square.Class("cui-console-play-playing")),
		g.If(labelled, g.Group{
			h.Span(h.Class("cui-console-play-paused"), g.Text(p.Label)),
			h.Span(h.Class("cui-console-play-playing"), g.Text(p.PlayingLabel)),
		}),
	).Render(w)
}
