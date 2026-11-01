package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleColor is a console part's color (data-color): the toolkit's color roles.
// Its zero value is [ConsoleColorNeutral], which writes no data-color.
type ConsoleColor struct{ value string }

var (
	// ConsoleColorNeutral is the neutral, the default: no data-color.
	ConsoleColorNeutral = ConsoleColor{value: ""}
	// ConsoleColorBrand is the brand's blue.
	ConsoleColorBrand = ConsoleColor{"brand"}
	// ConsoleColorInfo is information's blue.
	ConsoleColorInfo = ConsoleColor{"info"}
	// ConsoleColorSuccess is success's green.
	ConsoleColorSuccess = ConsoleColor{"success"}
	// ConsoleColorWarning is a warning's yellow.
	ConsoleColorWarning = ConsoleColor{"warning"}
	// ConsoleColorError is an error's red.
	ConsoleColorError = ConsoleColor{"error"}
)

// Attribute is its data-color value; "" for the neutral, which writes none.
func (c ConsoleColor) Attribute() string {
	return c.value
}

// ConsoleBadgeVariant is a [ConsoleBadge]'s ground (data-variant). Its zero value is
// [ConsoleBadgeVariantTinted].
type ConsoleBadgeVariant struct{ value string }

var (
	// ConsoleBadgeVariantTinted is its color at 0.15, the default: no data-variant.
	ConsoleBadgeVariantTinted = ConsoleBadgeVariant{value: ""}
	// ConsoleBadgeVariantGray is the gray fill under its colored label.
	ConsoleBadgeVariantGray = ConsoleBadgeVariant{"gray"}
	// ConsoleBadgeVariantPlain is no fill, its color alone.
	ConsoleBadgeVariantPlain = ConsoleBadgeVariant{"plain"}
)

// ConsoleBadge is cui-console-badge (components/console/console-badge): App Store
// Connect's newer look's badge, a short word on a pill tinted in its color,
// on the gray, or plain; with Dot, a dot before it.
type ConsoleBadge struct {
	// Label is its word. Required: an empty one panics.
	Label string
	// Color is its color.
	Color ConsoleColor
	// Variant is its ground.
	Variant ConsoleBadgeVariant
	// Small is whether it is small (data-size="sm").
	Small bool
	// Dot is whether a dot in its color goes before its word.
	Dot bool
	// Class is the page's utilities on the badge; "" for none.
	Class string
}

// Render writes the badge.
func (b ConsoleBadge) Render(w io.Writer) error {
	if b.Label == "" {
		panic("console.ConsoleBadge needs a Label")
	}
	return h.Span(class.Root("cui-console-badge", b.Class),
		optionalAttr("data-color", b.Color.value),
		optionalAttr("data-variant", b.Variant.value),
		g.If(b.Small, h.Data("size", "sm")),
		g.If(b.Dot, h.Span(h.Class("cui-console-badge-dot"))),
		g.Text(b.Label),
	).Render(w)
}
