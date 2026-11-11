// Package attr writes a component's attributes.
package attr

import g "maragu.dev/gomponents"

// Class is a component's class attribute: its own classes (block and
// modifiers), then the page's (the component's Class field).
func Class(block, page string) g.Node {
	if page == "" {
		return g.Attr("class", block)
	}
	return g.Attr("class", block+" "+page)
}

// Optional is the attribute name="value", left out when value is "".
func Optional(name, value string) g.Node {
	if value == "" {
		return nil
	}
	return g.Attr(name, value)
}
