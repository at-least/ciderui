//go:build browser

package consoleartwork

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const a = ".cui-console-artwork"

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "artwork",
			HTML:   h.Example("artwork"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: "img" + a, Props: []string{"border-top-left-radius"}, ASC: h.ASCProbe("media", "thumb.img")},
				{Target: "img" + a, Props: []string{"box-width", "box-height", "object-fit", "background-color"}, Golden: "as wide as its column, 16:9, the picture cropped to fill it on the console card's gray (App Store Connect's frame rgb(242, 242, 242), media.json)"},
				{Target: "span" + a, Props: []string{"display", "align-items", "justify-content", "color", "font-size", "line-height", "text-align", "box-height"}, Golden: "no picture: its words centered in its place, the secondary label at the console's xs size, the same 16:9 box"},
			},
		},
	})
}
