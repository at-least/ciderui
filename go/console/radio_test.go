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
		console.ConsoleRadio{Name: "kind", Value: "service", Checked: true, Caption: "A backend or a server. Listed under Services.", Attrs: nil, Class: "", Children: g.Text("Service")},
		console.ConsoleRadio{Name: "kind", Value: "app", Checked: false, Caption: "A mobile or desktop app. Listed under Apps.", Attrs: nil, Class: "", Children: g.Text("App")},
		console.ConsoleRadio{Name: "kind", Value: "job", Checked: false, Caption: "", Attrs: h.Disabled(), Class: "", Children: g.Text("Job")},
	)
	example.Check(t, radios, "console/console-radio/examples/radios.html")
}

func TestRadioAttrsGoOnItsInput(t *testing.T) {
	got := example.Render(console.ConsoleRadio{Name: "kind", Value: "service", Checked: false, Caption: "", Attrs: h.Required(), Class: "", Children: g.Text("Service")})
	if want := `<label class="cui-console-radio"><input class="cui-console-radio-input" type="radio" name="kind" value="service" required>Service</label>`; got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
}

func TestRadioWithoutItsLabelPanics(t *testing.T) {
	mustPanic(t, "ConsoleRadio needs its label as Children", func() {
		example.Render(console.ConsoleRadio{Name: "a", Value: "a", Checked: false, Caption: "", Attrs: nil, Class: "", Children: nil})
	})
}

func TestRadioAnEmptyNamePanics(t *testing.T) {
	mustPanic(t, "needs a Name", func() {
		example.Render(console.ConsoleRadio{Name: "", Value: "a", Checked: false, Caption: "", Attrs: nil, Class: "", Children: g.Text("A")})
	})
}
