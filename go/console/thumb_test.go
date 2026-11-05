package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestThumbRendersTheExample(t *testing.T) {
	const path = "console/console-thumb/examples/thumbs.html"
	thumbs := h.Div(h.Class("cui-console-page"),
		h.Div(h.Class("cui-console-thumbs"),
			console.ConsoleThumb{Href: "#checkout", Src: exampleSrc(path, "checkout.png"), Alt: "checkout.png", Caption: "checkout.png", Lazy: false, Class: ""},
			console.ConsoleThumb{Href: "#cart", Src: exampleSrc(path, "cart.png"), Alt: "cart.png", Caption: "cart.png", Lazy: false, Class: ""},
		),
	)
	example.Check(t, thumbs, path)
}

func TestThumbPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleThumb {
		return console.ConsoleThumb{Href: "/a.png", Src: "/a.png", Alt: "a.png", Caption: "a.png", Lazy: false, Class: ""}
	}, []example.Break[console.ConsoleThumb]{
		{Want: "console.ConsoleThumb needs an Alt: it is its link's name", Change: func(p *console.ConsoleThumb) { p.Alt = "" }},
		{Want: "console.ConsoleThumb needs a Caption", Change: func(p *console.ConsoleThumb) { p.Caption = "" }},
	})
}
