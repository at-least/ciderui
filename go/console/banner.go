package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/attr"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// @invariant: one glyph per color, so a reader learns it once: SF Symbols'
// info.circle, checkmark.circle, exclamationmark.triangle and
// exclamationmark.circle in Lucide
var bannerIcons = map[ConsoleColor]lucide.Icon{
	ConsoleColorNeutral: lucide.Info,
	ConsoleColorInfo:    lucide.Info,
	ConsoleColorSuccess: lucide.CircleCheck,
	ConsoleColorWarning: lucide.TriangleAlert,
	ConsoleColorError:   lucide.CircleAlert,
}

// ConsoleBanner is a message across the page, its icon its Color's.
type ConsoleBanner struct {
	Color ConsoleColor
	Title string
	// Body: the banner's words, or blocks of them, written in its
	// <div class="cui-console-banner-body">.
	Body g.Node
	// Children: what follows the body, its actions.
	Children g.Node
}

func (b ConsoleBanner) Render(w io.Writer) error {
	role := "status"
	if b.Color == ConsoleColorError {
		role = "alert"
	}
	return h.Div(h.Class("cui-console-banner"), attr.Optional("data-color", b.Color.value), h.Role(role),
		bannerIcons[b.Color].Class("cui-console-banner-icon"),
		h.Div(
			optionalText(h.P, "cui-console-banner-title", b.Title),
			g.If(b.Body != nil, h.Div(h.Class("cui-console-banner-body"), b.Body)),
			b.Children,
		),
	).Render(w)
}
