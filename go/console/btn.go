package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/attr"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type ConsoleBtnVariant struct{ value string }

var (
	ConsoleBtnVariantGray   = ConsoleBtnVariant{value: ""}
	ConsoleBtnVariantFilled = ConsoleBtnVariant{"filled"}
	ConsoleBtnVariantPlain  = ConsoleBtnVariant{"plain"}
)

type ConsoleBtnSize struct{ value string }

var (
	ConsoleBtnSizeSmall  = ConsoleBtnSize{"sm"}
	ConsoleBtnSizeMedium = ConsoleBtnSize{value: ""}
	ConsoleBtnSizeLarge  = ConsoleBtnSize{"lg"}
)

// ConsoleBtn is the console's button: a link with an Href, otherwise a <button>.
type ConsoleBtn struct {
	Href string
	// Label: "" for a button of its Icon alone, named by AriaLabel.
	Label     string
	AriaLabel string
	// Icon: before the Label; the zero Icon for none.
	Icon lucide.Icon
	// Submit: a <button> that sends its form; otherwise type="button".
	Submit      bool
	Variant     ConsoleBtnVariant
	Size        ConsoleBtnSize
	Destructive bool
	// Attrs: disabled, aria-*, an htmx attribute.
	Attrs g.Node
	Class string
}

func (b ConsoleBtn) Render(w io.Writer) error {
	if b.Href != "" && b.Submit {
		panic("console.ConsoleBtn has an Href and Submit: a link sends no form")
	}
	root := "cui-console-btn"
	shape := ""
	if b.Label == "" {
		if b.AriaLabel == "" || b.Icon == (lucide.Icon{}) {
			panic("console.ConsoleBtn needs a Label, or an Icon and an AriaLabel")
		}
		root += " cui-console-btn-icon"
		if b.Variant != ConsoleBtnVariantPlain {
			shape = "capsule"
		}
	} else if b.AriaLabel != "" {
		panic("console.ConsoleBtn has a Label and an AriaLabel: its words name it")
	}
	attrs := g.Group{attr.Class(root, b.Class), attr.Optional("data-variant", b.Variant.value), g.If(b.Destructive, h.Data("color", "error")), attr.Optional("data-size", b.Size.value), attr.Optional("data-shape", shape)}
	content := g.Group{g.If(b.Icon != (lucide.Icon{}), b.Icon), g.Text(b.Label)}
	rest := g.Group{attr.Optional("aria-label", b.AriaLabel), b.Attrs, content}
	if b.Href != "" {
		return h.A(attrs, h.Href(b.Href), rest).Render(w)
	}
	kind := "button"
	if b.Submit {
		kind = "submit"
	}
	return h.Button(attrs, h.Type(kind), rest).Render(w)
}
