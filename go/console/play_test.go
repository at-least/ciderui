package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// @invariant: a button at rest as the page's script sets it playing: data-state,
// and an icon-only one named for what a press does now
func playing(rest, ariaLabel string) string {
	html := strings.Replace(rest, `data-shape="capsule"`, `data-shape="capsule" data-state="playing"`, 1)
	if ariaLabel != "" {
		html = strings.Replace(html, `aria-label="Play"`, `aria-label="`+ariaLabel+`"`, 1)
	}
	return html
}

// @invariant: each example's second button is its first as the page's script sets it playing
func TestPlayRendersTheExamples(t *testing.T) {
	for _, c := range []struct {
		path, stopAriaLabel string
		rest                console.ConsolePlay
	}{
		{"console/console-play/examples/play.html", "Stop", console.ConsolePlay{AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: ""}},
		{"console/console-play/examples/labelled.html", "", console.ConsolePlay{AriaLabel: "", Label: "Play This Line", PlayingLabel: "Stop", Size: console.ConsoleBtnSizeSmall, Class: ""}},
	} {
		example.Check(t, h.Div(h.Class("flex flex-wrap items-center gap-3"), c.rest, g.Raw(playing(example.Render(c.rest), c.stopAriaLabel))), c.path)
	}
	example.Check(t, h.Div(h.Class("flex flex-wrap items-center gap-3"),
		console.ConsolePlay{AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeSmall, Class: ""},
		console.ConsolePlay{AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: ""},
		console.ConsolePlay{AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeLarge, Class: ""},
	), "console/console-play/examples/sizes.html")
}

func TestPlayPanics(t *testing.T) {
	example.Panics(t, func() console.ConsolePlay {
		return console.ConsolePlay{AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: ""}
	}, []example.Break[console.ConsolePlay]{
		{Want: "console.ConsolePlay needs an AriaLabel without a Label, and none with one", Change: func(p *console.ConsolePlay) { p.AriaLabel = "" }},
		{Want: "console.ConsolePlay needs an AriaLabel without a Label, and none with one", Change: func(p *console.ConsolePlay) { p.Label, p.PlayingLabel = "Play This Line", "Stop" }},
		{Want: "console.ConsolePlay needs a Label and a PlayingLabel, or neither", Change: func(p *console.ConsolePlay) { p.AriaLabel, p.Label = "", "Play This Line" }},
		{Want: "console.ConsolePlay needs a Label and a PlayingLabel, or neither", Change: func(p *console.ConsolePlay) { p.AriaLabel, p.PlayingLabel = "", "Stop" }},
	})
}
