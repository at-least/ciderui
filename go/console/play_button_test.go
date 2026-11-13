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
		rest                console.ConsolePlayButton
	}{
		{"console/console-play-button/examples/play.html", "Stop", console.ConsolePlayButton{Attrs: nil, AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: ""}},
		{"console/console-play-button/examples/labelled.html", "", console.ConsolePlayButton{Attrs: nil, AriaLabel: "", Label: "Play This Line", PlayingLabel: "Stop", Size: console.ConsoleBtnSizeSmall, Class: ""}},
	} {
		example.Check(t, h.Div(h.Class("flex flex-wrap items-center gap-3"), c.rest, g.Raw(playing(example.Render(c.rest), c.stopAriaLabel))), c.path)
	}
	example.Check(t, h.Div(h.Class("flex flex-wrap items-center gap-3"),
		console.ConsolePlayButton{Attrs: nil, AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeSmall, Class: ""},
		console.ConsolePlayButton{Attrs: nil, AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: ""},
		console.ConsolePlayButton{Attrs: nil, AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeLarge, Class: ""},
	), "console/console-play-button/examples/sizes.html")
}

func TestPlayPanics(t *testing.T) {
	example.Panics(t, func() console.ConsolePlayButton {
		return console.ConsolePlayButton{Attrs: nil, AriaLabel: "Play", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: ""}
	}, []example.Break[console.ConsolePlayButton]{
		{Want: "console.ConsolePlayButton needs an AriaLabel without a Label, and none with one", Change: func(p *console.ConsolePlayButton) { p.AriaLabel = "" }},
		{Want: "console.ConsolePlayButton needs an AriaLabel without a Label, and none with one", Change: func(p *console.ConsolePlayButton) { p.Label, p.PlayingLabel = "Play This Line", "Stop" }},
		{Want: "console.ConsolePlayButton needs a Label and a PlayingLabel, or neither", Change: func(p *console.ConsolePlayButton) { p.AriaLabel, p.Label = "", "Play This Line" }},
		{Want: "console.ConsolePlayButton needs a Label and a PlayingLabel, or neither", Change: func(p *console.ConsolePlayButton) { p.AriaLabel, p.PlayingLabel = "", "Stop" }},
	})
}
