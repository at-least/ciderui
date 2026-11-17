package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func artwork() console.ConsoleArtwork {
	return console.ConsoleArtwork{Src: exampleSrc("console/console-artwork/examples/artwork.html", "The feast"), Label: "The feast"}
}

func TestArtworkRendersTheExample(t *testing.T) {
	example.Check(t, h.Div(h.Class("grid w-80 grid-cols-3 gap-2"),
		artwork(),
		console.ConsoleArtwork{Src: "", Label: "The Little Red Hen"},
	), "console/console-artwork/examples/artwork.html")
}

func TestArtworkPanics(t *testing.T) {
	example.Panics(t, artwork, []example.Break[console.ConsoleArtwork]{
		{Want: "console.ConsoleArtwork needs a Label", Change: func(p *console.ConsoleArtwork) { p.Label = "" }},
		{Want: "console.ConsoleArtwork needs a Label", Change: func(p *console.ConsoleArtwork) { p.Src, p.Label = "", "" }},
	})
}
