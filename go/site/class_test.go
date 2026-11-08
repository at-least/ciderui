package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
)

func TestEachPartWritesThePagesClass(t *testing.T) {
	const page = example.PageClass
	example.PagesClass(t, ".", []string{"AccordionItem", "Page"}, []example.ClassRow{
		{Root: "cui-btn", Node: site.Btn{Href: "/x", Label: "x", Icon: nil, Submit: false, Variant: site.BtnVariantPrimary, Size: site.BtnSizeMedium, Block: false, Attrs: nil, Class: page}},
		{Root: "cui-accordion-item", Node: site.Accordion{Items: []site.AccordionItem{{Question: "x", Class: page, Children: nil}}}},
		{Root: "cui-field", Node: site.Field{ID: "x", Type: "text", Name: "x", Label: "x", Autocomplete: "", Value: "", Error: "", Invalid: false, Help: "", Prefix: "", Attrs: nil, Class: page}},
		{Root: "cui-footer", Node: site.Footer{Copyright: g.Text("©"), Class: page}},
		{Root: "cui-globalnav", Node: site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "x", Links: []site.NavLink{{Label: "x", Href: "/x", Post: false}}, AriaLabel: "x", MenuAriaLabel: "x", CloseAriaLabel: "x", Class: page}},
		{Root: "cui-option", Node: site.Option{ID: "x", Name: "x", Value: "x", Title: "x", Detail: "", Checked: false, Attrs: nil, Class: page}},
		{Root: "cui-ribbon", Node: site.Ribbon{Class: page, Children: nil}},
	})
}
