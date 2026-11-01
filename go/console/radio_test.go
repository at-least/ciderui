package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestRadioRendersTheExample(t *testing.T) {
	radios := h.FieldSet(h.Class("grid gap-3"),
		h.Legend(h.Class("cui-console-label"), g.Text("Kind")),
		console.ConsoleRadio{Name: "kind", Value: "service", Checked: true, Required: false, Disabled: false, Caption: "A backend or a server. Listed under Services.", Class: "", Children: g.Text("Service")},
		console.ConsoleRadio{Name: "kind", Value: "app", Checked: false, Required: false, Disabled: false, Caption: "A mobile or desktop app. Listed under Apps.", Class: "", Children: g.Text("App")},
		console.ConsoleRadio{Name: "kind", Value: "job", Checked: false, Required: false, Disabled: true, Caption: "", Class: "", Children: g.Text("Job")},
	)
	example.Check(t, radios, "console/console-radio/examples/radios.html")
}

func TestRadioARequiredOneSaysSo(t *testing.T) {
	got := example.Render(console.ConsoleRadio{Name: "kind", Value: "service", Checked: false, Required: true, Disabled: false, Caption: "", Class: "", Children: g.Text("Service")})
	if want := `<label class="cui-console-radio"><input class="cui-console-radio-input" type="radio" name="kind" value="service" required>Service</label>`; got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
}

func TestRadioAnEmptyNamePanics(t *testing.T) {
	mustPanic(t, "needs a Name", func() {
		example.Render(console.ConsoleRadio{Name: "", Value: "a", Checked: false, Required: false, Disabled: false, Caption: "", Class: "", Children: g.Text("A")})
	})
}
