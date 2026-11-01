package docs

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DocsAsideVariant is a [DocsAside]'s kind: DocC's six, each its ground and
// its edge's color, and the word its label says. The zero value is a note,
// the gray one.
type DocsAsideVariant struct{ value string }

var (
	// DocsAsideVariantImportant is something the reader must not miss.
	DocsAsideVariantImportant = DocsAsideVariant{"important"}
	// DocsAsideVariantWarning is a danger: data lost, a build broken.
	DocsAsideVariantWarning = DocsAsideVariant{"warning"}
	// DocsAsideVariantTip is a way to do it better.
	DocsAsideVariantTip = DocsAsideVariant{"tip"}
	// DocsAsideVariantDeprecated is what not to use any more.
	DocsAsideVariantDeprecated = DocsAsideVariant{"deprecated"}
	// DocsAsideVariantExperiment is something to try.
	DocsAsideVariantExperiment = DocsAsideVariant{"experiment"}
)

// the word its label says, and the aside's accessible name
func (v DocsAsideVariant) label() string {
	switch v.value {
	case "":
		return "Note"
	case "important":
		return "Important"
	case "warning":
		return "Warning"
	case "tip":
		return "Tip"
	case "deprecated":
		return "Deprecated"
	case "experiment":
		return "Experiment"
	}
	panic("docs.DocsAsideVariant: no such variant: " + v.value)
}

// DocsAside is cui-docs-aside (components/docs/docs-aside): Apple's developer
// documentation's aside, a note set apart from an article's text, its kind's
// word over what the page gives it, in its kind's colors.
type DocsAside struct {
	// Variant is its kind; the zero value, a note, writes no data-variant.
	Variant DocsAsideVariant
	// Class is the page's utilities on the aside; "" for none.
	Class string
	// Children is what it says, under its label: the page's paragraphs.
	Children g.Node
}

// Render writes the aside.
func (a DocsAside) Render(w io.Writer) error {
	label := a.Variant.label()
	return h.Aside(class.Root("cui-docs-aside", a.Class),
		g.If(a.Variant.value != "", h.Data("variant", a.Variant.value)),
		h.Aria("label", label),
		h.P(h.Class("cui-docs-aside-label"), g.Text(label)),
		a.Children,
	).Render(w)
}
