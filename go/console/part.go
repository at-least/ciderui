package console

import (
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
)

func iconPart(i lucide.Icon, class string) g.Node {
	if i == (lucide.Icon{}) {
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
