package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type BtnVariant struct{ value string }

var (
	BtnVariantPrimary          = BtnVariant{value: ""}
	BtnVariantSecondary        = BtnVariant{"secondary"}
	BtnVariantSecondaryNeutral = BtnVariant{"secondary-neutral"}
)

type BtnSize struct{ value string }

var (
	BtnSizeMedium = BtnSize{value: ""}
	BtnSizeLarge  = BtnSize{"lg"}
)

// Btn is apple.com's pill: a link with an Href, otherwise a <button>.
type Btn struct {
	// Href: "" for a <button>.
	Href  string
	Label string
	// Icon: before the Label, the page's node — Google's G, a spinner,
	// [lucide.Check]; nil for none. Not [lucide.Icon], the console button's
	// type: a site's button is the store's, whose marks are its page's own.
	Icon g.Node
	// Submit: a <button> that sends its form; otherwise type="button".
	Submit  bool
	Variant BtnVariant
	Size    BtnSize
	// Block fills the row.
	Block bool
	// Attrs: disabled, aria-*, an htmx attribute; nil for none.
	Attrs g.Node
	Class string
}

func (b Btn) Render(w io.Writer) error {
	if b.Label == "" {
		panic("site.Btn needs a Label")
	}
	if b.Href != "" && b.Submit {
		panic("site.Btn has an Href and Submit: a link sends no form")
	}
	root := "cui-btn"
	if b.Block {
		root += " cui-btn-block"
	}
	attrs := g.Group{class.Root(root, b.Class), g.If(b.Variant.value != "", h.Data("variant", b.Variant.value)), g.If(b.Size.value != "", h.Data("size", b.Size.value))}
	content := g.Group{b.Icon, g.Text(b.Label)}
	if b.Href != "" {
		return h.A(attrs, h.Href(b.Href), b.Attrs, content).Render(w)
	}
	kind := "button"
	if b.Submit {
		kind = "submit"
	}
	return h.Button(attrs, h.Type(kind), b.Attrs, content).Render(w)
}
