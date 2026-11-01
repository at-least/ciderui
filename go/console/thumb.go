package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleThumb is cui-console-thumb (components/console/console-thumb): a
// screenshot shown small, a link to itself whole, its name under it, App
// Store Connect's App Previews and Screenshots'. Their row,
// <div class="cui-console-thumbs">, is the page's own.
type ConsoleThumb struct {
	// Href is the image whole, what it links to.
	Href string
	// Src is the image shown.
	Src string
	// Alt is what the image shows, its alt.
	Alt string
	// Caption is its name under it. Required: an empty one panics.
	Caption string
	// Lazy is whether it loads only near the window (loading="lazy"), for a
	// row of many.
	Lazy bool
	// Class is the page's utilities on the <figure>; "" for none.
	Class string
}

// Render writes the thumbnail.
func (t ConsoleThumb) Render(w io.Writer) error {
	if t.Caption == "" {
		panic("console.ConsoleThumb needs a Caption")
	}
	return h.Figure(class.Root("cui-console-thumb", t.Class),
		h.A(h.Href(t.Href), h.Img(h.Src(t.Src), h.Alt(t.Alt), g.If(t.Lazy, h.Loading("lazy")))),
		h.FigCaption(h.Class("cui-console-thumb-caption"), g.Text(t.Caption)),
	).Render(w)
}
