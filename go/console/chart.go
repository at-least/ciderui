package console

import (
	"io"
	"strconv"
	"strings"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleChartSeries is one series of a [ConsoleChart]'s legend.
type ConsoleChartSeries struct {
	// Name is its name.
	Name string
	// Color is its marks' fill, a CSS color.
	Color string
}

// ConsoleChart is cui-console-chart (components/console/console-chart): App Store
// Connect's analytics chart, what is around the page's plot (Children: its
// <svg class="cui-console-chart-svg">, gridlines and marks, which the page
// draws from its data): the scale over it, the empty note, the axis's first
// and last labels under it, and the legend, a console [ConsoleStatus] for each
// series.
type ConsoleChart struct {
	// Peak is the top gridline's value, at the plot's top end; "" for no
	// scale.
	Peak string
	// Half is the middle gridline's value; "" where Peak is.
	Half string
	// Empty is what the plot has none of, centered on it; "" for a plot
	// with data.
	Empty string
	// AxisStart is the first bucket's label under the plot; "" with AxisEnd
	// for no axis.
	AxisStart string
	// AxisEnd is the last bucket's label.
	AxisEnd string
	// Legend is the series, for a chart of more than one; nil for no
	// legend.
	Legend []ConsoleChartSeries
	// Class is the page's utilities on the <figure>; "" for none.
	Class string
	// Children is the plot.
	Children g.Node
}

// Render writes the chart.
func (c ConsoleChart) Render(w io.Writer) error {
	if (c.Peak == "") != (c.Half == "") {
		panic("console.ConsoleChart's scale is a peak and a half, both or neither")
	}
	if (c.AxisStart == "") != (c.AxisEnd == "") {
		panic("console.ConsoleChart's axis is a start and an end, both or neither")
	}
	return h.Figure(class.Root("cui-console-chart", c.Class),
		h.Div(h.Class("cui-console-chart-plot"),
			g.If(c.Peak != "", g.Group{
				h.Span(h.Class("cui-console-chart-scale"), h.Aria("hidden", "true"), g.Text(c.Peak)),
				h.Span(h.Class("cui-console-chart-scale"), h.Data("at", "half"), h.Aria("hidden", "true"), g.Text(c.Half)),
			}),
			g.If(c.Empty != "", h.P(h.Class("cui-console-chart-empty"), g.Text(c.Empty))),
			c.Children,
		),
		g.If(c.AxisStart != "", h.Div(h.Class("cui-console-chart-axis"), h.Aria("hidden", "true"), h.Span(g.Text(c.AxisStart)), h.Span(g.Text(c.AxisEnd)))),
		g.If(len(c.Legend) > 0, h.Div(h.Class("cui-console-chart-legend"),
			g.Map(c.Legend, func(s ConsoleChartSeries) g.Node {
				return ConsoleStatus{Label: s.Name, Color: ConsoleColorNeutral, Series: s.Color, Class: ""}
			}),
		)),
	).Render(w)
}

// ConsoleSparklinePoint is a point of a [ConsoleSparkline]'s line, in its units; Y down.
type ConsoleSparklinePoint struct {
	X float64
	Y float64
}

// ConsoleSparkline is cui-console-chart-sparkline (components/console/console-chart):
// a small line chart, its line in the brand's blue over its area in the
// blue's tint, down to its foot.
type ConsoleSparkline struct {
	// Label is what it shows, in words: the svg's label. Required: an empty
	// one panics.
	Label string
	// Width is its width, in its own units and in CSS pixels.
	Width uint
	// Height is its height, the same.
	Height uint
	// Points are the line's points, from its start (X 0) to its end (X
	// Width).
	Points []ConsoleSparklinePoint
	// Class is the page's utilities on the svg; "" for none.
	Class string
}

// Render writes the sparkline's <svg>.
func (s ConsoleSparkline) Render(w io.Writer) error {
	if s.Label == "" {
		panic("console.ConsoleSparkline needs a Label")
	}
	points := make([]string, len(s.Points))
	for i, p := range s.Points {
		points[i] = strconv.FormatFloat(p.X, 'f', -1, 64) + "," + strconv.FormatFloat(p.Y, 'f', -1, 64)
	}
	line := strings.Join(points, " ")
	width := strconv.FormatUint(uint64(s.Width), 10)
	height := strconv.FormatUint(uint64(s.Height), 10)
	// the marks raw: g.El writes no self-closing tag, and the example's
	// marks are <polygon …/> and <polyline …/>; their points are numbers
	marks := `<polygon class="cui-console-chart-sparkline-area" points="0,` + height + " " + line + " " + width + "," + height + `"/>` +
		`<polyline class="cui-console-chart-sparkline-line" points="` + line + `"/>`
	return g.El("svg", class.Root("cui-console-chart-sparkline", s.Class),
		g.Attr("viewBox", "0 0 "+width+" "+height), h.Width(width), h.Height(height), h.Role("img"), h.Aria("label", s.Label),
		g.Raw(marks),
	).Render(w)
}
