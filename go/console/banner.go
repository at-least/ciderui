package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type ConsoleBanner struct {
	Color ConsoleColor
	Icon  lucide.Icon
	Title string
	Class string
	// Children: a <p class="cui-console-banner-body">, actions.
	Children g.Node
}

func (b ConsoleBanner) Render(w io.Writer) error {
	role := "status"
	if b.Color == ConsoleColorError {
		role = "alert"
	}
	return h.Div(class.Root("cui-console-banner", b.Class), optionalAttr("data-color", b.Color.value), h.Role(role),
		iconPart(b.Icon, "cui-console-banner-icon"),
		h.Div(
			g.If(b.Title != "", h.P(h.Class("cui-console-banner-title"), g.Text(b.Title))),
			b.Children,
		),
	).Render(w)
}
