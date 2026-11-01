package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestCheckboxBoxRendersTheExample(t *testing.T) {
	checkboxes := h.Div(h.Class("grid gap-3"),
		h.Label(h.Class("cui-console-checkbox"),
			h.Input(h.Class("cui-console-checkbox-input"), h.Type("checkbox"), h.Name("mute")),
			console.ConsoleCheckboxBox{},
			g.Text("Mute this issue"),
		),
		h.Label(h.Class("cui-console-checkbox"),
			h.Input(h.Class("cui-console-checkbox-input"), h.Type("checkbox"), h.Name("notify"), h.Checked()),
			console.ConsoleCheckboxBox{},
			g.Text("Notify me when it comes back"),
		),
		h.Label(h.Class("cui-console-checkbox"),
			h.Input(h.Class("cui-console-checkbox-input"), h.Type("checkbox"), h.Name("share"), h.Disabled()),
			console.ConsoleCheckboxBox{},
			g.Text("Share with every project"),
		),
		h.Span(h.Class("cui-console-checkbox"),
			h.Input(h.Class("cui-console-checkbox-input"), h.Type("checkbox"), h.Name("all"), h.Aria("label", "Select all")),
			console.ConsoleCheckboxBox{},
		),
	)
	example.Check(t, checkboxes, "console/console-checkbox/examples/checkbox.html")
}
