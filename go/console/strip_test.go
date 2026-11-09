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
	shot := func(scene, n int, tip string, mark console.ConsoleStripMark, markAriaLabel string, current bool) console.ConsoleStripItem {
		alt := fmt.Sprintf("Shot %d.%d", scene, n)
		return console.ConsoleStripItem{Href: fmt.Sprintf("#shot-%d-%d", scene, n), Src: exampleSrc(path, alt), Alt: alt, Current: current, Crop: console.ConsoleStripItemCropCenter, Caption: fmt.Sprintf("S%02d-%02d", scene, n), Detail: "", Unused: false, Tip: tip, Mark: mark, MarkAriaLabel: markAriaLabel}
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
	asset := func(href, alt, caption, detail string, crop console.ConsoleStripItemCrop, unused, current bool) console.ConsoleStripItem {
		return console.ConsoleStripItem{Href: href, Src: exampleSrc(path, alt), Alt: alt, Current: current, Crop: crop, Caption: caption, Detail: detail, Unused: unused, Tip: "", Mark: console.ConsoleStripMarkNone, MarkAriaLabel: ""}
	}
	center := console.ConsoleStripItemCropCenter
	leading := console.ConsoleStripItemCropLeading
	strip := h.Div(h.Class("cui-console-page"),
		console.ConsoleStrip{AriaLabel: "Assets", Class: "", Groups: []console.ConsoleStripGroup{
			{Label: "Characters", Items: []console.ConsoleStripItem{asset("#ann", "Ann’s sheet", "Ann", "3 shots", leading, false, true), asset("#ben", "Ben’s sheet", "Ben", "Not used", leading, true, false)}},
			{Label: "Props", Items: []console.ConsoleStripItem{asset("#lamp", "The lamp", "Lamp", "1 shot", center, false, false)}},
			{Label: "Places", Items: []console.ConsoleStripItem{asset("#field", "The field", "Field", "5 shots", center, false, false)}},
		}},
	)
	example.Check(t, strip, path)
}

func validStrip() console.ConsoleStrip {
	item := console.ConsoleStripItem{Href: "#a", Src: "/a.png", Alt: "Shot 1", Current: false, Crop: console.ConsoleStripItemCropCenter, Caption: "", Detail: "", Unused: false, Tip: "", Mark: console.ConsoleStripMarkNone, MarkAriaLabel: ""}
	return console.ConsoleStrip{AriaLabel: "Shots", Groups: []console.ConsoleStripGroup{{Label: "Scene 1", Items: []console.ConsoleStripItem{item, item}}}, Class: ""}
}

func TestStripPanics(t *testing.T) {
	example.Panics(t, validStrip, []example.Break[console.ConsoleStrip]{
		{Want: "console.ConsoleStrip needs an AriaLabel", Change: func(p *console.ConsoleStrip) { p.AriaLabel = "" }},
		{Want: "console.ConsoleStrip needs a ConsoleStripGroup", Change: func(p *console.ConsoleStrip) { p.Groups = nil }},
		{Want: "console.ConsoleStripGroup needs a Label", Change: func(p *console.ConsoleStrip) { p.Groups[0].Label = "" }},
		{Want: "console.ConsoleStripGroup needs a ConsoleStripItem", Change: func(p *console.ConsoleStrip) { p.Groups[0].Items = nil }},
		{Want: "console.ConsoleStrip has one current item at most", Change: func(p *console.ConsoleStrip) { p.Groups[0].Items[0].Current, p.Groups[0].Items[1].Current = true, true }},
		{Want: "console.ConsoleStripItem needs an Alt: it is its link's name", Change: func(p *console.ConsoleStrip) { p.Groups[0].Items[0].Alt = "" }},
		{Want: "console.ConsoleStripItem's Detail needs a Caption above it", Change: func(p *console.ConsoleStrip) { p.Groups[0].Items[0].Detail = "3 shots" }},
		{Want: "console.ConsoleStripItem needs a MarkAriaLabel with a Mark, and none without", Change: func(p *console.ConsoleStrip) { p.Groups[0].Items[0].Mark = console.ConsoleStripMarkWaiting }},
		{Want: "console.ConsoleStripItem needs a MarkAriaLabel with a Mark, and none without", Change: func(p *console.ConsoleStrip) { p.Groups[0].Items[0].MarkAriaLabel = "Waiting for you" }},
	})
}

func TestStripItemWritesNoCaptionItHasNone(t *testing.T) {
	want := `<a class="cui-console-strip-item" href="#a"><img src="/a.png" alt="Shot 1" loading="lazy"></a>`
	if got := example.Render(validStrip()); !strings.Contains(got, want) {
		t.Errorf("an item with no Caption, Detail, Tip or Mark wrote more than its picture:\n got: %s\nwant: %s", got, want)
	}
}
