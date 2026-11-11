package console

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

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
	// Scale: with one, the part draws the gridlines it labels.
	Scale *ConsoleChartScale
	// Empty is the note centered on a plot with no data ("No data in this
	// window"); "" for a plot with data.
	Empty string
	Axis  *ConsoleChartAxis
	// Children: <svg class="cui-console-chart-svg"> holding the marks.
	Children g.Node
}

func (c ConsoleChart) Render(w io.Writer) error {
	need("console.ConsoleChart", "an AriaLabel", c.AriaLabel)
	if c.Scale != nil && (c.Scale.Peak == "" || c.Scale.Half == "") {
		panic("console.ConsoleChartScale needs a Peak and a Half")
	}
	if c.Axis != nil && (c.Axis.Start == "" || c.Axis.End == "") {
		panic("console.ConsoleChartAxis needs a Start and an End")
	}
	return h.Figure(h.Class("cui-console-chart"), h.Aria("label", c.AriaLabel),
		h.Div(h.Class("cui-console-chart-plot"),
			g.Iff(c.Scale != nil, func() g.Node {
				return g.Group{
					h.SVG(h.Class("cui-console-chart-grid"), h.Aria("hidden", "true"), gridLine("0"), gridLine("50%")),
					h.Span(h.Class("cui-console-chart-scale"), h.Aria("hidden", "true"), g.Text(c.Scale.Peak)),
					h.Span(h.Class("cui-console-chart-scale"), h.Data("at", "half"), h.Aria("hidden", "true"), g.Text(c.Scale.Half)),
				}
			}),
			optionalText(h.P, "cui-console-chart-empty", c.Empty),
			c.Children,
		),
		g.Iff(c.Axis != nil, func() g.Node {
			return h.Div(h.Class("cui-console-chart-axis"), h.Aria("hidden", "true"), h.Span(g.Text(c.Axis.Start)), h.Span(g.Text(c.Axis.End)))
		}),
	).Render(w)
}

func gridLine(y string) g.Node {
	return g.El("line", g.Attr("x1", "0"), g.Attr("x2", "100%"), g.Attr("y1", y), g.Attr("y2", y))
}
