package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/need"
	h "maragu.dev/gomponents/html"
)

// ConsoleViewerID is the viewer's id, for a page's commandfor and script.
//
// @invariant: one viewer shows each of a page's pictures, so its id is the
// kit's, not the page's
const ConsoleViewerID = "viewer"

// ConsoleViewer is a picture shown whole over the page, fit to the window
// (Quick Look), its label under it and a close button in a row above it. The
// page opens it as a <dialog> (showModal() or a button's command="show-modal");
// its script sets the picture's src and alt and the label's words before it
// opens it, so one viewer shows each of many pictures.
type ConsoleViewer struct {
	AriaLabel      string
	CloseAriaLabel string
}

func (v ConsoleViewer) Render(w io.Writer) error {
	need.Field("console.ConsoleViewer", "an AriaLabel", v.AriaLabel)
	need.Field("console.ConsoleViewer", "a CloseAriaLabel", v.CloseAriaLabel)
	return h.Dialog(h.Class("cui-console-viewer"), h.ID(ConsoleViewerID), h.Aria("label", v.AriaLabel),
		h.Form(h.Method("dialog"), h.Class("cui-console-viewer-close"),
			viewerCloseButton(v.CloseAriaLabel),
		),
		h.Img(h.Class("cui-console-viewer-image"), h.Alt("")),
		h.P(h.Class("cui-console-viewer-label")),
	).Render(w)
}
