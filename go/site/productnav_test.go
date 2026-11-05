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

func productNavOf(item site.ProductNavItem) site.ProductNav {
	return site.ProductNav{ID: "p", AriaLabel: "Products", Paddles: nil, Class: "", Items: []site.ProductNavItem{item}}
}

func TestProductNavItemAnEmptyNamePanics(t *testing.T) {
	mustPanic(t, "needs a Name", func() {
		example.Render(productNavOf(site.ProductNavItem{Href: "/x", Name: "", Class: "", Children: g.Text("x")}))
	})
}

func TestProductNavItemTakesThePagesUtilities(t *testing.T) {
	got := example.Render(productNavOf(site.ProductNavItem{Href: "/x", Name: "x", Class: "hidden", Children: g.Text("x")}))
	if !strings.Contains(got, `<li class="hidden"><a class="cui-productnav-item"`) {
		t.Errorf("got %s", got)
	}
}

func TestProductNavAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "ProductNav needs an ID", func() {
		example.Render(site.ProductNav{ID: "", AriaLabel: "Products", Paddles: nil, Class: "", Items: nil})
	})
}

func TestProductNavAnEmptyAriaLabelPanics(t *testing.T) {
	mustPanic(t, "ProductNav needs an AriaLabel", func() {
		example.Render(site.ProductNav{ID: "p", AriaLabel: "", Paddles: nil, Class: "", Items: nil})
	})
}

func TestProductNavPaddlesNeedTheirWords(t *testing.T) {
	mustPanic(t, "site.ProductNav's Paddles need a PreviousAriaLabel", func() {
		example.Render(site.ProductNav{ID: "p", AriaLabel: "x", Paddles: &site.Paddles{PreviousAriaLabel: "", NextAriaLabel: "x"}, Class: "", Items: nil})
	})
	mustPanic(t, "site.ProductNav's Paddles need a NextAriaLabel", func() {
		example.Render(site.ProductNav{ID: "p", AriaLabel: "x", Paddles: &site.Paddles{PreviousAriaLabel: "x", NextAriaLabel: ""}, Class: "", Items: nil})
	})
}
