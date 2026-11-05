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
	shot := func(scene, n int, tip string, mark console.ConsoleStripMark, markAriaLabel string, current bool) console.ConsoleStripItem {
		alt := fmt.Sprintf("Shot %d.%d", scene, n)
		return console.ConsoleStripItem{Href: fmt.Sprintf("#shot-%d-%d", scene, n), Src: src(alt), Alt: alt, Lazy: false, Current: current, Variant: console.ConsoleStripItemVariantFrame, Caption: fmt.Sprintf("S%02d-%02d", scene, n), Detail: "", Unused: false, Tip: tip, Mark: mark, MarkAriaLabel: markAriaLabel}
	}
	const line = "Once upon a time, a miller had three sons."
	none := console.ConsoleStripMarkNone
	waiting := console.ConsoleStripMarkWaiting
	working := console.ConsoleStripMarkWorking
	strip := h.Div(h.Class("cui-console-page"),
		console.ConsoleStrip{AriaLabel: "Shots", Class: "", Groups: []console.ConsoleStripGroup{
			{Label: "Scene 1", Items: []console.ConsoleStripItem{shot(1, 1, line, none, "", false), shot(1, 2, "", none, "", false), shot(1, 3, "", waiting, "Waiting for you", false)}},
			{Label: "Scene 2", Items: []console.ConsoleStripItem{shot(2, 1, "", none, "", false), shot(2, 2, "", working, "To redraw", true), shot(2, 3, "", none, "", false), shot(2, 4, "", waiting, "Waiting for you", false)}},
			{Label: "Scene 3", Items: []console.ConsoleStripItem{shot(3, 1, "", none, "", false), shot(3, 2, "", working, "Being fixed", false), shot(3, 3, "", none, "", false)}},
		}},
	)
	example.Check(t, strip, path)
}

func TestStripRendersTheAssetsExample(t *testing.T) {
	const path = "console/console-strip/examples/assets.html"
	raw := readRaw(path)
	src := func(alt string) string {
		at := strings.Index(raw, `" alt="`+alt+`"`)
		start := strings.LastIndex(raw[:at], `src="`) + len(`src="`)
		return raw[start:at]
	}
	asset := func(href, alt, caption, detail string, variant console.ConsoleStripItemVariant, unused, current bool) console.ConsoleStripItem {
		return console.ConsoleStripItem{Href: href, Src: src(alt), Alt: alt, Lazy: false, Current: current, Variant: variant, Caption: caption, Detail: detail, Unused: unused, Tip: "", Mark: console.ConsoleStripMarkNone, MarkAriaLabel: ""}
	}
	frame := console.ConsoleStripItemVariantFrame
	sheet := console.ConsoleStripItemVariantSheet
	strip := h.Div(h.Class("cui-console-page"),
		console.ConsoleStrip{AriaLabel: "Assets", Class: "", Groups: []console.ConsoleStripGroup{
			{Label: "Characters", Items: []console.ConsoleStripItem{asset("#ann", "Ann’s sheet", "Ann", "3 shots", sheet, false, true), asset("#ben", "Ben’s sheet", "Ben", "Not used", sheet, true, false)}},
			{Label: "Props", Items: []console.ConsoleStripItem{asset("#lamp", "The lamp", "Lamp", "1 shot", frame, false, false)}},
			{Label: "Places", Items: []console.ConsoleStripItem{asset("#field", "The field", "Field", "5 shots", frame, false, false)}},
		}},
	)
	example.Check(t, strip, path)
}

func stripOf(items ...console.ConsoleStripItem) console.ConsoleStrip {
	return console.ConsoleStrip{AriaLabel: "Shots", Groups: []console.ConsoleStripGroup{{Label: "Scene 1", Items: items}}, Class: ""}
}

func stripItem(alt string, current bool, mark console.ConsoleStripMark, markAriaLabel string) console.ConsoleStripItem {
	return console.ConsoleStripItem{Href: "#a", Src: "/a.png", Alt: alt, Lazy: false, Current: current, Variant: console.ConsoleStripItemVariantFrame, Caption: "", Detail: "", Unused: false, Tip: "", Mark: mark, MarkAriaLabel: markAriaLabel}
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
		{"needs a Caption under a Detail", stripOf(console.ConsoleStripItem{Href: "#a", Src: "/a.png", Alt: "Ann", Lazy: false, Current: false, Variant: console.ConsoleStripItemVariantFrame, Caption: "", Detail: "3 shots", Unused: false, Tip: "", Mark: none, MarkAriaLabel: ""})},
		{"needs a MarkAriaLabel with a Mark", stripOf(stripItem("Shot 1", false, console.ConsoleStripMarkWaiting, ""))},
		{"a MarkAriaLabel with a Mark, and none without", stripOf(stripItem("Shot 1", false, none, "Waiting for you"))},
	} {
		t.Run(c.want, func(t *testing.T) {
			mustPanic(t, c.want, func() { example.Render(c.strip) })
		})
	}
}
