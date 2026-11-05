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

func TestCheckboxAttrsGoOnItsInput(t *testing.T) {
	got := example.Render(site.Checkbox{ID: "x", Name: "x", Label: "x", Checked: true, Attrs: h.Disabled(), Class: ""})
	if !strings.Contains(got, `<input class="cui-checkbox-input" type="checkbox" id="x" name="x" checked disabled>`) {
		t.Errorf("got %s", got)
	}
}

func TestCheckboxPanics(t *testing.T) {
	example.Panics(t, func() site.Checkbox {
		return site.Checkbox{ID: "x", Name: "x", Label: "x", Checked: false, Attrs: nil, Class: ""}
	}, []example.Break[site.Checkbox]{
		{Want: "site.Checkbox needs an ID", Change: func(p *site.Checkbox) { p.ID = "" }},
		{Want: "site.Checkbox needs a Name", Change: func(p *site.Checkbox) { p.Name = "" }},
		{Want: "site.Checkbox needs a Label", Change: func(p *site.Checkbox) { p.Label = "" }},
	})
}
