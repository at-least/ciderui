package console

import (
	"fmt"
	"io"

	"github.com/at-least/ciderui/go/internal/attr"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleStripMark is who an item waits on, a dot in its picture's corner:
// the reader, or someone at work on it.
type ConsoleStripMark struct{ value string }

var (
	ConsoleStripMarkNone = ConsoleStripMark{value: ""}
	// ConsoleStripMarkWaiting: waiting for the reader (a call to make, a fix
	// to look at), a dot in console-status's warning color.
	ConsoleStripMarkWaiting = ConsoleStripMark{"waiting"}
	// ConsoleStripMarkWorking: someone is at work on it (a redraw, a report
	// being fixed), a gray dot.
	ConsoleStripMarkWorking = ConsoleStripMark{"working"}
)

// ConsoleStripItemCrop is how an item's picture is cropped to 16:9.
type ConsoleStripItemCrop struct{ value string }

var (
	// ConsoleStripItemCropCenter: round its middle, a film's frame.
	ConsoleStripItemCropCenter = ConsoleStripItemCrop{value: ""}
	// ConsoleStripItemCropLeading: to its leading edge, a character
	// sheet's face.
	ConsoleStripItemCropLeading = ConsoleStripItemCrop{"leading"}
)

// ConsoleStrip is one row of 16:9 thumbnails scrolled sideways, in groups
// under their names (a storyboard's shots by scene under its player), one of
// them current.
type ConsoleStrip struct {
	AriaLabel string
	Groups    []ConsoleStripGroup
}

type ConsoleStripGroup struct {
	Label string
	Items []ConsoleStripItem
}

type ConsoleStripItem struct {
	Href string
	Src  string
	// Alt is the link's name, with Caption, Detail, Tip and MarkAriaLabel
	// after it.
	Alt string
	// Current: ringed.
	Current bool
	Crop    ConsoleStripItemCrop
	// Caption: the shot's name under its picture.
	Caption string
	// Detail: a second line under the Caption.
	Detail string
	// Unused: dimmed.
	Unused bool
	// Tip: a line shown over the item while the pointer is on it (a shot's
	// first spoken line).
	Tip           string
	Mark          ConsoleStripMark
	MarkAriaLabel string
}

func (s ConsoleStrip) Render(w io.Writer) error {
	need("console.ConsoleStrip", "an AriaLabel", s.AriaLabel)
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
		panic(fmt.Sprintf("console.ConsoleStrip has %d current items: one at most", currents))
	}
	return h.Nav(h.Class("cui-console-strip"), h.Aria("label", s.AriaLabel),
		g.Map(s.Groups, func(group ConsoleStripGroup) g.Node {
			need("console.ConsoleStripGroup", "a Label", group.Label)
			if len(group.Items) == 0 {
				panic("console.ConsoleStripGroup needs a ConsoleStripItem")
			}
			return h.Div(h.Class("cui-console-strip-group"),
				h.P(h.Class("cui-console-strip-header"), h.Span(h.Class("cui-console-strip-label"), g.Text(group.Label))),
				g.Map(group.Items, stripItem),
			)
		}),
	).Render(w)
}

func stripItem(item ConsoleStripItem) g.Node {
	need("console.ConsoleStripItem", "an Href", item.Href)
	need("console.ConsoleStripItem", "a Src", item.Src)
	if item.Alt == "" {
		panic("console.ConsoleStripItem needs an Alt: it is its link's name")
	}
	if item.Detail != "" && item.Caption == "" {
		panic("console.ConsoleStripItem's Detail needs a Caption above it")
	}
	if (item.Mark == ConsoleStripMarkNone) != (item.MarkAriaLabel == "") {
		panic("console.ConsoleStripItem needs a MarkAriaLabel with a Mark, and none without")
	}
	return h.A(h.Class("cui-console-strip-item"), h.Href(item.Href), attr.Optional("data-crop", item.Crop.value),
		g.If(item.Unused, h.Data("state", "unused")), g.If(item.Current, h.Aria("current", "true")),
		h.Img(h.Src(item.Src), h.Alt(item.Alt), h.Loading("lazy")),
		optionalText(h.Span, "cui-console-strip-caption", item.Caption),
		optionalText(h.Span, "cui-console-strip-detail", item.Detail),
		optionalText(h.Span, "cui-console-strip-tip", item.Tip),
		g.If(item.Mark != ConsoleStripMarkNone, h.Span(h.Class("cui-console-strip-mark"), h.Data("state", item.Mark.value), h.Role("img"), h.Aria("label", item.MarkAriaLabel))),
	)
}
