package docs

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DocsAsideVariant is DocC's six aside kinds.
type DocsAsideVariant struct{ value string }

var (
	DocsAsideVariantNote       = DocsAsideVariant{value: ""}
	DocsAsideVariantImportant  = DocsAsideVariant{"important"}
	DocsAsideVariantWarning    = DocsAsideVariant{"warning"}
	DocsAsideVariantTip        = DocsAsideVariant{"tip"}
	DocsAsideVariantDeprecated = DocsAsideVariant{"deprecated"}
	DocsAsideVariantExperiment = DocsAsideVariant{"experiment"}
)

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

type DocsAside struct {
	Variant DocsAsideVariant
	Class   string
	// Children: the page's paragraphs.
	Children g.Node
}

func (a DocsAside) Render(w io.Writer) error {
	label := a.Variant.label()
	return h.Aside(class.Root("cui-docs-aside", a.Class),
		g.If(a.Variant.value != "", h.Data("variant", a.Variant.value)),
		h.Aria("label", label),
		h.P(h.Class("cui-docs-aside-label"), g.Text(label)),
		a.Children,
	).Render(w)
}
