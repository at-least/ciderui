package console

import (
	"io"

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
	// Src: "" for none, its Placeholder shown.
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
			panic("console.ConsoleArtwork with no Src has no Alt: its Placeholder's words are read")
		}
		return h.Span(h.Class("cui-console-artwork"), optionalAttr("data-aspect", a.Aspect.value), g.Text(a.Placeholder)).Render(w)
	}
	if a.Placeholder != "" {
		panic("console.ConsoleArtwork with a Src has no Placeholder")
	}
	return h.Img(h.Class("cui-console-artwork"), optionalAttr("data-aspect", a.Aspect.value), h.Src(a.Src), h.Alt(a.Alt), h.Loading("lazy")).Render(w)
}
