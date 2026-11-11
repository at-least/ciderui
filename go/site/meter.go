package site

import (
	"io"
	"math"
	"strconv"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Meter is the App Store's ratings bars: one row per share, a thin track
// filled as far as its share, its label before it and its value after.
type Meter struct {
	Rows []MeterRow
}

// MeterRow is one of a [Meter]'s rows. The track is hidden from assistive
// technology: Label and Value carry the share in words.
type MeterRow struct {
	Label string
	// Share: how much of the track is filled, 0 to 1.
	Share float64
	Value string
}

func (m Meter) Render(w io.Writer) error {
	if len(m.Rows) == 0 {
		panic("site.Meter needs Rows")
	}
	return h.Div(h.Class("cui-meter"), g.Map(m.Rows, MeterRow.node)).Render(w)
}

func (r MeterRow) node() g.Node {
	if math.IsNaN(r.Share) || r.Share < 0 || r.Share > 1 {
		panic("site.MeterRow's Share is not between 0 and 1: " + strconv.FormatFloat(r.Share, 'g', -1, 64))
	}
	if r.Value == "" {
		panic("site.MeterRow needs a Value: the track is hidden from assistive technology")
	}
	return h.Div(h.Class("cui-meter-row"),
		g.If(r.Label != "", h.Span(h.Class("cui-meter-label"), g.Text(r.Label))),
		h.Span(h.Class("cui-meter-track"), h.Aria("hidden", "true"),
			h.Span(h.Class("cui-meter-fill"), h.Style("width: "+strconv.FormatFloat(math.Round(r.Share*1000)/10, 'f', -1, 64)+"%")),
		),
		h.Span(h.Class("cui-meter-value"), g.Text(r.Value)),
	)
}
