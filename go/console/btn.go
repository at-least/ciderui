package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/attr"
	"github.com/at-least/ciderui/go/internal/need"
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
	Href  string
	Label string
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
	need.Field("console.ConsoleBtn", "a Label", b.Label)
	attrs := g.Group{attr.Class("cui-console-btn", b.Class), attr.Optional("data-variant", b.Variant.value), g.If(b.Destructive, h.Data("color", "error")), attr.Optional("data-size", b.Size.value)}
	rest := g.Group{b.Attrs, g.Text(b.Label)}
	if b.Href != "" {
		return h.A(attrs, h.Href(b.Href), rest).Render(w)
	}
	kind := "button"
	if b.Submit {
		kind = "submit"
	}
	return h.Button(attrs, h.Type(kind), rest).Render(w)
}

// @invariant: iconButton (the sidebar's open and close) and viewerCloseButton are
// the kit's buttons of an icon alone, here because the gate holds every button
// class to this file: a console button a page writes is words alone, and a
// page's own icon button is these classes on its <button> (examples/icon.html,
// capsule.html).
func iconButton(icon lucide.Icon, ariaLabel string, attrs g.Node, class string) g.Node {
	return h.Button(attr.Class("cui-console-btn cui-console-btn-icon", class), h.Data("variant", "plain"), h.Type("button"), h.Aria("label", ariaLabel), attrs, icon)
}

func viewerCloseButton(ariaLabel string) g.Node {
	return h.Button(h.Class("cui-console-btn cui-console-btn-icon"), h.Data("size", "sm"), h.Data("shape", "capsule"), h.Type("submit"), h.Aria("label", ariaLabel), lucide.X)
}
