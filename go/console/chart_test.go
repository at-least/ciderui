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
	legend := []console.ConsoleChartSeries{
		{Name: "2.4.1", Color: "var(--cui-c-console-brand-accent-primary)"},
		{Name: "2.4.0", Color: "var(--cui-c-console-neutral-label-secondary)"},
	}
	chart := h.Div(h.Class("cui-console-page max-w-xl"),
		console.ConsoleChart{Peak: "peak 24 · Oct 27", Half: "12", Empty: "", AxisStart: "Oct 22", AxisEnd: "Oct 28", Legend: legend, Class: "",
			Children: g.Raw(`<svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140" preserveAspectRatio="none" role="img" aria-label="Events: 96, peak 24 in one day, by release: 2.4.1 ×60, 2.4.0 ×36">` + chartGrid + bars + `</svg>`)},
	)
	example.Check(t, chart, "console/console-chart/examples/chart.html")

	empty := h.Div(h.Class("cui-console-page max-w-xl"),
		console.ConsoleChart{Peak: "", Half: "", Empty: "No unhandled errors in this window", AxisStart: "Oct 22", AxisEnd: "Oct 28", Legend: nil, Class: "",
			Children: g.Raw(`<svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140" preserveAspectRatio="none" role="img" aria-label="Events: 0">` + chartGrid + `</svg>`)},
	)
	example.Check(t, empty, "console/console-chart/examples/empty.html")

	points := []console.ConsoleSparklinePoint{{X: 0, Y: 20}, {X: 20, Y: 16}, {X: 40, Y: 22}, {X: 60, Y: 8}, {X: 80, Y: 12}, {X: 100, Y: 4}, {X: 120, Y: 10}}
	sparkline := h.Div(h.Class("cui-console-page"),
		console.ConsoleSparkline{Label: "42 events over 7 days", Width: 120, Height: 28, Points: points, Class: ""},
	)
	example.Check(t, sparkline, "console/console-chart/examples/sparkline.html")
}

func TestChartAPeakWithoutItsHalfPanics(t *testing.T) {
	mustPanic(t, "a peak and a half, both or neither", func() {
		example.Render(console.ConsoleChart{Peak: "peak 3", Half: "", Empty: "", AxisStart: "", AxisEnd: "", Legend: nil, Class: "", Children: nil})
	})
}

func TestChartAnAxisWithoutItsEndPanics(t *testing.T) {
	mustPanic(t, "a start and an end, both or neither", func() {
		example.Render(console.ConsoleChart{Peak: "", Half: "", Empty: "", AxisStart: "Oct 22", AxisEnd: "", Legend: nil, Class: "", Children: nil})
	})
}

func TestSparklineWithoutALabelPanics(t *testing.T) {
	mustPanic(t, "needs a Label", func() {
		example.Render(console.ConsoleSparkline{Label: "", Width: 1, Height: 1, Points: nil, Class: ""})
	})
}
