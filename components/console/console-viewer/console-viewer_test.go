//go:build browser

package consoleviewer

import (
	"strings"
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	asc "github.com/at-least/ciderui/reference/asc/results"
)

const v = ".cui-console-viewer"

// @invariant: the example renders the viewer empty; a page's script sets the picture and the label before
// it opens it, as this does (a 1600x900 gray picture)
const picture = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNjAwIiBoZWlnaHQ9IjkwMCI+PHJlY3Qgd2lkdGg9IjE2MDAiIGhlaWdodD0iOTAwIiBmaWxsPSIjOGU4ZTkzIi8+PC9zdmc+"

func filled(label string) string {
	html := strings.Replace(h.Example("viewer"), `<img class="cui-console-viewer-image" alt="">`, `<img class="cui-console-viewer-image" src="`+picture+`" alt="The cover">`, 1)
	return strings.Replace(html, `<p class="cui-console-viewer-label"></p>`, `<p class="cui-console-viewer-label">`+label+`</p>`, 1)
}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "viewer",
			HTML:   filled("Drawn 2 November"),
			Click:  []string{`[command="show-modal"]`},
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: v, Props: []string{"display", "flex-direction", "row-gap", "padding-top", "border-top-left-radius", "background-color"}, Golden: "Quick Look's: the picture on the console's ground, 8px in, 20px corners, its label 8px under it"},
				{Target: v, Props: []string{"box-width", "box-height"}, Golden: "the picture whole, fit to the window less 16px at each side, never past it"},
				{Target: v + "-image", Props: []string{"border-top-left-radius"}, ASC: h.ASCProbe(asc.Media, "thumb.img")},
				{Target: v + "-image", Props: []string{"object-fit"}, Golden: "contain, not App Store Connect's thumbnail's fill: the picture whole, at its own shape"},
				{Target: v + "-label", Props: []string{"color", "font-size", "line-height", "text-align"}, Golden: "the secondary label at the body's size, centered under the picture"},
				{Target: v + "-close", Props: []string{"position", "align-self", "box-top"}, Golden: "the close button in its own row at the trailing end above the picture, as Quick Look's title bar: over the picture's corner a white picture hid it"},
				{Target: v + "-image", Props: []string{"box-top"}, Golden: "the picture under the close button's row, 8px below it"},
				{Target: v, Pseudo: "::backdrop", Props: []string{"background-color"}, Golden: "the dialog's scrim behind it (--cui-c-scrim)"},
			},
		},
		{
			Name:   "viewer no label",
			HTML:   filled(""),
			Click:  []string{`[command="show-modal"]`},
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: v + "-label", Props: []string{"display"}, Golden: "no label, no room for one"},
			},
		},
	})
}
