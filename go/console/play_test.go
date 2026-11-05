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
	), "console/console-play/examples/sizes.html")
}

func TestPlayPanics(t *testing.T) {
	example.Panics(t, func() console.ConsolePlay {
		return console.ConsolePlay{State: console.ConsolePlayStatePaused, AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: ""}
	}, []example.Break[console.ConsolePlay]{
		{Want: "console.ConsolePlay needs an AriaLabel without a Label, and none with one", Change: func(p *console.ConsolePlay) { p.AriaLabel = "" }},
		{Want: "console.ConsolePlay needs an AriaLabel without a Label, and none with one", Change: func(p *console.ConsolePlay) { p.Label, p.PlayingLabel = "Play This Line", "Stop" }},
		{Want: "console.ConsolePlay needs a Label and a PlayingLabel, or neither", Change: func(p *console.ConsolePlay) { p.AriaLabel, p.Label = "", "Play This Line" }},
		{Want: "console.ConsolePlay needs a Label and a PlayingLabel, or neither", Change: func(p *console.ConsolePlay) { p.AriaLabel, p.PlayingLabel = "", "Stop" }},
	})
}
