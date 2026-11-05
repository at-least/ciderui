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

func TestChartPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleChart {
		return console.ConsoleChart{AriaLabel: "Events: 3", Scale: nil, Empty: "", Axis: nil, Legend: nil, Class: "", Children: nil}
	}, []example.Break[console.ConsoleChart]{
		{Want: "console.ConsoleChart needs an AriaLabel", Change: func(p *console.ConsoleChart) { p.AriaLabel = "" }},
		{Want: "console.ConsoleChartScale needs a Peak and a Half", Change: func(p *console.ConsoleChart) { p.Scale = &console.ConsoleChartScale{Peak: "3", Half: ""} }},
		{Want: "console.ConsoleChartScale needs a Peak and a Half", Change: func(p *console.ConsoleChart) { p.Scale = &console.ConsoleChartScale{Peak: "", Half: "3"} }},
		{Want: "console.ConsoleChartAxis needs a Start and an End", Change: func(p *console.ConsoleChart) { p.Axis = &console.ConsoleChartAxis{Start: "Oct 22", End: ""} }},
		{Want: "console.ConsoleChartAxis needs a Start and an End", Change: func(p *console.ConsoleChart) { p.Axis = &console.ConsoleChartAxis{Start: "", End: "Oct 28"} }},
		{Want: "console.ConsoleChartSeries needs a Name and a Color", Change: func(p *console.ConsoleChart) { p.Legend = []console.ConsoleChartSeries{{Name: "", Color: "red"}} }},
		{Want: "console.ConsoleChartSeries needs a Name and a Color", Change: func(p *console.ConsoleChart) { p.Legend = []console.ConsoleChartSeries{{Name: "iOS", Color: ""}} }},
	})
}

func TestSparklinePanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleSparkline {
		return console.ConsoleSparkline{AriaLabel: "x", Width: 10, Height: 28, Values: []float64{0, 4}, Class: ""}
	}, []example.Break[console.ConsoleSparkline]{
		{Want: "console.ConsoleSparkline needs an AriaLabel", Change: func(p *console.ConsoleSparkline) { p.AriaLabel = "" }},
		{Want: "console.ConsoleSparkline needs two values or more", Change: func(p *console.ConsoleSparkline) { p.Values = []float64{2} }},
		{Want: "console.ConsoleSparkline's Values are finite and 0 or more", Change: func(p *console.ConsoleSparkline) { p.Values = []float64{-1, 4} }},
		{Want: "console.ConsoleSparkline's Values are finite and 0 or more", Change: func(p *console.ConsoleSparkline) { p.Values = []float64{math.NaN(), 4} }},
		{Want: "console.ConsoleSparkline's Values are finite and 0 or more", Change: func(p *console.ConsoleSparkline) { p.Values = []float64{math.Inf(1), 4} }},
	})
}

// the line beside it is raw: a label is not
func TestSparklineEscapesItsTitle(t *testing.T) {
	got := example.Render(console.ConsoleSparkline{AriaLabel: "a<b", Width: 10, Height: 28, Values: []float64{0, 4}, Class: ""})
	if !strings.Contains(got, "<title>a&lt;b</title>") {
		t.Errorf("got %s", got)
	}
}

// the largest 2 below the top, nothing below Height-1, so the round stroke is
// never clipped; x from 0 to Width, to hundredths
func TestSparklinePoints(t *testing.T) {
	for _, c := range []struct {
		width, height uint
		values        []float64
		want          string
	}{
		{10, 28, []float64{0, 4}, "0,27 10,2"},
		{10, 28, []float64{0, 0.5}, "0,27 10,2"},
		{10, 28, []float64{0, 0}, "0,27 10,27"},
		{1, 4, []float64{0, 0, 0, 0}, "0,3 0.33,3 0.67,3 1,3"},
	} {
		got := example.Render(console.ConsoleSparkline{AriaLabel: "x", Width: c.width, Height: c.height, Values: c.values, Class: ""})
		if want := `points="` + c.want + `"`; !strings.Contains(got, want) {
			t.Errorf("no %s in %s", want, got)
		}
	}
}
