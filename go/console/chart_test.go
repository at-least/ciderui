package console_test

import (
	"math"
	"strings"
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
		console.ConsoleChart{AriaLabel: "Events: 96, peak 24 in one day, by release: 2.4.1 ×60, 2.4.0 ×36", Scale: &console.ConsoleChartScale{Peak: "peak 24 · Oct 27", Half: "12"}, Empty: "", Axis: &console.ConsoleChartAxis{Start: "Oct 22", End: "Oct 28"}, Legend: legend, Class: "",
			Children: g.Raw(`<svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140" preserveAspectRatio="none">` + chartGrid + bars + `</svg>`)},
	)
	example.Check(t, chart, "console/console-chart/examples/chart.html")

	empty := h.Div(h.Class("cui-console-page max-w-xl"),
		console.ConsoleChart{AriaLabel: "Events: 0", Scale: nil, Empty: "No unhandled errors in this window", Axis: &console.ConsoleChartAxis{Start: "Oct 22", End: "Oct 28"}, Legend: nil, Class: "",
			Children: g.Raw(`<svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140" preserveAspectRatio="none">` + chartGrid + `</svg>`)},
	)
	example.Check(t, empty, "console/console-chart/examples/empty.html")

	sparkline := h.Div(h.Class("cui-console-page"),
		console.ConsoleSparkline{AriaLabel: "53 events over 7 days", Width: 120, Height: 28, Values: []float64{4, 6, 3, 10, 8, 13, 9}, Class: ""},
	)
	example.Check(t, sparkline, "console/console-chart/examples/sparkline.html")
}

func chartWith(change func(*console.ConsoleChart)) console.ConsoleChart {
	c := console.ConsoleChart{AriaLabel: "Events: 3", Scale: nil, Empty: "", Axis: nil, Legend: nil, Class: "", Children: nil}
	change(&c)
	return c
}

func TestChartPanics(t *testing.T) {
	for _, c := range []struct {
		want   string
		change func(*console.ConsoleChart)
	}{
		{"ConsoleChart needs an AriaLabel", func(c *console.ConsoleChart) { c.AriaLabel = "" }},
		{"ConsoleChartScale needs a Peak and a Half", func(c *console.ConsoleChart) { c.Scale = &console.ConsoleChartScale{Peak: "3", Half: ""} }},
		{"ConsoleChartAxis needs a Start and an End", func(c *console.ConsoleChart) { c.Axis = &console.ConsoleChartAxis{Start: "Oct 22", End: ""} }},
		{"ConsoleChartSeries needs a Name and a Color", func(c *console.ConsoleChart) { c.Legend = []console.ConsoleChartSeries{{Name: "", Color: "red"}} }},
		{"ConsoleChartSeries needs a Name and a Color", func(c *console.ConsoleChart) { c.Legend = []console.ConsoleChartSeries{{Name: "iOS", Color: ""}} }},
	} {
		mustPanic(t, c.want, func() { example.Render(chartWith(c.change)) })
	}
}

func TestSparklineDrawsTheLineAloneAndSaysItsLabelAsATitle(t *testing.T) {
	got := example.Render(console.ConsoleSparkline{AriaLabel: "3 events", Width: 10, Height: 28, Values: []float64{0, 4}, Class: ""})
	if strings.Contains(got, "polygon") || !strings.Contains(got, "<title>3 events</title><polyline ") {
		t.Errorf("got %s", got)
	}
}

func TestSparklineEscapesItsTitle(t *testing.T) {
	got := example.Render(console.ConsoleSparkline{AriaLabel: "a<b", Width: 10, Height: 28, Values: []float64{0, 4}, Class: ""})
	if !strings.Contains(got, "<title>a&lt;b</title>") {
		t.Errorf("got %s", got)
	}
}

func TestSparklineWithoutAnAriaLabelPanics(t *testing.T) {
	mustPanic(t, "ConsoleSparkline needs an AriaLabel", func() {
		example.Render(console.ConsoleSparkline{AriaLabel: "", Width: 1, Height: 1, Values: nil, Class: ""})
	})
}

func TestSparklineOfOneValuePanics(t *testing.T) {
	mustPanic(t, "ConsoleSparkline needs two values or more", func() {
		example.Render(console.ConsoleSparkline{AriaLabel: "x", Width: 10, Height: 4, Values: []float64{2}, Class: ""})
	})
}

func TestSparklineSpacesItsValuesEvenlyFromXZeroToXWidth(t *testing.T) {
	got := example.Render(console.ConsoleSparkline{AriaLabel: "x", Width: 1, Height: 4, Values: make([]float64, 50), Class: ""})
	if !strings.Contains(got, ` 1,3"/>`) || !strings.Contains(got, `points="0,3 `) {
		t.Errorf("got %s", got)
	}
}

func TestSparklineDrawsA0HalfItsStrokeAboveTheBottom(t *testing.T) {
	got := example.Render(console.ConsoleSparkline{AriaLabel: "x", Width: 10, Height: 28, Values: []float64{0, 4}, Class: ""})
	if !strings.Contains(got, `<polyline class="cui-console-chart-sparkline-line" points="0,27 10,2"/>`) {
		t.Errorf("got %s", got)
	}
}

func TestSparklineOfOnly0sLiesHalfItsStrokeAboveTheBottom(t *testing.T) {
	got := example.Render(console.ConsoleSparkline{AriaLabel: "x", Width: 10, Height: 28, Values: []float64{0, 0}, Class: ""})
	if !strings.Contains(got, `points="0,27 10,27"`) {
		t.Errorf("got %s", got)
	}
}

func TestSparklineANegativeNaNOrInfiniteValuePanics(t *testing.T) {
	for _, bad := range []float64{-1, math.NaN(), math.Inf(1)} {
		mustPanic(t, "ConsoleSparkline's Values are finite and 0 or more", func() {
			example.Render(console.ConsoleSparkline{AriaLabel: "x", Width: 10, Height: 28, Values: []float64{bad, 4}, Class: ""})
		})
	}
}
