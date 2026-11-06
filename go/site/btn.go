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
	BtnVariantNeutral          = BtnVariant{"neutral"}
	BtnVariantSecondaryNeutral = BtnVariant{"secondary-neutral"}
)

type BtnSize struct{ value string }

var (
	BtnSizeSmall  = BtnSize{"sm"}
	BtnSizeMedium = BtnSize{value: ""}
	BtnSizeLarge  = BtnSize{"lg"}
	BtnSizeXLarge = BtnSize{"xl"}
	// BtnSizeCTA is a hero's call to action, stepping at 735px.
	BtnSizeCTA = BtnSize{"cta"}
)

// Btn is apple.com's pill: a link with an Href, otherwise a <button>.
type Btn struct {
	// Href: "" for a <button>.
	Href  string
	Label string
	// Submit: a <button> that sends its form; otherwise type="button".
	Submit  bool
	Variant BtnVariant
	Size    BtnSize
	// Block fills the row; Rect is the store's form button's corners.
	Block bool
	Rect  bool
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
	if b.Rect {
		root += " cui-btn-rect"
	}
	attrs := g.Group{class.Root(root, b.Class), g.If(b.Variant.value != "", h.Data("variant", b.Variant.value)), g.If(b.Size.value != "", h.Data("size", b.Size.value))}
	if b.Href != "" {
		return h.A(attrs, h.Href(b.Href), b.Attrs, g.Text(b.Label)).Render(w)
	}
	kind := "button"
	if b.Submit {
		kind = "submit"
	}
	return h.Button(attrs, h.Type(kind), b.Attrs, g.Text(b.Label)).Render(w)
}
