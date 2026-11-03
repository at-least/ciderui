package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func TestCheckboxRendersTheExamples(t *testing.T) {
	example.Check(t, site.Checkbox{ID: "remember", Name: "remember", Label: "Remember me", Checked: false, Class: ""}, "site/checkbox/examples/checkbox.html")
	example.Check(t, site.Checkbox{ID: "terms", Name: "terms", Label: "I agree to the terms", Checked: true, Class: ""}, "site/checkbox/examples/checked.html")
}

func TestCheckboxAnEmptyLabelPanics(t *testing.T) {
	mustPanic(t, "needs a Label", func() {
		example.Render(site.Checkbox{ID: "x", Name: "x", Label: "", Checked: false, Class: ""})
	})
}

func TestCheckboxAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "Checkbox needs an ID", func() {
		example.Render(site.Checkbox{ID: "", Name: "x", Label: "x", Checked: false, Class: ""})
	})
}
