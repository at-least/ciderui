package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func TestSearchFieldRendersTheExamples(t *testing.T) {
	field := site.SearchField{Name: "q", Label: "Search apple.com", Value: "", ClearHref: "#search", ClearAriaLabel: "Clear Search", Attrs: nil}
	example.Check(t, field, "site/field/examples/search.html")
	field.Value = "airpods"
	example.Check(t, field, "site/field/examples/search-filled.html")
}

func TestSearchFieldPanics(t *testing.T) {
	example.Panics(t, func() site.SearchField {
		return site.SearchField{Name: "q", Label: "Search", Value: "", ClearHref: "/search", ClearAriaLabel: "Clear", Attrs: nil}
	}, []example.Break[site.SearchField]{
		{Want: "site.SearchField needs a Name", Change: func(f *site.SearchField) { f.Name = "" }},
		{Want: "site.SearchField needs a Label", Change: func(f *site.SearchField) { f.Label = "" }},
		{Want: "site.SearchField needs a ClearHref and a ClearAriaLabel", Change: func(f *site.SearchField) { f.ClearHref = "" }},
		{Want: "site.SearchField needs a ClearHref and a ClearAriaLabel", Change: func(f *site.SearchField) { f.ClearAriaLabel = "" }},
	})
}
