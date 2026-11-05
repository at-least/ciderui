package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	h "maragu.dev/gomponents/html"
)

func TestCheckboxRendersTheExamples(t *testing.T) {
	example.Check(t, site.Checkbox{ID: "remember", Name: "remember", Label: "Remember me", Checked: false, Attrs: nil, Class: ""}, "site/checkbox/examples/checkbox.html")
	example.Check(t, site.Checkbox{ID: "terms", Name: "terms", Label: "I agree to the terms", Checked: true, Attrs: nil, Class: ""}, "site/checkbox/examples/checked.html")
}

func TestCheckboxAnEmptyLabelPanics(t *testing.T) {
	mustPanic(t, "needs a Label", func() {
		example.Render(site.Checkbox{ID: "x", Name: "x", Label: "", Checked: false, Attrs: nil, Class: ""})
	})
}

func TestCheckboxAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "Checkbox needs an ID", func() {
		example.Render(site.Checkbox{ID: "", Name: "x", Label: "x", Checked: false, Attrs: nil, Class: ""})
	})
}

func TestCheckboxAnEmptyNamePanics(t *testing.T) {
	mustPanic(t, "site.Checkbox needs a Name", func() {
		example.Render(site.Checkbox{ID: "x", Name: "", Label: "x", Checked: false, Attrs: nil, Class: ""})
	})
}

func TestCheckboxAttrsGoOnItsInput(t *testing.T) {
	got := example.Render(site.Checkbox{ID: "x", Name: "x", Label: "x", Checked: true, Attrs: h.Disabled(), Class: ""})
	if !strings.Contains(got, `<input class="cui-checkbox-input" type="checkbox" id="x" name="x" checked disabled>`) {
		t.Errorf("got %s", got)
	}
}
