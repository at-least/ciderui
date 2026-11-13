package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func viewer() console.ConsoleViewer {
	return console.ConsoleViewer{ID: "cover", AriaLabel: "Cover", CloseAriaLabel: "Close"}
}

func TestViewerRendersTheExample(t *testing.T) {
	example.Check(t, g.Group{
		h.Button(h.Class("cui-console-btn"), h.Type("button"), g.Attr("commandfor", "cover"), g.Attr("command", "show-modal"), g.Text("View Cover")),
		viewer(),
	}, "console/console-viewer/examples/viewer.html")
}

func TestViewerPanics(t *testing.T) {
	example.Panics(t, viewer, []example.Break[console.ConsoleViewer]{
		{Want: "console.ConsoleViewer needs an ID", Change: func(p *console.ConsoleViewer) { p.ID = "" }},
		{Want: "console.ConsoleViewer needs an AriaLabel", Change: func(p *console.ConsoleViewer) { p.AriaLabel = "" }},
		{Want: "console.ConsoleViewer needs a CloseAriaLabel", Change: func(p *console.ConsoleViewer) { p.CloseAriaLabel = "" }},
	})
}
