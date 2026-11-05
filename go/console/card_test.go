package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestMetricRendersTheExample(t *testing.T) {
	cards := []struct {
		title, value, change string
		color                console.ConsoleChangeColor
		caption              string
	}{
		{"Crash-Free Sessions", "99.2%", "+0.4%", console.ConsoleChangeColorSuccess, "Latest release, 2.4.1"},
		{"Unhandled Errors", "1.4K", "+44%", console.ConsoleChangeColorError, "In the last 7 days"},
		{"New Issues", "8", "No change", console.ConsoleChangeColorNeutral, "First seen in the last 7 days"},
	}
	var metrics g.Group
	for _, c := range cards {
		metrics = append(metrics, console.ConsoleMetric{Title: c.title, Value: g.Text(c.value), Change: c.change, ChangeColor: c.color, Caption: c.caption, Href: "", Class: ""})
	}
	example.Check(t, h.Div(h.Class("grid grid-cols-3 gap-5"), metrics), "console/console-card/examples/metrics.html")
}

func TestMetricRendersTheLinkExample(t *testing.T) {
	card := h.Div(h.Class("grid grid-cols-2 gap-5"),
		console.ConsoleMetric{Title: "New Issues", Value: g.Text("8"), Change: "", ChangeColor: console.ConsoleChangeColorNeutral, Caption: "First seen in the last 7 days", Href: "#new-issues", Class: ""},
	)
	example.Check(t, card, "console/console-card/examples/link.html")
}

func TestMetricAValueOfMarkupIsWrittenAsItIs(t *testing.T) {
	got := example.Render(console.ConsoleMetric{Title: "Crash-Free", Value: h.Span(h.Class("text-danger"), g.Text("0%")), Change: "", ChangeColor: console.ConsoleChangeColorNeutral, Caption: "", Href: "", Class: ""})
	if !strings.Contains(got, `<span class="cui-console-card-value"><span class="text-danger">0%</span></span>`) {
		t.Errorf("got %s", got)
	}
}

func TestMetricPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleMetric {
		return console.ConsoleMetric{Title: "x", Value: g.Text("1"), Change: "", ChangeColor: console.ConsoleChangeColorNeutral, Caption: "", Href: "", Class: ""}
	}, []example.Break[console.ConsoleMetric]{
		{Want: "console.ConsoleMetric needs a Title", Change: func(p *console.ConsoleMetric) { p.Title = "" }},
		{Want: "console.ConsoleMetric needs a Value", Change: func(p *console.ConsoleMetric) { p.Value = nil }},
		{Want: "console.ConsoleMetric needs a Value", Change: func(p *console.ConsoleMetric) { p.Value = g.Text("") }},
	})
}
