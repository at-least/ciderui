package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/attr"
	"github.com/at-least/ciderui/go/internal/need"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleStripMark is who an item waits on, a dot in its picture's corner:
// the reader, or someone at work on it.
type ConsoleStripMark struct{ value string }

var (
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
	// Current: the current item's Href, ringed; "" for none.
	Current string
}

type ConsoleStripGroup struct {
	Title string
	Items []ConsoleStripItem
}

type ConsoleStripItem struct {
	Href string
	Src  string
	// Alt is the link's name, with Label, Detail, Tip and Mark's AriaLabel
	// after it.
	Alt  string
	Crop ConsoleStripItemCrop
	// Label: the shot's name under its picture.
	Label string
	// Detail: a second line under the Label.
	Detail string
	// Unused: dimmed.
	Unused bool
	// Tip: a line shown over the item while the pointer is on it (a shot's
	// first spoken line).
	Tip  string
	Mark *ConsoleStripItemMark
}

// ConsoleStripItemMark is the dot in an item's picture's corner; an item
// with no mark has a nil one.
type ConsoleStripItemMark struct {
	State     ConsoleStripMark
	AriaLabel string
}

func (s ConsoleStrip) Render(w io.Writer) error {
	need.Field("console.ConsoleStrip", "an AriaLabel", s.AriaLabel)
	need.Some("console.ConsoleStrip", "Groups", s.Groups)
	var hrefs []string
	groups := make(g.Group, len(s.Groups))
	for i, group := range s.Groups {
		path := need.At("console.ConsoleStrip.Groups", i)
		need.Field(path, "a Title", group.Title)
		need.Some(path, "Items", group.Items)
		items := make(g.Group, len(group.Items))
		for j, item := range group.Items {
			items[j] = stripItem(need.At(path+".Items", j), item, item.Href == s.Current)
			hrefs = append(hrefs, item.Href)
		}
		groups[i] = h.Div(h.Class("cui-console-strip-group"),
			h.P(h.Class("cui-console-strip-header"), h.Span(h.Class("cui-console-strip-title"), g.Text(group.Title))),
			items,
		)
	}
	needCurrent("console.ConsoleStrip", s.Current, hrefs)
	return h.Nav(h.Class("cui-console-strip"), h.Aria("label", s.AriaLabel), groups).Render(w)
}

func stripItem(path string, item ConsoleStripItem, current bool) g.Node {
	need.Field(path, "an Href", item.Href)
	need.Field(path, "a Src", item.Src)
	need.Field(path, "an Alt: it is its link's name", item.Alt)
	if item.Detail != "" && item.Label == "" {
		panic(path + " needs a Label with a Detail")
	}
	var mark g.Node
	if item.Mark != nil {
		need.Field(path+".Mark", "a State: ConsoleStripMarkWaiting or ConsoleStripMarkWorking", item.Mark.State.value)
		need.Field(path+".Mark", "an AriaLabel", item.Mark.AriaLabel)
		mark = h.Span(h.Class("cui-console-strip-mark"), h.Data("state", item.Mark.State.value), h.Role("img"), h.Aria("label", item.Mark.AriaLabel))
	}
	return h.A(h.Class("cui-console-strip-item"), h.Href(item.Href), attr.Optional("data-crop", item.Crop.value),
		g.If(item.Unused, h.Data("state", "unused")), g.If(current, h.Aria("current", "true")),
		h.Img(h.Src(item.Src), h.Alt(item.Alt), h.Loading("lazy")),
		optionalText(h.Span, "cui-console-strip-label", item.Label),
		optionalText(h.Span, "cui-console-strip-detail", item.Detail),
		optionalText(h.Span, "cui-console-strip-tip", item.Tip),
		mark,
	)
}
