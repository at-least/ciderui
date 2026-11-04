package console_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestStripRendersTheExample(t *testing.T) {
	const path = "console/console-strip/examples/strip.html"
	raw := readRaw(path)
	// the example's src for the image of alt, a data: URL too long to copy
	src := func(alt string) string {
		at := strings.Index(raw, `" alt="`+alt+`"`)
		start := strings.LastIndex(raw[:at], `src="`) + len(`src="`)
		return raw[start:at]
	}
	shot := func(scene, n int, doubtAriaLabel string, mark console.ConsoleStripMark, markAriaLabel string, current bool) console.ConsoleStripItem {
		alt := fmt.Sprintf("Shot %d.%d", scene, n)
		return console.ConsoleStripItem{Href: fmt.Sprintf("#shot-%d-%d", scene, n), Src: src(alt), Alt: alt, Lazy: false, Current: current, Caption: fmt.Sprintf("S%02d-%02d", scene, n), DoubtAriaLabel: doubtAriaLabel, Mark: mark, MarkAriaLabel: markAriaLabel}
	}
	const doubt = "Waiting for you"
	none := console.ConsoleStripMarkNone
	kept := console.ConsoleStripMarkKept
	redraw := console.ConsoleStripMarkRedraw
	strip := h.Div(h.Class("cui-console-page"),
		console.ConsoleStrip{AriaLabel: "Shots", Class: "", Groups: []console.ConsoleStripGroup{
			{Label: "Scene 1", Items: []console.ConsoleStripItem{shot(1, 1, "", kept, "Kept", false), shot(1, 2, "", kept, "Kept", false), shot(1, 3, doubt, kept, "Kept", false)}},
			{Label: "Scene 2", Items: []console.ConsoleStripItem{shot(2, 1, "", kept, "Kept", false), shot(2, 2, doubt, redraw, "To redraw", true), shot(2, 3, "", none, "", false), shot(2, 4, doubt, none, "", false)}},
			{Label: "Scene 3", Items: []console.ConsoleStripItem{shot(3, 1, "", none, "", false), shot(3, 2, "", none, "", false), shot(3, 3, "", none, "", false)}},
		}},
	)
	example.Check(t, strip, path)
}

func stripOf(items ...console.ConsoleStripItem) console.ConsoleStrip {
	return console.ConsoleStrip{AriaLabel: "Shots", Groups: []console.ConsoleStripGroup{{Label: "Scene 1", Items: items}}, Class: ""}
}

func stripItem(alt string, current bool, mark console.ConsoleStripMark, markAriaLabel string) console.ConsoleStripItem {
	return console.ConsoleStripItem{Href: "#a", Src: "/a.png", Alt: alt, Lazy: false, Current: current, Caption: "", DoubtAriaLabel: "", Mark: mark, MarkAriaLabel: markAriaLabel}
}

func TestStripPanics(t *testing.T) {
	none := console.ConsoleStripMarkNone
	for _, c := range []struct {
		want  string
		strip console.ConsoleStrip
	}{
		{"needs an AriaLabel", console.ConsoleStrip{AriaLabel: "", Groups: stripOf(stripItem("Shot 1", false, none, "")).Groups, Class: ""}},
		{"needs a ConsoleStripGroup", console.ConsoleStrip{AriaLabel: "Shots", Groups: nil, Class: ""}},
		{"needs a Label", console.ConsoleStrip{AriaLabel: "Shots", Groups: []console.ConsoleStripGroup{{Label: "", Items: []console.ConsoleStripItem{stripItem("Shot 1", false, none, "")}}}, Class: ""}},
		{"needs a ConsoleStripItem", console.ConsoleStrip{AriaLabel: "Shots", Groups: []console.ConsoleStripGroup{{Label: "Scene 1", Items: nil}}, Class: ""}},
		{"one current item at most", stripOf(stripItem("Shot 1", true, none, ""), stripItem("Shot 2", true, none, ""))},
		{"needs an Alt", stripOf(stripItem("", false, none, ""))},
		{"needs a MarkAriaLabel with a Mark", stripOf(stripItem("Shot 1", false, console.ConsoleStripMarkKept, ""))},
		{"a MarkAriaLabel with a Mark, and none without", stripOf(stripItem("Shot 1", false, none, "Kept"))},
	} {
		t.Run(c.want, func(t *testing.T) {
			mustPanic(t, c.want, func() { example.Render(c.strip) })
		})
	}
}
