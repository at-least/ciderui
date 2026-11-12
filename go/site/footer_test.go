package site_test

import (
	"io"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func company() g.Node { return h.A(h.Href("/company/"), g.Text("Example Inc.")) }

func TestFooterRendersTheExample(t *testing.T) {
	example.Check(t, site.Footer{Directory: nil, DirectoryAriaLabel: "", Copyright: company()}, "site/footer/examples/footer.html")
}

func directory() []site.FooterSection {
	return []site.FooterSection{
		{Title: "Products", Links: []site.FooterLink{{Label: "Store", Href: "/store/"}, {Label: "Laptops", Href: "/laptops/"}, {Label: "Phones", Href: "/phones/"}}},
		{Title: "Account", Links: []site.FooterLink{{Label: "Your Account", Href: "/account/"}, {Label: "Order History", Href: "/orders/"}}},
	}
}

func TestFooterRendersTheDirectoryExample(t *testing.T) {
	example.Check(t, site.Footer{Directory: directory(), DirectoryAriaLabel: "Directory", Copyright: company()}, "site/footer/examples/directory.html")
}

type counted struct{ n *int }

func (c counted) Render(w io.Writer) error {
	*c.n++
	_, err := io.WriteString(w, "x")
	return err
}

// Copyright is the page's node: rendering one twice may cost, or render
// differently the second time
func TestFooterRendersItsCopyrightOnce(t *testing.T) {
	var copyright int
	example.Render(site.Footer{Directory: nil, DirectoryAriaLabel: "", Copyright: counted{&copyright}})
	if copyright != 1 {
		t.Errorf("the footer rendered the page's copyright %d times, want once", copyright)
	}
}

func TestFooterPanics(t *testing.T) {
	example.Panics(t, func() site.Footer {
		return site.Footer{Directory: directory(), DirectoryAriaLabel: "Directory", Copyright: g.Text("©")}
	}, []example.Break[site.Footer]{
		{Want: "site.Footer needs a Copyright", Change: func(p *site.Footer) { p.Copyright = nil }},
		{Want: "site.Footer has a Directory or a DirectoryAriaLabel without the other", Change: func(p *site.Footer) { p.DirectoryAriaLabel = "" }},
		{Want: "site.Footer has a Directory or a DirectoryAriaLabel without the other", Change: func(p *site.Footer) { p.Directory = nil }},
		{Want: "site.FooterSection needs a Title and Links", Change: func(p *site.Footer) { p.Directory[1].Title = "" }},
		{Want: "site.FooterSection needs a Title and Links", Change: func(p *site.Footer) { p.Directory[1].Links = nil }},
		{Want: "site.FooterSection's Links need their Label and Href", Change: func(p *site.Footer) { p.Directory[1].Links[0].Label = "" }},
		{Want: "site.FooterSection's Links need their Label and Href", Change: func(p *site.Footer) { p.Directory[1].Links[0].Href = "" }},
	})
}
