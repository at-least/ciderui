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
		{Root: "cui-accordion", Node: site.Accordion{Large: false, Class: page, Items: nil}},
		{Root: "cui-accordion-item", Node: site.Accordion{Large: false, Class: "", Items: []site.AccordionItem{{Question: "x", Open: false, Class: page, Children: nil}}}},
		{Root: "cui-field", Node: site.Field{ID: "x", Type: "text", Name: "x", Label: "x", Autocomplete: "", Value: "", Help: "", Error: "", Prefix: "", Attrs: nil, Class: page}},
		{Root: "cui-field", Node: site.Select{ID: "x", Name: "x", Label: "x", Placeholder: "x", Options: models, Selected: "", Error: "", Attrs: nil, Class: page}},
		{Root: "cui-field", Node: site.TextArea{ID: "x", Name: "x", Label: "x", Value: "", Error: "", Attrs: nil, Class: page}},
		{Root: "cui-footer", Node: site.Footer{Notes: nil, NotesAriaLabel: "", Directory: nil, DirectoryAriaLabel: "", Shop: nil, Copyright: g.Text("©"), Links: nil, Locale: nil, Class: page}},
		{Root: "cui-globalnav", Node: site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "x", Links: globalNavLinks, Search: nil, Bag: nil, Text: globalNavText, Clear: false, Class: page}},
		{Root: "cui-link", Node: site.MoreLink{Href: "/x", Label: "x", Class: page}},
		{Root: "cui-option", Node: site.OptionTile{ID: "x", Name: "x", Value: "x", Title: "x", Detail: "", Price: "", Checked: false, Attrs: nil, Class: page}},
		{Root: "cui-ribbon", Node: site.Ribbon{Class: page, Children: nil}},
	})
}
