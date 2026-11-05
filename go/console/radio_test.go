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

func TestRadioPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleRadio {
		return console.ConsoleRadio{Name: "a", Value: "a", Checked: false, Caption: "", Attrs: nil, Class: "", Children: g.Text("A")}
	}, []example.Break[console.ConsoleRadio]{
		{Want: "console.ConsoleRadio needs a Name", Change: func(p *console.ConsoleRadio) { p.Name = "" }},
		{Want: "console.ConsoleRadio needs its label as Children", Change: func(p *console.ConsoleRadio) { p.Children = nil }},
	})
}
