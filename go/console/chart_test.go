package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

const chartGrid = `<line class="cui-console-chart-grid" x1="0" x2="600" y1="105" y2="105"/><line class="cui-console-chart-grid" x1="0" x2="600" y1="70" y2="70"/><line class="cui-console-chart-grid" x1="0" x2="600" y1="35" y2="35"/>`

func TestChartRendersTheExamples(t *testing.T) {
	bars := `<g class="cui-console-chart-bar"><title>Oct 22: 8</title><rect class="cui-console-chart-mark" x="3" y="94" width="80" height="22" fill="var(--cui-c-console-brand-accent-primary)"/><rect class="cui-console-chart-mark" x="3" y="116" width="80" height="24" fill="var(--cui-c-console-neutral-label-secondary)"/></g>` +
		`<g class="cui-console-chart-bar"><title>Oct 24: 14</title><rect class="cui-console-chart-mark" x="89" y="60" width="80" height="40" fill="var(--cui-c-console-brand-accent-primary)"/><rect class="cui-console-chart-mark" x="89" y="100" width="80" height="40" fill="var(--cui-c-console-neutral-label-secondary)"/></g>` +
		`<g class="cui-console-chart-bar"><title>Oct 27: 24</title><rect class="cui-console-chart-mark" x="175" y="4" width="80" height="96" fill="var(--cui-c-console-brand-accent-primary)"/><rect class="cui-console-chart-mark" x="175" y="100" width="80" height="40" fill="var(--cui-c-console-neutral-label-secondary)"/></g>`
	chart := h.Div(h.Class("cui-console-page max-w-xl"),
		console.ConsoleChart{AriaLabel: "Events: 96, peak 24 in one day, by release: 2.4.1 ×60, 2.4.0 ×36", Scale: &console.ConsoleChartScale{Peak: "peak 24 · Oct 27", Half: "12"}, Empty: "", Axis: &console.ConsoleChartAxis{Start: "Oct 22", End: "Oct 28"}, Class: "",
			Children: g.Raw(`<svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140" preserveAspectRatio="none">` + chartGrid + bars + `</svg>`)},
	)
	example.Check(t, chart, "console/console-chart/examples/chart.html")

	empty := h.Div(h.Class("cui-console-page max-w-xl"),
		console.ConsoleChart{AriaLabel: "Events: 0", Scale: nil, Empty: "No unhandled errors in this window", Axis: &console.ConsoleChartAxis{Start: "Oct 22", End: "Oct 28"}, Class: "",
			Children: g.Raw(`<svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140" preserveAspectRatio="none">` + chartGrid + `</svg>`)},
	)
	example.Check(t, empty, "console/console-chart/examples/empty.html")

}

func TestChartPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleChart {
		return console.ConsoleChart{AriaLabel: "Events: 3", Scale: nil, Empty: "", Axis: nil, Class: "", Children: nil}
	}, []example.Break[console.ConsoleChart]{
		{Want: "console.ConsoleChart needs an AriaLabel", Change: func(p *console.ConsoleChart) { p.AriaLabel = "" }},
		{Want: "console.ConsoleChartScale needs a Peak and a Half", Change: func(p *console.ConsoleChart) { p.Scale = &console.ConsoleChartScale{Peak: "3", Half: ""} }},
		{Want: "console.ConsoleChartScale needs a Peak and a Half", Change: func(p *console.ConsoleChart) { p.Scale = &console.ConsoleChartScale{Peak: "", Half: "3"} }},
		{Want: "console.ConsoleChartAxis needs a Start and an End", Change: func(p *console.ConsoleChart) { p.Axis = &console.ConsoleChartAxis{Start: "Oct 22", End: ""} }},
		{Want: "console.ConsoleChartAxis needs a Start and an End", Change: func(p *console.ConsoleChart) { p.Axis = &console.ConsoleChartAxis{Start: "", End: "Oct 28"} }},
	})
}
