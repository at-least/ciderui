package console

import (
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type ConsoleChartSeries struct {
	Name string
	// Color is a CSS color, the series' marks' fill.
	Color string
}

// ConsoleChartScale labels a plot's top and middle gridlines.
type ConsoleChartScale struct {
	Peak string
	Half string
}

// ConsoleChartAxis labels a plot's first and last buckets.
type ConsoleChartAxis struct {
	Start string
	End   string
}

// ConsoleChart is the frame around the page's plot; the page draws the plot.
type ConsoleChart struct {
	// AriaLabel: what the chart shows, in words ("Events: 96, peak 24 in one
	// day"): the scale and the axis are hidden from assistive technology.
	AriaLabel string
	// Scale: nil for none.
	Scale *ConsoleChartScale
	// Empty is the note centered on a plot with no data ("No data in this
	// window"); "" for a plot with data.
	Empty string
	// Axis: nil for none.
	Axis *ConsoleChartAxis
	// Legend: for a chart of more than one series.
	Legend []ConsoleChartSeries
	Class  string
	// Children: <svg class="cui-console-chart-svg"> holding the gridlines and
	// marks.
	Children g.Node
}

func (c ConsoleChart) Render(w io.Writer) error {
	if c.AriaLabel == "" {
		panic("console.ConsoleChart needs an AriaLabel")
	}
	if c.Scale != nil && (c.Scale.Peak == "" || c.Scale.Half == "") {
		panic("console.ConsoleChartScale needs a Peak and a Half")
	}
	if c.Axis != nil && (c.Axis.Start == "" || c.Axis.End == "") {
		panic("console.ConsoleChartAxis needs a Start and an End")
	}
	return h.Figure(class.Root("cui-console-chart", c.Class), h.Aria("label", c.AriaLabel),
		h.Div(h.Class("cui-console-chart-plot"),
			g.Iff(c.Scale != nil, func() g.Node {
				return g.Group{
					h.Span(h.Class("cui-console-chart-scale"), h.Aria("hidden", "true"), g.Text(c.Scale.Peak)),
					h.Span(h.Class("cui-console-chart-scale"), h.Data("at", "half"), h.Aria("hidden", "true"), g.Text(c.Scale.Half)),
				}
			}),
			g.If(c.Empty != "", h.P(h.Class("cui-console-chart-empty"), g.Text(c.Empty))),
			c.Children,
		),
		g.Iff(c.Axis != nil, func() g.Node {
			return h.Div(h.Class("cui-console-chart-axis"), h.Aria("hidden", "true"), h.Span(g.Text(c.Axis.Start)), h.Span(g.Text(c.Axis.End)))
		}),
		g.If(len(c.Legend) > 0, h.Div(h.Class("cui-console-chart-legend"),
			g.Map(c.Legend, func(s ConsoleChartSeries) g.Node {
				if s.Name == "" || s.Color == "" {
					panic("console.ConsoleChartSeries needs a Name and a Color")
				}
				return statusPart(s.Name, ConsoleColorNeutral, s.Color, "")
			}),
		)),
	).Render(w)
}

type ConsoleSparkline struct {
	// AriaLabel: what the line shows, in words ("3 events over 2 days").
	AriaLabel string
	Width     uint
	Height    uint
	// Values: the data, drawn on a scale from 0 to the largest.
	Values []float64
	Class  string
}

func (s ConsoleSparkline) Render(w io.Writer) error {
	if s.AriaLabel == "" {
		panic("console.ConsoleSparkline needs an AriaLabel")
	}
	if len(s.Values) < 2 {
		panic("console.ConsoleSparkline needs two values or more")
	}
	largest := 0.0
	for _, v := range s.Values {
		// not v < 0 alone: NaN passes it and max then makes largest NaN
		if !(v >= 0) || math.IsInf(v, 1) {
			panic("console.ConsoleSparkline's Values are finite and 0 or more")
		}
		largest = max(largest, v)
	}
	// App Store Connect's (reference/asc/README.md, Charts): the largest 2
	// below the top, nothing below Height-1, so the round stroke is never
	// clipped
	bottom := float64(s.Height)
	top := 2.0
	floor := bottom - 1
	last := float64(len(s.Values) - 1)
	points := make([]string, len(s.Values))
	for i, v := range s.Values {
		y := floor
		if largest > 0 {
			y = min(floor, bottom-v/largest*(bottom-top))
		}
		x := float64(i) * float64(s.Width) / last
		points[i] = sparklineCoordinate(x) + "," + sparklineCoordinate(y)
	}
	width := strconv.FormatUint(uint64(s.Width), 10)
	height := strconv.FormatUint(uint64(s.Height), 10)
	// raw, not g.El: g.El writes no self-closing tag, and the example's mark is
	// <polyline …/>; raw is safe, the points are numbers
	line := `<polyline class="cui-console-chart-sparkline-line" points="` + strings.Join(points, " ") + `"/>`
	return g.El("svg", class.Root("cui-console-chart-sparkline", s.Class),
		g.Attr("viewBox", "0 0 "+width+" "+height), h.Width(width), h.Height(height), h.Role("img"), h.Aria("label", s.AriaLabel),
		g.El("title", g.Text(s.AriaLabel)),
		g.Raw(line),
	).Render(w)
}

// to hundredths, as App Store Connect's read coordinates are: no float noise
// (7.999999999999999)
func sparklineCoordinate(f float64) string {
	return strconv.FormatFloat(math.Round(f*100)/100, 'f', -1, 64)
}
