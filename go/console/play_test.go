package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestPlayRendersTheExamples(t *testing.T) {
	example.Check(t, h.Div(h.Class("flex flex-wrap items-center gap-3"),
		console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "Play", Size: console.ConsoleBtnSizeMedium, Class: ""},
		console.ConsolePlay{State: console.ConsolePlayStatePlaying, AriaLabel: "Pause", Size: console.ConsoleBtnSizeMedium, Class: ""},
	), "console/console-play/examples/play.html")
	example.Check(t, h.Div(h.Class("flex flex-wrap items-center gap-3"),
		console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "Play", Size: console.ConsoleBtnSizeSmall, Class: ""},
		console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "Play", Size: console.ConsoleBtnSizeMedium, Class: ""},
		console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "Play", Size: console.ConsoleBtnSizeLarge, Class: ""},
		console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "Play", Size: console.ConsoleBtnSizeExtraLarge, Class: ""},
	), "console/console-play/examples/sizes.html")
}

func TestPlayAnEmptyAriaLabelPanics(t *testing.T) {
	mustPanic(t, "ConsolePlay needs an AriaLabel", func() {
		example.Render(console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "", Size: console.ConsoleBtnSizeMedium, Class: ""})
	})
}
