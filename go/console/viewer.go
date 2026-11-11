package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/attr"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleViewer is a picture shown whole over the page, fit to the window
// (Quick Look), its caption under it and a close button over its corner. The
// page opens it as a [ConsoleDialog]; a page whose one viewer shows each of
// many pictures writes Src, Alt and Caption "" and its script sets them
// before it opens it.
type ConsoleViewer struct {
	ID             string
	AriaLabel      string
	CloseAriaLabel string
	Src            string
	Alt            string
	Caption        string
}

func (v ConsoleViewer) Render(w io.Writer) error {
	need("console.ConsoleViewer", "an ID", v.ID)
	need("console.ConsoleViewer", "an AriaLabel", v.AriaLabel)
	need("console.ConsoleViewer", "a CloseAriaLabel", v.CloseAriaLabel)
	if v.Src == "" && (v.Alt != "" || v.Caption != "") {
		panic("console.ConsoleViewer has an Alt or a Caption and no Src: the page's script sets all three")
	}
	return h.Dialog(h.Class("cui-console-viewer"), h.ID(v.ID), h.Aria("label", v.AriaLabel),
		h.Img(h.Class("cui-console-viewer-image"), attr.Optional("src", v.Src), h.Alt(v.Alt)),
		h.P(h.Class("cui-console-viewer-caption"), g.Text(v.Caption)),
		h.Form(h.Method("dialog"), h.Class("cui-console-viewer-close"),
			ConsoleBtn{Href: "", Label: "", AriaLabel: v.CloseAriaLabel, Icon: lucide.X, Submit: true, Variant: ConsoleBtnVariantGray, Size: ConsoleBtnSizeSmall, Destructive: false, Attrs: nil, Class: ""},
		),
	).Render(w)
}
