package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestSwatchRendersTheExample(t *testing.T) {
	example.Check(t, h.Ul(h.Class("cui-swatches"),
		site.Swatch{ID: "color-burgundy", Name: "color", Value: "burgundy", AriaLabel: "Burgundy", Checked: true, Attrs: nil, Class: "", Children: h.Span(h.Class("cui-swatch-color"), h.Style("background-color: #5c2a2e"))},
		site.Swatch{ID: "color-silver", Name: "color", Value: "silver", AriaLabel: "Silver", Checked: false, Attrs: nil, Class: "", Children: h.Span(h.Class("cui-swatch-color"), h.Style("background-color: #e3e4e5"))},
		site.Swatch{ID: "color-sky", Name: "color", Value: "sky", AriaLabel: "Sky Blue", Checked: false, Attrs: nil, Class: "", Children: h.Span(h.Class("cui-swatch-color"))},
	), "site/swatch/examples/swatches.html")
}

func TestSwatchWritesItsTitleAndAttrsOnItsInput(t *testing.T) {
	got := example.Render(site.Swatch{ID: "x", Name: "x", Value: "v", AriaLabel: "銀色", Checked: false, Attrs: h.Disabled(), Class: "", Children: g.Text("x")})
	if !strings.Contains(got, `<input class="cui-swatch-input" type="radio" id="x" name="x" value="v" title="銀色" disabled>`) {
		t.Errorf("got %s", got)
	}
}

func TestSwatchPanics(t *testing.T) {
	example.Panics(t, func() site.Swatch {
		return site.Swatch{ID: "x", Name: "x", Value: "x", AriaLabel: "x", Checked: false, Attrs: nil, Class: "", Children: g.Text("x")}
	}, []example.Break[site.Swatch]{
		{Want: "site.Swatch needs an ID", Change: func(p *site.Swatch) { p.ID = "" }},
		{Want: "site.Swatch needs a Name", Change: func(p *site.Swatch) { p.Name = "" }},
		{Want: "site.Swatch needs an AriaLabel", Change: func(p *site.Swatch) { p.AriaLabel = "" }},
	})
}
