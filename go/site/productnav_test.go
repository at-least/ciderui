package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestProductNavRendersTheExample(t *testing.T) {
	item := func(name, href string) site.ProductNavItem {
		image := h.Img(h.Class("cui-productnav-image"), h.Src("data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyMDAiIGhlaWdodD0iMTMwIi8+"), h.Width("200"), h.Height("130"), h.Alt(""))
		return site.ProductNavItem{Href: href, Name: name, Class: "", Children: image}
	}
	example.Check(t, site.ProductNav{ID: "products", AriaLabel: "Products", Paddles: nil, Class: "", Items: []site.ProductNavItem{
		item("Laptops", "/shop/laptops"),
		item("Phones", "/shop/phones"),
		item("Tablets", "/shop/tablets"),
		item("Watches", "/shop/watches"),
		item("Audio", "/shop/audio"),
		item("Accessories", "/shop/accessories"),
	}}, "site/productnav/examples/productnav.html")
}

func validProductNav() site.ProductNav {
	return site.ProductNav{ID: "p", AriaLabel: "Products", Paddles: nil, Class: "", Items: []site.ProductNavItem{{Href: "/x", Name: "x", Class: "", Children: g.Text("x")}}}
}

func TestProductNavPanics(t *testing.T) {
	example.Panics(t, validProductNav, []example.Break[site.ProductNav]{
		{Want: "site.ProductNav needs an ID", Change: func(p *site.ProductNav) { p.ID = "" }},
		{Want: "site.ProductNav needs an AriaLabel", Change: func(p *site.ProductNav) { p.AriaLabel = "" }},
		{Want: "site.ProductNavItem needs a Name", Change: func(p *site.ProductNav) { p.Items[0].Name = "" }},
		{Want: "site.ProductNav's Paddles need a PreviousAriaLabel", Change: func(p *site.ProductNav) { p.Paddles = &site.Paddles{PreviousAriaLabel: "", NextAriaLabel: "x"} }},
		{Want: "site.ProductNav's Paddles need a NextAriaLabel", Change: func(p *site.ProductNav) { p.Paddles = &site.Paddles{PreviousAriaLabel: "x", NextAriaLabel: ""} }},
	})
}

func TestProductNavPaddlesControlItsRow(t *testing.T) {
	nav := validProductNav()
	nav.Paddles = &site.Paddles{PreviousAriaLabel: "Previous", NextAriaLabel: "Next"}
	got := example.Render(nav)
	for _, want := range []string{
		`<button class="cui-gallery-paddle cui-shelf-paddle" type="button" aria-controls="p" aria-label="Previous" disabled>`,
		`<button class="cui-gallery-paddle cui-shelf-paddle" type="button" aria-controls="p" aria-label="Next">`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("no %s in %s", want, got)
		}
	}
}
