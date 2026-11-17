package console

import (
	"slices"
	"strconv"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func optionalText(el func(...g.Node) g.Node, class, text string) g.Node {
	if text == "" {
		return nil
	}
	return el(h.Class(class), g.Text(text))
}

func pageLink(class, href string, current bool, content g.Node) g.Node {
	return h.A(h.Class(class), h.Href(href), g.If(current, h.Aria("current", "page")), content)
}

func needCurrent(part, current string, hrefs []string) {
	if current != "" && !slices.Contains(hrefs, current) {
		panic(part + ".Current " + strconv.Quote(current) + " matches no item's Href")
	}
}
