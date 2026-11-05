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

func TestOptionTileAnEmptyTitlePanics(t *testing.T) {
	mustPanic(t, "needs a Title", func() {
		example.Render(site.OptionTile{ID: "x", Name: "x", Value: "x", Title: "", Detail: "", Price: "", Checked: false, Attrs: nil, Class: ""})
	})
}

func TestOptionTileAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "OptionTile needs an ID", func() {
		example.Render(site.OptionTile{ID: "", Name: "x", Value: "x", Title: "x", Detail: "", Price: "", Checked: false, Attrs: nil, Class: ""})
	})
}

func TestOptionTileAnEmptyNamePanics(t *testing.T) {
	mustPanic(t, "site.OptionTile needs a Name", func() {
		example.Render(site.OptionTile{ID: "x", Name: "", Value: "x", Title: "x", Detail: "", Price: "", Checked: false, Attrs: nil, Class: ""})
	})
}

func TestOptionTileAttrsGoOnItsInput(t *testing.T) {
	got := example.Render(site.OptionTile{ID: "x", Name: "x", Value: "v", Title: "x", Detail: "", Price: "", Checked: false, Attrs: h.Disabled(), Class: ""})
	if !strings.Contains(got, `<input class="cui-option-input" type="radio" id="x" name="x" value="v" disabled>`) {
		t.Errorf("got %s", got)
	}
}
