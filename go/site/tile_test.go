package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

const (
	tileCircle = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI0MDAiIGhlaWdodD0iNTAwIj48cmVjdCB3aWR0aD0iNDAwIiBoZWlnaHQ9IjUwMCIgZmlsbD0iI2U4ZThlZCIvPjxjaXJjbGUgY3g9IjIwMCIgY3k9IjIzMCIgcj0iMTEwIiBmaWxsPSIjZDJkMmQ3Ii8+PC9zdmc+"
	tileCase   = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI0MDAiIGhlaWdodD0iNTAwIj48cmVjdCB3aWR0aD0iNDAwIiBoZWlnaHQ9IjUwMCIgZmlsbD0iI2Y1ZjVmNyIvPjxyZWN0IHg9IjExMCIgeT0iMTQwIiB3aWR0aD0iMTgwIiBoZWlnaHQ9IjIyMCIgcng9IjM2IiBmaWxsPSIjZDJkMmQ3Ii8+PC9zdmc+"
	tileBand   = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI0MDAiIGhlaWdodD0iNTAwIj48cmVjdCB3aWR0aD0iNDAwIiBoZWlnaHQ9IjUwMCIgZmlsbD0iI2RmZTdmMiIvPjxyZWN0IHk9IjMyMCIgd2lkdGg9IjQwMCIgaGVpZ2h0PSIxODAiIGZpbGw9IiNjM2QxZTYiLz48L3N2Zz4="
)

func TestTilesRendersTheExamples(t *testing.T) {
	example.Check(t, site.Tiles{Items: []site.Tile{
		{Href: "#airpods-pro", Image: tileCircle, Title: "AirPods Pro 3", Detail: "$249.00"},
		{Href: "#airpods", Image: tileCase, Title: "AirPods 5 with Wireless Charging Case", Detail: "$179.00"},
		{Href: "#case", Image: tileBand, Title: "Silicone Case", Detail: ""},
	}}, "site/tile/examples/tiles.html")
}

func TestTilesPanics(t *testing.T) {
	example.Panics(t, func() site.Tiles {
		return site.Tiles{Items: []site.Tile{{Href: "#x", Image: "x.png", Title: "x", Detail: ""}}}
	}, []example.Break[site.Tiles]{
		{Want: "site.Tiles needs Items: a page with none says so in its own words", Change: func(t *site.Tiles) { t.Items = nil }},
		{Want: "site.Tile needs an Href, an Image and a Title", Change: func(t *site.Tiles) { t.Items[0].Href = "" }},
		{Want: "site.Tile needs an Href, an Image and a Title", Change: func(t *site.Tiles) { t.Items[0].Image = "" }},
		{Want: "site.Tile needs an Href, an Image and a Title", Change: func(t *site.Tiles) { t.Items[0].Title = "" }},
	})
}
