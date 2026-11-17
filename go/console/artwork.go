package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/need"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleArtwork is a picture's thumbnail, 16:9 and as wide as the page's room
// for it, or its Label's words in its place where there is no picture.
type ConsoleArtwork struct {
	Src string
	// Label: with a Src, the picture's alt text; without one, the words shown.
	Label string
}

func (a ConsoleArtwork) Render(w io.Writer) error {
	need.Field("console.ConsoleArtwork", "a Label", a.Label)
	if a.Src == "" {
		return h.Span(h.Class("cui-console-artwork"), g.Text(a.Label)).Render(w)
	}
	return h.Img(h.Class("cui-console-artwork"), h.Src(a.Src), h.Alt(a.Label), h.Loading("lazy")).Render(w)
}
