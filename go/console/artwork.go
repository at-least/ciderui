package console

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleArtwork is a picture's thumbnail, 16:9 and as wide as the page's room
// for it, or its Placeholder's words in its place where there is no picture.
type ConsoleArtwork struct {
	Src         string
	Alt         string
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
		return h.Span(h.Class("cui-console-artwork"), g.Text(a.Placeholder)).Render(w)
	}
	if a.Placeholder != "" {
		panic("console.ConsoleArtwork has a Src and a Placeholder: the Placeholder is for no picture")
	}
	return h.Img(h.Class("cui-console-artwork"), h.Src(a.Src), h.Alt(a.Alt), h.Loading("lazy")).Render(w)
}
