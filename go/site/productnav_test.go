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
	items := []site.Link{
		{Text: "Laptops", Href: "/shop/laptops"},
		{Text: "Phones", Href: "/shop/phones"},
		{Text: "Tablets", Href: "/shop/tablets"},
		{Text: "Watches", Href: "/shop/watches"},
		{Text: "Audio", Href: "/shop/audio"},
		{Text: "Accessories", Href: "/shop/accessories"},
	}
	example.Check(t, site.ProductNav{ID: "products", AriaLabel: "Products", Paddles: false, Class: "", Children: g.Map(items, func(item site.Link) g.Node {
		image := h.Img(h.Class("cui-productnav-image"), h.Src("data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyMDAiIGhlaWdodD0iMTMwIi8+"), h.Width("200"), h.Height("130"), h.Alt(""))
		return site.ProductNavItem{Href: item.Href, Name: item.Text, Class: "", Children: image}
	})}, "site/productnav/examples/productnav.html")
}

func TestProductNavItemAnEmptyNamePanics(t *testing.T) {
	mustPanic(t, "needs a Name", func() {
		example.Render(site.ProductNavItem{Href: "/x", Name: "", Class: "", Children: g.Text("x")})
	})
}

func TestProductNavItemTakesThePagesUtilities(t *testing.T) {
	got := example.Render(site.ProductNavItem{Href: "/x", Name: "x", Class: "hidden", Children: g.Text("x")})
	if !strings.HasPrefix(got, `<li class="hidden"><a class="cui-productnav-item"`) {
		t.Errorf("got %s", got)
	}
}

func TestProductNavAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "ProductNav needs an ID", func() {
		example.Render(site.ProductNav{ID: "", AriaLabel: "Products", Paddles: false, Class: "", Children: g.Text("")})
	})
}

func TestProductNavAnEmptyAriaLabelPanics(t *testing.T) {
	mustPanic(t, "ProductNav needs an AriaLabel", func() {
		example.Render(site.ProductNav{ID: "p", AriaLabel: "", Paddles: false, Class: "", Children: g.Text("")})
	})
}
