package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	h "maragu.dev/gomponents/html"
)

func TestOptionRendersTheExamples(t *testing.T) {
	example.Check(t, site.Option{ID: "storage-256", Name: "storage", Value: "256", Title: "256GB", Detail: "", Checked: false, Attrs: nil}, "site/option/examples/option.html")
	example.Check(t, site.Option{ID: "model-pro", Name: "model", Value: "pro", Title: "Cider Phone 3", Detail: "6.3-inch display", Checked: true, Attrs: nil}, "site/option/examples/checked.html")
}

func TestOptionAttrsGoOnItsInput(t *testing.T) {
	got := example.Render(site.Option{ID: "x", Name: "x", Value: "v", Title: "x", Detail: "", Checked: false, Attrs: h.Disabled()})
	if !strings.Contains(got, `<input class="cui-option-input" type="radio" id="x" name="x" value="v" disabled>`) {
		t.Errorf("the radio input lost its Attrs: %s", got)
	}
}

func TestOptionPanics(t *testing.T) {
	example.Panics(t, func() site.Option {
		return site.Option{ID: "x", Name: "x", Value: "x", Title: "x", Detail: "", Checked: false, Attrs: nil}
	}, []example.Break[site.Option]{
		{Want: "site.Option needs an ID", Change: func(p *site.Option) { p.ID = "" }},
		{Want: "site.Option needs a Name", Change: func(p *site.Option) { p.Name = "" }},
		{Want: "site.Option needs a Title", Change: func(p *site.Option) { p.Title = "" }},
	})
}
