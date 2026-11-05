package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
)

func TestEachPartWritesThePagesClass(t *testing.T) {
	const page = example.PageClass
	example.PagesClass(t, ".", []string{"AccordionItem", "ModalButton", "Post", "ProductNavItem", "Story"}, []example.ClassRow{
		{Root: "cui-accordion", Node: site.Accordion{Large: false, Class: page, Items: nil}},
		{Root: "cui-accordion-item", Node: site.Accordion{Large: false, Class: "", Items: []site.AccordionItem{{Question: "x", Open: false, Class: page, Children: nil}}}},
		{Root: "cui-article", Node: site.Article{Category: "", Date: nil, Headline: "x", Subhead: "", Body: nil, Class: page}},
		{Root: "cui-card", Node: site.Card{Href: "/x/", Media: nil, Eyebrow: "", Headline: "x", Description: "", Price: "", Class: page}},
		{Root: "cui-checkbox", Node: site.Checkbox{ID: "x", Name: "x", Label: "x", Checked: false, Attrs: nil, Class: page}},
		{Root: "cui-field", Node: site.Field{ID: "x", Type: "text", Name: "x", Label: "x", Autocomplete: "", Value: "", Help: "", Error: "", Prefix: "", Attrs: nil, Class: page}},
		{Root: "cui-field", Node: site.Select{ID: "x", Name: "x", Label: "x", Placeholder: "x", Options: models, Selected: "", Error: "", Attrs: nil, Class: page}},
		{Root: "cui-field", Node: site.TextArea{ID: "x", Name: "x", Label: "x", Value: "", Error: "", Attrs: nil, Class: page}},
		{Root: "cui-footer", Node: site.Footer{Notes: nil, NotesAriaLabel: "", Directory: nil, DirectoryAriaLabel: "", Shop: nil, Copyright: g.Text("©"), Links: nil, Locale: nil, Class: page}},
		{Root: "cui-footnote", Node: site.Footnote{Href: "#f", Mark: "1", AriaLabel: "x", Class: page}},
		{Root: "cui-gallery", Node: site.Gallery{ID: "g", Items: nil, Paddles: nil, Current: 0, Class: page}},
		{Root: "cui-globalnav", Node: site.GlobalNav{Logo: globalNavLogo(), Home: "/", LogoAriaLabel: "x", Links: globalNavLinks, Search: nil, Bag: nil, Text: globalNavText, Clear: false, Class: page}},
		{Root: "cui-link", Node: site.MoreLink{Href: "/x", Label: "x", Class: page}},
		{Root: "cui-localnav", Node: site.LocalNav{Title: "x", Href: "/x/", Links: nil, Current: "", Button: nil, AriaLabel: "x", MenuAriaLabel: "", Class: page}},
		{Root: "cui-modal", Node: site.Modal{ID: "m", Labelledby: "t", Variant: site.ModalVariantSheet, CloseAriaLabel: "x", Class: page, Children: nil}},
		{Root: "cui-option", Node: site.OptionTile{ID: "x", Name: "x", Value: "x", Title: "x", Detail: "", Price: "", Checked: false, Attrs: nil, Class: page}},
		{Root: "cui-post", Node: site.Posts{Small: false, Heading: "x", Class: "", Items: []site.Post{{Href: "/x/", Media: nil, Category: "", Date: nil, Headline: "x", Class: page}}}},
		{Root: "cui-posts", Node: site.Posts{Small: false, Heading: "x", Class: page, Items: nil}},
		{Root: "cui-shelf cui-productnav", Node: site.ProductNav{ID: "p", AriaLabel: "x", Paddles: nil, Class: page, Items: nil}},
		{Root: "cui-quicklinks", Node: site.QuickLinks{ID: "q", Title: "x", Links: nil, Class: page}},
		{Root: "cui-ribbon", Node: site.Ribbon{Class: page, Children: nil}},
		{Root: "cui-section", Node: site.Section{Alt: false, Large: false, Class: page, Children: nil}},
		{Root: "cui-section-header", Node: site.SectionHeader{Headline: "x", Links: nil, Class: page}},
		{Root: "cui-shelf", Node: site.Shelf{ID: "x", Headline: "x", Subhead: "", Paddles: nil, Class: page, Items: nil}},
		{Root: "cui-skip cui-btn", Node: site.SkipLink{Href: "#main", Label: "x", Class: page}},
		{Root: "cui-stat", Node: site.Stat{Large: false, Before: "", Value: "x", After: "", Class: page}},
		{Root: "cui-stories", Node: site.Stories{Heading: "x", Class: page, Items: nil}},
		{Root: "cui-story", Node: site.Stories{Heading: "x", Class: "", Items: []site.Story{{Size: site.StorySizeSmall, Href: "/x/", Media: nil, Category: "", Date: nil, Headline: "x", Class: page}}}},
		{Root: "cui-swatch", Node: site.Swatch{ID: "x", Name: "x", Value: "x", AriaLabel: "x", Checked: false, Attrs: nil, Class: page, Children: nil}},
		{Root: "cui-tile", Node: site.Tile{Large: false, Rounded: false, Link: nil, Media: nil, Headline: "x", Subhead: "", Callout: "", Class: page, Children: nil}},
		{Root: "", Node: site.ProductNav{ID: "p", AriaLabel: "x", Paddles: nil, Class: "", Items: []site.ProductNavItem{{Href: "/x", Name: "x", Class: page, Children: g.Text("x")}}}},
	})
}
