package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestChartRendersTheExamples(t *testing.T) {
	bars := `<g class="cui-console-chart-bar"><title>Oct 22: 8</title><rect class="cui-console-chart-mark" x="3" y="95" width="80" height="45"/></g>` +
		`<g class="cui-console-chart-bar"><title>Oct 24: 14</title><rect class="cui-console-chart-mark" x="89" y="61" width="80" height="79"/></g>` +
		`<g class="cui-console-chart-bar"><title>Oct 27: 24</title><rect class="cui-console-chart-mark" x="175" y="4" width="80" height="136"/></g>`
	chart := h.Div(h.Class("cui-console-page max-w-xl"),
		console.ConsoleChart{AriaLabel: "Events: 46, peak 24 in one day", Scale: &console.ConsoleChartScale{Peak: "peak 24 · Oct 27", Half: "12"}, Empty: "", Axis: &console.ConsoleChartAxis{Start: "Oct 22", End: "Oct 28"},
			Children: g.Raw(`<svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140" preserveAspectRatio="none">` + bars + `</svg>`)},
	)
	example.Check(t, chart, "console/console-chart/examples/chart.html")

	empty := h.Div(h.Class("cui-console-page max-w-xl"),
		console.ConsoleChart{AriaLabel: "Events: 0", Scale: nil, Empty: "No unhandled errors in this window", Axis: &console.ConsoleChartAxis{Start: "Oct 22", End: "Oct 28"},
			Children: g.Raw(`<svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140" preserveAspectRatio="none"></svg>`)},
	)
	example.Check(t, empty, "console/console-chart/examples/empty.html")
}

func TestChartPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleChart {
		return console.ConsoleChart{AriaLabel: "Events: 3", Scale: nil, Empty: "", Axis: nil, Children: nil}
	}, []example.Break[console.ConsoleChart]{
		{Want: "console.ConsoleChart needs an AriaLabel", Change: func(p *console.ConsoleChart) { p.AriaLabel = "" }},
		{Want: "console.ConsoleChartScale needs a Peak and a Half", Change: func(p *console.ConsoleChart) { p.Scale = &console.ConsoleChartScale{Peak: "3", Half: ""} }},
		{Want: "console.ConsoleChartScale needs a Peak and a Half", Change: func(p *console.ConsoleChart) { p.Scale = &console.ConsoleChartScale{Peak: "", Half: "3"} }},
		{Want: "console.ConsoleChartAxis needs a Start and an End", Change: func(p *console.ConsoleChart) { p.Axis = &console.ConsoleChartAxis{Start: "Oct 22", End: ""} }},
		{Want: "console.ConsoleChartAxis needs a Start and an End", Change: func(p *console.ConsoleChart) { p.Axis = &console.ConsoleChartAxis{Start: "", End: "Oct 28"} }},
	})
}
