package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func TestOptionTileRendersTheExamples(t *testing.T) {
	example.Check(t, site.OptionTile{ID: "storage-256", Name: "storage", Value: "256", Title: "256GB", Detail: "", Price: "From $1099", Checked: false, Class: ""}, "site/option/examples/option.html")
	example.Check(t, site.OptionTile{ID: "model-pro", Name: "model", Value: "pro", Title: "Cider Phone 3", Detail: "6.3-inch display", Price: "", Checked: true, Class: ""}, "site/option/examples/checked.html")
}

func TestOptionTileAnEmptyTitlePanics(t *testing.T) {
	mustPanic(t, "needs a Title", func() {
		example.Render(site.OptionTile{ID: "x", Name: "x", Value: "x", Title: "", Detail: "", Price: "", Checked: false, Class: ""})
	})
}

func TestOptionTileAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "OptionTile needs an ID", func() {
		example.Render(site.OptionTile{ID: "", Name: "x", Value: "x", Title: "x", Detail: "", Price: "", Checked: false, Class: ""})
	})
}
