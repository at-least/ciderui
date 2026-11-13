package console

import (
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// @trap: what carries its own article ("an ID"): a vowel rule would write
// "an URL"
func need(part, what, value string) {
	if value == "" {
		panic(part + " needs " + what)
	}
}

func iconPart(i lucide.Icon, class string) g.Node {
	if i == (lucide.Icon{}) {
		return nil
	}
	return i.Class(class)
}

func optionalText(el func(...g.Node) g.Node, class, text string) g.Node {
	if text == "" {
		return nil
	}
	return el(h.Class(class), g.Text(text))
}

func pageLink(class, href string, current bool, content g.Node) g.Node {
	return h.A(h.Class(class), h.Href(href), g.If(current, h.Aria("current", "page")), content)
}
