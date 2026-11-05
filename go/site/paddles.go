package site

import (
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Paddles are a row's previous and next buttons, for a page with a script that
// scrolls the row and disables the paddle at an end; nil for none. The kit
// renders only the start state.
type Paddles struct {
	PreviousAriaLabel string
	NextAriaLabel     string
}

func (p *Paddles) node(owner, class, controls string) g.Node {
	if p == nil {
		return nil
	}
	if p.PreviousAriaLabel == "" {
		panic(owner + "'s Paddles need a PreviousAriaLabel")
	}
	if p.NextAriaLabel == "" {
		panic(owner + "'s Paddles need a NextAriaLabel")
	}
	return g.Group{
		h.Button(h.Class(class), h.Type("button"), h.Aria("controls", controls), h.Aria("label", p.PreviousAriaLabel), h.Disabled(), glyph.PaddlePrevious),
		h.Button(h.Class(class), h.Type("button"), h.Aria("controls", controls), h.Aria("label", p.NextAriaLabel), glyph.PaddleNext),
	}
}
