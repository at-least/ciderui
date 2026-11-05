package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestDialogRendersTheExample(t *testing.T) {
	dialog := g.Group{
		h.Button(h.Class("cui-console-btn"), h.Type("button"), g.Attr("commandfor", "ignore"), g.Attr("command", "show-modal"), g.Text("Ignore…")),
		console.ConsoleDialog{ID: "ignore", Title: g.Text("Ignore 3 issues?"), Attrs: nil, Class: "", Children: g.Group{
			h.P(g.Text("Ignored issues come back when the condition you choose is met.")),
			h.Form(h.Method("dialog"), h.Class("cui-console-dialog-actions"),
				h.Button(h.Class("cui-console-btn"), g.Text("Cancel")),
				h.Button(h.Class("cui-console-btn"), h.Data("variant", "filled"), g.Text("Ignore")),
			),
		}},
	}
	example.Check(t, dialog, "console/console-dialog/examples/dialog.html")
}

func TestDialogTakesAttributes(t *testing.T) {
	got := example.Render(console.ConsoleDialog{ID: "confirm", Title: g.Text("Sure?"), Attrs: h.Aria("describedby", "confirm-detail"), Class: "", Children: nil})
	if want := `<dialog class="cui-console-dialog" id="confirm" aria-labelledby="confirm-title" aria-describedby="confirm-detail"><h2 class="cui-console-dialog-title" id="confirm-title">Sure?</h2></dialog>`; got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
}

func TestDialogPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleDialog {
		return console.ConsoleDialog{ID: "x", Title: g.Text("x"), Attrs: nil, Class: "", Children: nil}
	}, []example.Break[console.ConsoleDialog]{
		{Want: "console.ConsoleDialog needs an ID", Change: func(p *console.ConsoleDialog) { p.ID = "" }},
		{Want: "console.ConsoleDialog needs a Title", Change: func(p *console.ConsoleDialog) { p.Title = nil }},
	})
}
