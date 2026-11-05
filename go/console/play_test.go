package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestPlayRendersTheExamples(t *testing.T) {
	example.Check(t, h.Div(h.Class("flex flex-wrap items-center gap-3"),
		console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: ""},
		console.ConsolePlay{State: console.ConsolePlayStatePlaying, AriaLabel: "Stop", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: ""},
	), "console/console-play/examples/play.html")
	example.Check(t, h.Div(h.Class("flex flex-wrap items-center gap-3"),
		console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "", Label: "Play This Line", PlayingLabel: "Stop", Size: console.ConsoleBtnSizeSmall, Class: ""},
		console.ConsolePlay{State: console.ConsolePlayStatePlaying, AriaLabel: "", Label: "Play This Line", PlayingLabel: "Stop", Size: console.ConsoleBtnSizeSmall, Class: ""},
	), "console/console-play/examples/labelled.html")
	example.Check(t, h.Div(h.Class("flex flex-wrap items-center gap-3"),
		console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeSmall, Class: ""},
		console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: ""},
		console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeLarge, Class: ""},
		console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeExtraLarge, Class: ""},
	), "console/console-play/examples/sizes.html")
}

func TestPlayPanics(t *testing.T) {
	for _, c := range []struct {
		want                           string
		ariaLabel, label, playingLabel string
	}{
		{"needs an AriaLabel without a Label", "", "", ""},
		{"and none with one", "Play", "Play This Line", "Stop"},
		{"needs a Label and a PlayingLabel, or neither", "", "Play This Line", ""},
		{"needs a Label and a PlayingLabel, or neither", "", "", "Stop"},
	} {
		mustPanic(t, c.want, func() {
			example.Render(console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: c.ariaLabel, Label: c.label, PlayingLabel: c.playingLabel, Size: console.ConsoleBtnSizeMedium, Class: ""})
		})
	}
}
