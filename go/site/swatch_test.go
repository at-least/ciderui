package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestSwatchRendersTheExample(t *testing.T) {
	example.Check(t, h.Ul(h.Class("cui-swatches"),
		site.Swatch{ID: "color-burgundy", Name: "color", Value: "burgundy", Label: "Burgundy", Checked: true, Class: "", Children: h.Span(h.Class("cui-swatch-color"), h.Style("background-color: #5c2a2e"))},
		site.Swatch{ID: "color-silver", Name: "color", Value: "silver", Label: "Silver", Checked: false, Class: "", Children: h.Span(h.Class("cui-swatch-color"), h.Style("background-color: #e3e4e5"))},
		site.Swatch{ID: "color-sky", Name: "color", Value: "sky", Label: "Sky Blue", Checked: false, Class: "", Children: h.Span(h.Class("cui-swatch-color"))},
	), "site/swatch/examples/swatches.html")
}

func TestSwatchAnEmptyLabelPanics(t *testing.T) {
	mustPanic(t, "needs a Label", func() {
		example.Render(site.Swatch{ID: "x", Name: "x", Value: "x", Label: "", Checked: false, Class: "", Children: g.Text("x")})
	})
}

func TestSwatchAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "Swatch needs an ID", func() {
		example.Render(site.Swatch{ID: "", Name: "x", Value: "x", Label: "x", Checked: false, Class: "", Children: g.Text("x")})
	})
}
