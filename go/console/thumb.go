package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleThumb is a screenshot shown small, linking to the full image (App
// Store Connect's App Previews and Screenshots). Their row,
// <div class="cui-console-thumbs">, is the page's own.
type ConsoleThumb struct {
	// Href is the full image.
	Href string
	Src  string
	Alt  string
	// Caption is the file's name.
	Caption string
	// Lazy: for a row of many.
	Lazy  bool
	Class string
}

func (t ConsoleThumb) Render(w io.Writer) error {
	if t.Caption == "" {
		panic("console.ConsoleThumb needs a Caption")
	}
	return h.Figure(class.Root("cui-console-thumb", t.Class),
		h.A(h.Href(t.Href), h.Img(h.Src(t.Src), h.Alt(t.Alt), g.If(t.Lazy, h.Loading("lazy")))),
		h.FigCaption(h.Class("cui-console-thumb-caption"), g.Text(t.Caption)),
	).Render(w)
}
