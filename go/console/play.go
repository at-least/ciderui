package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	h "maragu.dev/gomponents/html"
)

// ConsoleBtnSize is cui-console-btn's data-size.
type ConsoleBtnSize struct{ value string }

var (
	ConsoleBtnSizeSmall      = ConsoleBtnSize{"sm"}
	ConsoleBtnSizeMedium     = ConsoleBtnSize{value: ""}
	ConsoleBtnSizeLarge      = ConsoleBtnSize{"lg"}
	ConsoleBtnSizeExtraLarge = ConsoleBtnSize{"xl"}
)

// ConsolePlayState is what the clip is doing; the page's script sets
// data-state as it plays and stops.
type ConsolePlayState struct{ value string }

var (
	// ConsolePlayStatePaused shows the play glyph.
	ConsolePlayStatePaused = ConsolePlayState{value: ""}
	// ConsolePlayStatePlaying shows the pause glyph.
	ConsolePlayStatePlaying = ConsolePlayState{"playing"}
)

// ConsolePlay is a round play and pause button for one audio clip. It
// holds both glyphs; the page's script switches data-state and AriaLabel
// together, and Class is its hook.
type ConsolePlay struct {
	State ConsolePlayState
	// AriaLabel is what pressing it does now ("Play" while paused, "Pause"
	// while playing).
	AriaLabel string
	Size      ConsoleBtnSize
	Class     string
}

func (p ConsolePlay) Render(w io.Writer) error {
	if p.AriaLabel == "" {
		panic("console.ConsolePlay needs an AriaLabel")
	}
	return h.Button(class.Root("cui-console-btn cui-console-btn-icon cui-console-play", p.Class),
		optionalAttr("data-size", p.Size.value), h.Data("shape", "capsule"), optionalAttr("data-state", p.State.value),
		h.Type("button"), h.Aria("label", p.AriaLabel),
		lucide.Play.Class("cui-console-play-paused"), lucide.Pause.Class("cui-console-play-playing"),
	).Render(w)
}
