package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	h "maragu.dev/gomponents/html"
)

func TestOptionTileRendersTheExamples(t *testing.T) {
	example.Check(t, site.OptionTile{ID: "storage-256", Name: "storage", Value: "256", Title: "256GB", Detail: "", Price: "From $1099", Checked: false, Attrs: nil, Class: ""}, "site/option/examples/option.html")
	example.Check(t, site.OptionTile{ID: "model-pro", Name: "model", Value: "pro", Title: "Cider Phone 3", Detail: "6.3-inch display", Price: "", Checked: true, Attrs: nil, Class: ""}, "site/option/examples/checked.html")
}

func TestOptionTileAttrsGoOnItsInput(t *testing.T) {
	got := example.Render(site.OptionTile{ID: "x", Name: "x", Value: "v", Title: "x", Detail: "", Price: "", Checked: false, Attrs: h.Disabled(), Class: ""})
	if !strings.Contains(got, `<input class="cui-option-input" type="radio" id="x" name="x" value="v" disabled>`) {
		t.Errorf("got %s", got)
	}
}

func TestOptionTilePanics(t *testing.T) {
	example.Panics(t, func() site.OptionTile {
		return site.OptionTile{ID: "x", Name: "x", Value: "x", Title: "x", Detail: "", Price: "", Checked: false, Attrs: nil, Class: ""}
	}, []example.Break[site.OptionTile]{
		{Want: "site.OptionTile needs an ID", Change: func(p *site.OptionTile) { p.ID = "" }},
		{Want: "site.OptionTile needs a Name", Change: func(p *site.OptionTile) { p.Name = "" }},
		{Want: "site.OptionTile needs a Title", Change: func(p *site.OptionTile) { p.Title = "" }},
	})
}
