package console

import (
	"io"
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

// ConsoleChart is the frame around the page's plot; the page draws the plot.
type ConsoleChart struct {
	// Peak and Half label the top and middle gridlines.
	Peak string
	Half string
	// Empty is the note centered on a plot with no data ("No data in this
	// window"); "" for a plot with data.
	Empty string
	// AxisStart and AxisEnd label the first and last buckets.
	AxisStart string
	AxisEnd   string
	// Legend: for a chart of more than one series.
	Legend []ConsoleChartSeries
	Class  string
	// Children: <svg class="cui-console-chart-svg"> holding the gridlines and
	// marks.
	Children g.Node
}

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

// ConsoleSparklinePoint is in the sparkline's viewBox units, Y down.
type ConsoleSparklinePoint struct {
	X float64
	Y float64
}

type ConsoleSparkline struct {
	// Label is what the line shows, in words ("3 events over 2 days").
	Label  string
	Width  uint
	Height uint
	// Points run from X 0 to X Width: the area under the line is closed at
	// those two ends.
	Points []ConsoleSparklinePoint
	Class  string
}

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
	// raw, not g.El: g.El writes no self-closing tag, and the example's marks
	// are <polygon …/> and <polyline …/>; raw is safe, the points are numbers
	marks := `<polygon class="cui-console-chart-sparkline-area" points="0,` + height + " " + line + " " + width + "," + height + `"/>` +
		`<polyline class="cui-console-chart-sparkline-line" points="` + line + `"/>`
	return g.El("svg", class.Root("cui-console-chart-sparkline", s.Class),
		g.Attr("viewBox", "0 0 "+width+" "+height), h.Width(width), h.Height(height), h.Role("img"), h.Aria("label", s.Label),
		g.Raw(marks),
	).Render(w)
}
