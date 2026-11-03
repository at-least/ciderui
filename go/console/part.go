package console

import (
	"github.com/at-least/ciderui/go/icon"
	g "maragu.dev/gomponents"
)

func iconPart(i icon.Icon, class string) g.Node {
	if i == (icon.Icon{}) {
		return nil
	}
	return i.Class(class)
}

func optionalAttr(name, value string) g.Node {
	if value == "" {
		return nil
	}
	return g.Attr(name, value)
}
