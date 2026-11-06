package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type ConsoleDisclosureVariant struct{ value string }

var (
	ConsoleDisclosureVariantLine   = ConsoleDisclosureVariant{value: ""}
	ConsoleDisclosureVariantButton = ConsoleDisclosureVariant{"button"}
	// ConsoleDisclosureVariantPlainButton is not a data-variant value of the
	// disclosure: it writes data-variant="button" on the <details> and
	// data-variant="plain" on the button.
	ConsoleDisclosureVariantPlainButton = ConsoleDisclosureVariant{"plain-button"}
)

// ConsoleDisclosure is the HIG's disclosure control.
type ConsoleDisclosure struct {
	Variant      ConsoleDisclosureVariant
	Open         bool
	Summary      string
	SummaryClass string
	Class        string
	Children     g.Node
}

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
		h.Summary(class.Root("cui-console-btn", d.SummaryClass), g.If(d.Variant == ConsoleDisclosureVariantPlainButton, h.Data("variant", ConsoleBtnVariantPlain.value)), g.Text(d.Summary), lucide.ChevronDown.Class(chevron)),
		d.Children,
	).Render(w)
}
