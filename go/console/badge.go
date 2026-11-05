package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleColor is data-color, the toolkit's color roles.
type ConsoleColor struct{ value string }

var (
	ConsoleColorNeutral = ConsoleColor{value: ""}
	ConsoleColorBrand   = ConsoleColor{"brand"}
	ConsoleColorInfo    = ConsoleColor{"info"}
	ConsoleColorSuccess = ConsoleColor{"success"}
	ConsoleColorWarning = ConsoleColor{"warning"}
	ConsoleColorError   = ConsoleColor{"error"}
)

type ConsoleBadgeVariant struct{ value string }

var (
	ConsoleBadgeVariantTinted = ConsoleBadgeVariant{value: ""}
	ConsoleBadgeVariantGray   = ConsoleBadgeVariant{"gray"}
	ConsoleBadgeVariantPlain  = ConsoleBadgeVariant{"plain"}
)

// ConsoleBadge: a colored mark before a word is [ConsoleStatus], not a badge
// (decision 27).
type ConsoleBadge struct {
	Label   string
	Color   ConsoleColor
	Variant ConsoleBadgeVariant
	Small   bool
	Class   string
}

func (b ConsoleBadge) Render(w io.Writer) error {
	if b.Label == "" {
		panic("console.ConsoleBadge needs a Label")
	}
	return h.Span(class.Root("cui-console-badge", b.Class),
		optionalAttr("data-color", b.Color.value),
		optionalAttr("data-variant", b.Variant.value),
		g.If(b.Small, h.Data("size", "sm")),
		g.Text(b.Label),
	).Render(w)
}
