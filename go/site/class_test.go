package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func TestEachPartWritesThePagesClass(t *testing.T) {
	const page = example.PageClass
	example.PagesClass(t, []example.ClassRow{
		{Root: "cui-btn", Node: site.Btn{Href: "/x", Label: "x", Icon: nil, Submit: false, Variant: site.BtnVariantPrimary, Size: site.BtnSizeMedium, Block: false, Attrs: nil, Class: page}},
		{Root: "cui-ribbon", Node: site.Ribbon{Class: page, Children: nil}},
	})
}
