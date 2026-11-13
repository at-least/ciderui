package console

import (
	"io"

	"github.com/at-least/ciderui/go/lucide"
	h "maragu.dev/gomponents/html"
)

// ConsoleViewer is a picture shown whole over the page, fit to the window
// (Quick Look), its caption under it and a close button in a row above it. The
// page opens it as a [ConsoleDialog]; its script sets the picture's src and alt
// and the caption's words before it opens it, so one viewer shows each of many
// pictures.
type ConsoleViewer struct {
	ID             string
	AriaLabel      string
	CloseAriaLabel string
}

func (v ConsoleViewer) Render(w io.Writer) error {
	need("console.ConsoleViewer", "an ID", v.ID)
	need("console.ConsoleViewer", "an AriaLabel", v.AriaLabel)
	need("console.ConsoleViewer", "a CloseAriaLabel", v.CloseAriaLabel)
	return h.Dialog(h.Class("cui-console-viewer"), h.ID(v.ID), h.Aria("label", v.AriaLabel),
		h.Form(h.Method("dialog"), h.Class("cui-console-viewer-close"),
			ConsoleBtn{Href: "", Label: "", AriaLabel: v.CloseAriaLabel, Icon: lucide.X, Submit: true, Variant: ConsoleBtnVariantGray, Size: ConsoleBtnSizeSmall, Destructive: false, Attrs: nil, Class: ""},
		),
		h.Img(h.Class("cui-console-viewer-image"), h.Alt("")),
		h.P(h.Class("cui-console-viewer-caption")),
	).Render(w)
}
