package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleBtnSize is cui-console-btn's data-size.
type ConsoleBtnSize struct{ value string }

var (
	ConsoleBtnSizeSmall  = ConsoleBtnSize{"sm"}
	ConsoleBtnSizeMedium = ConsoleBtnSize{value: ""}
	ConsoleBtnSizeLarge  = ConsoleBtnSize{"lg"}
)

// ConsolePlayState is what the clip is doing; the page's script sets
// data-state as it plays and stops.
type ConsolePlayState struct{ value string }

var (
	// ConsolePlayStatePaused shows the play glyph and the Label.
	ConsolePlayStatePaused = ConsolePlayState{value: ""}
	// ConsolePlayStatePlaying shows the stop glyph in a ring of how much has
	// played, and the PlayingLabel.
	ConsolePlayStatePlaying = ConsolePlayState{"playing"}
)

// ConsolePlay is a button that plays one audio clip and stops it. It holds
// both glyphs and both labels; the page's script switches data-state (and
// an icon-only one's AriaLabel) and, while playing, sets
// --cui-console-play-progress (0% to 100%), the ring round the glyph.
type ConsolePlay struct {
	State ConsolePlayState
	// AriaLabel is an icon-only button's name, what pressing it does now
	// ("Play" while paused, "Stop" while playing); "" with a Label, whose
	// words name it.
	AriaLabel string
	// Label is the words after the glyph while paused, PlayingLabel while
	// playing; "" both for a round button of the glyph alone.
	Label        string
	PlayingLabel string
	Size         ConsoleBtnSize
	Class        string
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
		optionalAttr("data-size", p.Size.value), h.Data("shape", "capsule"), optionalAttr("data-state", p.State.value),
		h.Type("button"), optionalAttr("aria-label", p.AriaLabel),
		h.Span(h.Class("cui-console-play-glyph"), lucide.Play.Class("cui-console-play-paused"), lucide.Square.Class("cui-console-play-playing")),
		g.If(labelled, g.Group{
			h.Span(h.Class("cui-console-play-paused"), g.Text(p.Label)),
			h.Span(h.Class("cui-console-play-playing"), g.Text(p.PlayingLabel)),
		}),
	).Render(w)
}
