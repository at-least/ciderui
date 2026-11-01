package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleDisclosureVariant is a [ConsoleDisclosure]'s look (data-variant). Its zero value
// is [ConsoleDisclosureVariantLine].
type ConsoleDisclosureVariant struct{ value string }

var (
	// ConsoleDisclosureVariantLine is a line, the chevron before its label, the
	// default: no data-variant.
	ConsoleDisclosureVariantLine = ConsoleDisclosureVariant{value: ""}
	// ConsoleDisclosureVariantButton is a console button, the chevron after its
	// label.
	ConsoleDisclosureVariantButton = ConsoleDisclosureVariant{"button"}
	// ConsoleDisclosureVariantPlainButton is a plain console button
	// (data-variant="plain": no fill, the brand's label), the chevron after
	// its label.
	ConsoleDisclosureVariantPlainButton = ConsoleDisclosureVariant{"plain-button"}
)

// ConsoleDisclosure is cui-console-disclosure (components/console/console-disclosure):
// a <details> hiding what is not needed yet (Children) behind a line or a
// console button, the HIG's disclosure control.
type ConsoleDisclosure struct {
	// Variant is its look.
	Variant ConsoleDisclosureVariant
	// Open is whether it starts open.
	Open bool
	// Summary is its label. Required: an empty one panics.
	Summary string
	// SummaryClass is the page's utilities on its <summary>; "" for none.
	SummaryClass string
	// Class is the page's utilities on its <details>; "" for none.
	Class string
	// Children is what it hides.
	Children g.Node
}

// Render writes the disclosure.
func (d ConsoleDisclosure) Render(w io.Writer) error {
	if d.Summary == "" {
		panic("console.ConsoleDisclosure needs a Summary")
	}
	const chevron = "cui-console-disclosure-chevron"
	if d.Variant == ConsoleDisclosureVariantLine {
		return h.Details(class.Root("cui-console-disclosure", d.Class), g.If(d.Open, h.Open()),
			h.Summary(class.Root("cui-console-disclosure-summary", d.SummaryClass), lucide.ChevronRight.Class(chevron), g.Text(d.Summary)),
			d.Children,
		).Render(w)
	}
	return h.Details(class.Root("cui-console-disclosure", d.Class), h.Data("variant", "button"), g.If(d.Open, h.Open()),
		h.Summary(class.Root("cui-console-btn", d.SummaryClass), g.If(d.Variant == ConsoleDisclosureVariantPlainButton, h.Data("variant", "plain")), g.Text(d.Summary), lucide.ChevronDown.Class(chevron)),
		d.Children,
	).Render(w)
}
