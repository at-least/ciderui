package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func artwork() console.ConsoleArtwork {
	return console.ConsoleArtwork{Src: exampleSrc("console/console-artwork/examples/artwork.html", "The feast"), Alt: "The feast", Placeholder: ""}
}

func TestArtworkRendersTheExample(t *testing.T) {
	example.Check(t, h.Div(h.Class("grid w-80 grid-cols-3 gap-2"),
		artwork(),
		console.ConsoleArtwork{Src: "", Alt: "", Placeholder: "The Little Red Hen"},
	), "console/console-artwork/examples/artwork.html")
}

func TestArtworkPanics(t *testing.T) {
	example.Panics(t, artwork, []example.Break[console.ConsoleArtwork]{
		{Want: "console.ConsoleArtwork needs a Src or a Placeholder", Change: func(p *console.ConsoleArtwork) { p.Src, p.Alt = "", "" }},
		{Want: "console.ConsoleArtwork has an Alt and no Src: its Placeholder's words are read", Change: func(p *console.ConsoleArtwork) { p.Src, p.Placeholder = "", "x" }},
		{Want: "console.ConsoleArtwork has a Src and a Placeholder: the Placeholder is for no picture", Change: func(p *console.ConsoleArtwork) { p.Placeholder = "x" }},
	})
}
