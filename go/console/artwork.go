package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/attr"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type ConsoleArtworkAspect struct{ value string }

var (
	// ConsoleArtworkAspectWide is 16:9, a video's or a page's picture.
	ConsoleArtworkAspectWide = ConsoleArtworkAspect{""}
	// ConsoleArtworkAspectSquare is 1:1, an icon's or a cover's.
	ConsoleArtworkAspectSquare = ConsoleArtworkAspect{"square"}
)

// ConsoleArtwork is a picture's thumbnail, as wide as the page's room for it,
// or its Placeholder's words in its place where there is no picture.
type ConsoleArtwork struct {
	Src         string
	Alt         string
	Aspect      ConsoleArtworkAspect
	Placeholder string
}

func (a ConsoleArtwork) Render(w io.Writer) error {
	if a.Src == "" {
		if a.Placeholder == "" {
			panic("console.ConsoleArtwork needs a Src or a Placeholder")
		}
		if a.Alt != "" {
			panic("console.ConsoleArtwork has an Alt and no Src: its Placeholder's words are read")
		}
		return h.Span(h.Class("cui-console-artwork"), attr.Optional("data-aspect", a.Aspect.value), g.Text(a.Placeholder)).Render(w)
	}
	if a.Placeholder != "" {
		panic("console.ConsoleArtwork has a Src and a Placeholder: the Placeholder is for no picture")
	}
	return h.Img(h.Class("cui-console-artwork"), attr.Optional("data-aspect", a.Aspect.value), h.Src(a.Src), h.Alt(a.Alt), h.Loading("lazy")).Render(w)
}
