package console

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleBanner is cui-console-banner (components/console/console-banner): App
// Store Connect's newer look's banner, a message on a rounded panel tinted
// in its color: its icon, its title, what the page gives it under the title
// (a <p class="cui-console-banner-body">, a caption, its actions), and a
// button that closes it, which the page's script wires. An error banner is
// role="alert", any other role="status".
type ConsoleBanner struct {
	// Color is its color: the neutral, info, success, warning or error. The
	// brand has no banner: it panics.
	Color ConsoleColor
	// Icon is its icon at its start, a package lucide icon; the zero
	// icon.Icon for none.
	Icon icon.Icon
	// Title is its title; "" for none.
	Title string
	// Dismiss is its close button's label (aria-label); "" for no close
	// button.
	Dismiss string
	// Class is the page's utilities on the banner; "" for none.
	Class string
	// Children is what goes under its title; nil for nothing.
	Children g.Node
}

// Render writes the banner.
func (b ConsoleBanner) Render(w io.Writer) error {
	if b.Color == ConsoleColorBrand {
		panic("console.ConsoleBanner has no brand color")
	}
	// an error interrupts: a screen reader announces it at once
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
		g.If(b.Dismiss != "", h.Button(h.Class("cui-console-btn cui-console-btn-icon cui-console-banner-dismiss"), h.Data("variant", "plain"), h.Data("size", "sm"), h.Type("button"), h.Aria("label", b.Dismiss), lucide.X)),
	).Render(w)
}
