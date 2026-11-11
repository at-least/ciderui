package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func viewer() console.ConsoleViewer {
	return console.ConsoleViewer{ID: "cover", AriaLabel: "Cover", CloseLabel: "Close", Src: exampleSrc("console/console-viewer/examples/viewer.html", "The cover"), Alt: "The cover", Caption: "Drawn 2 November"}
}

func TestViewerRendersTheExample(t *testing.T) {
	example.Check(t, g.Group{
		h.Button(h.Class("cui-console-btn"), h.Type("button"), g.Attr("commandfor", "cover"), g.Attr("command", "show-modal"), g.Text("View Cover")),
		viewer(),
	}, "console/console-viewer/examples/viewer.html")
}

func TestViewerTheScriptFillsWritesNoSrc(t *testing.T) {
	v := viewer()
	v.Src, v.Alt, v.Caption = "", "", ""
	got := example.Render(v)
	want := `<img class="cui-console-viewer-image" alt=""><p class="cui-console-viewer-caption"></p>`
	if !strings.Contains(got, want) {
		t.Errorf("a viewer the script fills wrote a src or words:\n got: %s\nwant within: %s", got, want)
	}
}

func TestViewerPanics(t *testing.T) {
	example.Panics(t, viewer, []example.Break[console.ConsoleViewer]{
		{Want: "console.ConsoleViewer needs an ID", Change: func(p *console.ConsoleViewer) { p.ID = "" }},
		{Want: "console.ConsoleViewer needs an AriaLabel", Change: func(p *console.ConsoleViewer) { p.AriaLabel = "" }},
		{Want: "console.ConsoleViewer needs a CloseLabel", Change: func(p *console.ConsoleViewer) { p.CloseLabel = "" }},
		{Want: "console.ConsoleViewer with no Src has no Alt or Caption: the page's script sets all three", Change: func(p *console.ConsoleViewer) { p.Src = "" }},
	})
}
