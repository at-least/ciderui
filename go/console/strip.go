package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleStripMark is what the reader decided of an item's shot: kept, or to
// redraw. A doubt is ConsoleStripItem.DoubtAriaLabel, beside either.
type ConsoleStripMark struct{ value string }

var (
	ConsoleStripMarkNone = ConsoleStripMark{value: ""}
	// ConsoleStripMarkKept: a check in the picture's corner, beside a doubt's
	// dot.
	ConsoleStripMarkKept = ConsoleStripMark{"kept"}
	// ConsoleStripMarkRedraw: to redraw, a red edge inside the picture.
	ConsoleStripMarkRedraw = ConsoleStripMark{"redraw"}
)

// ConsoleStrip is one row of 16:9 thumbnails scrolled sideways, in groups
// under their names (a storyboard's shots by scene under its player), one of
// them current.
type ConsoleStrip struct {
	AriaLabel string
	Groups    []ConsoleStripGroup
	Class     string
}

type ConsoleStripGroup struct {
	Label string
	Items []ConsoleStripItem
}

type ConsoleStripItem struct {
	Href string
	Src  string
	// Alt is the link's name, with Caption, DoubtAriaLabel and MarkAriaLabel
	// after it.
	Alt  string
	Lazy bool
	// Current: aria-current="true", ringed.
	Current bool
	// Caption: the shot's name under its picture.
	Caption string
	// DoubtAriaLabel: a doubt waiting for the reader's call, a dot in
	// console-status's warning color; "" none.
	DoubtAriaLabel string
	Mark           ConsoleStripMark
	MarkAriaLabel  string
}

func (s ConsoleStrip) Render(w io.Writer) error {
	if s.AriaLabel == "" {
		panic("console.ConsoleStrip needs an AriaLabel")
	}
	if len(s.Groups) == 0 {
		panic("console.ConsoleStrip needs a ConsoleStripGroup")
	}
	currents := 0
	for _, group := range s.Groups {
		for _, item := range group.Items {
			if item.Current {
				currents++
			}
		}
	}
	if currents > 1 {
		panic("console.ConsoleStrip has one current item at most")
	}
	return h.Nav(class.Root("cui-console-strip", s.Class), h.Aria("label", s.AriaLabel),
		g.Map(s.Groups, func(group ConsoleStripGroup) g.Node {
			if group.Label == "" {
				panic("console.ConsoleStripGroup needs a Label")
			}
			if len(group.Items) == 0 {
				panic("console.ConsoleStripGroup needs a ConsoleStripItem")
			}
			return h.Div(h.Class("cui-console-strip-group"),
				h.P(h.Class("cui-console-strip-head"), h.Span(h.Class("cui-console-strip-label"), g.Text(group.Label))),
				g.Map(group.Items, stripItem),
			)
		}),
	).Render(w)
}

func stripItem(item ConsoleStripItem) g.Node {
	if item.Alt == "" {
		panic("console.ConsoleStripItem needs an Alt: it is its link's name")
	}
	if (item.Mark == ConsoleStripMarkNone) != (item.MarkAriaLabel == "") {
		panic("console.ConsoleStripItem needs a MarkAriaLabel with a Mark, and none without")
	}
	var mark g.Node
	if item.Mark != ConsoleStripMarkNone {
		mark = stripMark(item.Mark.value, item.MarkAriaLabel, g.If(item.Mark == ConsoleStripMarkKept, lucide.Check))
	}
	// the doubt before the mark: the stylesheet moves a check that follows a
	// dot out of its corner
	return h.A(h.Class("cui-console-strip-item"), h.Href(item.Href), g.If(item.Current, h.Aria("current", "true")),
		h.Img(h.Src(item.Src), h.Alt(item.Alt), g.If(item.Lazy, h.Loading("lazy"))),
		g.If(item.Caption != "", h.Span(h.Class("cui-console-strip-caption"), g.Text(item.Caption))),
		g.If(item.DoubtAriaLabel != "", stripMark("doubt", item.DoubtAriaLabel, nil)),
		mark,
	)
}

func stripMark(state, ariaLabel string, glyph g.Node) g.Node {
	return h.Span(h.Class("cui-console-strip-mark"), h.Data("state", state), h.Role("img"), h.Aria("label", ariaLabel), glyph)
}
