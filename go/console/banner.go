package console

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleBanner: the page's script wires the close button.
type ConsoleBanner struct {
	Color ConsoleColor
	Icon  icon.Icon
	Title string
	// DismissAriaLabel: "" for no close button.
	DismissAriaLabel string
	Class            string
	// Children: a <p class="cui-console-banner-body">, a caption, actions.
	Children g.Node
}

func (b ConsoleBanner) Render(w io.Writer) error {
	if b.Color == ConsoleColorBrand {
		panic("console.ConsoleBanner has no brand color")
	}
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
		g.If(b.DismissAriaLabel != "", h.Button(h.Class("cui-console-btn cui-console-btn-icon cui-console-banner-dismiss"), h.Data("variant", "plain"), h.Data("size", "sm"), h.Type("button"), h.Aria("label", b.DismissAriaLabel), lucide.X)),
	).Render(w)
}
