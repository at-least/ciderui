package console_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestStripRendersTheExample(t *testing.T) {
	const path = "console/console-strip/examples/strip.html"
	shot := func(scene, n int, tip string, mark *console.ConsoleStripItemMark) console.ConsoleStripItem {
		alt := fmt.Sprintf("Shot %d.%d", scene, n)
		return console.ConsoleStripItem{Href: fmt.Sprintf("#shot-%d-%d", scene, n), Src: exampleSrc(path, alt), Alt: alt, Crop: console.ConsoleStripItemCropCenter, Label: fmt.Sprintf("S%02d-%02d", scene, n), Detail: "", Unused: false, Tip: tip, Mark: mark}
	}
	const line = "Once upon a time, a miller had three sons."
	waiting := &console.ConsoleStripItemMark{State: console.ConsoleStripMarkWaiting, AriaLabel: "Waiting for you"}
	redraw := &console.ConsoleStripItemMark{State: console.ConsoleStripMarkWorking, AriaLabel: "To redraw"}
	fixed := &console.ConsoleStripItemMark{State: console.ConsoleStripMarkWorking, AriaLabel: "Being fixed"}
	strip := h.Div(h.Class("cui-console-page"),
		console.ConsoleStrip{AriaLabel: "Shots", Current: "#shot-2-2", Groups: []console.ConsoleStripGroup{
			{Title: "Scene 1", Items: []console.ConsoleStripItem{shot(1, 1, line, nil), shot(1, 2, "", nil), shot(1, 3, "", waiting)}},
			{Title: "Scene 2", Items: []console.ConsoleStripItem{shot(2, 1, "", nil), shot(2, 2, "", redraw), shot(2, 3, "", nil), shot(2, 4, "", waiting)}},
			{Title: "Scene 3", Items: []console.ConsoleStripItem{shot(3, 1, "", nil), shot(3, 2, "", fixed), shot(3, 3, "", nil)}},
		}},
	)
	example.Check(t, strip, path)
}

func TestStripRendersTheAssetsExample(t *testing.T) {
	const path = "console/console-strip/examples/assets.html"
	asset := func(href, alt, label, detail string, crop console.ConsoleStripItemCrop, unused bool) console.ConsoleStripItem {
		return console.ConsoleStripItem{Href: href, Src: exampleSrc(path, alt), Alt: alt, Crop: crop, Label: label, Detail: detail, Unused: unused, Tip: "", Mark: nil}
	}
	center := console.ConsoleStripItemCropCenter
	leading := console.ConsoleStripItemCropLeading
	strip := h.Div(h.Class("cui-console-page"),
		console.ConsoleStrip{AriaLabel: "Assets", Current: "#ann", Groups: []console.ConsoleStripGroup{
			{Title: "Characters", Items: []console.ConsoleStripItem{asset("#ann", "Ann’s sheet", "Ann", "3 shots", leading, false), asset("#ben", "Ben’s sheet", "Ben", "Not used", leading, true)}},
			{Title: "Props", Items: []console.ConsoleStripItem{asset("#lamp", "The lamp", "Lamp", "1 shot", center, false)}},
			{Title: "Places", Items: []console.ConsoleStripItem{asset("#field", "The field", "Field", "5 shots", center, false)}},
		}},
	)
	example.Check(t, strip, path)
}

func validStrip() console.ConsoleStrip {
	item := console.ConsoleStripItem{Href: "#a", Src: "/a.png", Alt: "Shot 1", Crop: console.ConsoleStripItemCropCenter, Label: "", Detail: "", Unused: false, Tip: "", Mark: nil}
	return console.ConsoleStrip{AriaLabel: "Shots", Groups: []console.ConsoleStripGroup{{Title: "Scene 1", Items: []console.ConsoleStripItem{item, item}}}, Current: ""}
}

func TestStripPanics(t *testing.T) {
	example.Panics(t, validStrip, []example.Break[console.ConsoleStrip]{
		{Want: "console.ConsoleStrip needs an AriaLabel", Change: func(p *console.ConsoleStrip) { p.AriaLabel = "" }},
		{Want: "console.ConsoleStrip needs Groups", Change: func(p *console.ConsoleStrip) { p.Groups = nil }},
		{Want: "console.ConsoleStrip.Groups[0] needs a Title", Change: func(p *console.ConsoleStrip) { p.Groups[0].Title = "" }},
		{Want: "console.ConsoleStrip.Groups[1] needs Items", Change: func(p *console.ConsoleStrip) {
			p.Groups = append(p.Groups, console.ConsoleStripGroup{Title: "Scene 2", Items: nil})
		}},
		{Want: `console.ConsoleStrip.Current "#b" matches no item's Href`, Change: func(p *console.ConsoleStrip) { p.Current = "#b" }},
		{Want: "console.ConsoleStrip.Groups[1].Items[1] needs an Href", Change: func(p *console.ConsoleStrip) {
			p.Groups = append(p.Groups, console.ConsoleStripGroup{Title: "Scene 2", Items: slices.Clone(p.Groups[0].Items)})
			p.Groups[1].Items[1].Href = ""
		}},
		{Want: "console.ConsoleStrip.Groups[0].Items[0] needs a Src", Change: func(p *console.ConsoleStrip) { p.Groups[0].Items[0].Src = "" }},
		{Want: "console.ConsoleStrip.Groups[0].Items[0] needs an Alt: it is its link's name", Change: func(p *console.ConsoleStrip) { p.Groups[0].Items[0].Alt = "" }},
		{Want: "console.ConsoleStrip.Groups[0].Items[0] needs a Label with a Detail", Change: func(p *console.ConsoleStrip) { p.Groups[0].Items[0].Detail = "3 shots" }},
		{Want: "console.ConsoleStrip.Groups[0].Items[0].Mark needs an AriaLabel", Change: func(p *console.ConsoleStrip) {
			p.Groups[0].Items[0].Mark = &console.ConsoleStripItemMark{State: console.ConsoleStripMarkWaiting, AriaLabel: ""}
		}},
		{Want: "console.ConsoleStrip.Groups[0].Items[0].Mark needs a State: ConsoleStripMarkWaiting or ConsoleStripMarkWorking", Change: func(p *console.ConsoleStrip) {
			p.Groups[0].Items[0].Mark = &console.ConsoleStripItemMark{State: console.ConsoleStripMark{}, AriaLabel: "Waiting for you"}
		}},
	})
}

func TestStripItemWritesNoLabelItHasNone(t *testing.T) {
	want := `<a class="cui-console-strip-item" href="#a"><img src="/a.png" alt="Shot 1" loading="lazy"></a>`
	if got := example.Render(validStrip()); !strings.Contains(got, want) {
		t.Errorf("an item with no Label, Detail, Tip or Mark wrote more than its picture:\n got: %s\nwant: %s", got, want)
	}
}
