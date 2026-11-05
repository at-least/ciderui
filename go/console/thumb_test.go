package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestThumbRendersTheExample(t *testing.T) {
	const path = "console/console-thumb/examples/thumbs.html"
	raw := readRaw(path)
	// the example's src for the image of alt, a data: URL too long to copy
	src := func(alt string) string {
		at := strings.Index(raw, `" alt="`+alt+`"`)
		start := strings.LastIndex(raw[:at], `src="`) + len(`src="`)
		return raw[start:at]
	}
	thumbs := h.Div(h.Class("cui-console-page"),
		h.Div(h.Class("cui-console-thumbs"),
			console.ConsoleThumb{Href: "#checkout", Src: src("checkout.png"), Alt: "checkout.png", Caption: "checkout.png", Lazy: false, Class: ""},
			console.ConsoleThumb{Href: "#cart", Src: src("cart.png"), Alt: "cart.png", Caption: "cart.png", Lazy: false, Class: ""},
		),
	)
	example.Check(t, thumbs, path)
}

func TestThumbAnEmptyAltPanics(t *testing.T) {
	mustPanic(t, "ConsoleThumb needs an Alt", func() {
		example.Render(console.ConsoleThumb{Href: "/a.png", Src: "/a.png", Alt: "", Caption: "a.png", Lazy: false, Class: ""})
	})
}

func TestThumbAnEmptyCaptionPanics(t *testing.T) {
	mustPanic(t, "needs a Caption", func() {
		example.Render(console.ConsoleThumb{Href: "/a.png", Src: "/a.png", Alt: "a.png", Caption: "", Lazy: false, Class: ""})
	})
}
