// Package class writes a component's root class.
package class

import g "maragu.dev/gomponents"

// Root is a component's class attribute: its own classes (block and
// modifiers), then the page's (the component's Class field).
func Root(block, page string) g.Node {
	if page == "" {
		return g.Attr("class", block)
	}
	return g.Attr("class", block+" "+page)
}
