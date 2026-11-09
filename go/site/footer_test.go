package site_test

import (
	"io"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestFooterRendersTheExample(t *testing.T) {
	example.Check(t, site.Footer{Copyright: h.A(h.Href("/company/"), g.Text("Example Inc."))}, "site/footer/examples/footer.html")
}

type counted struct{ n *int }

func (c counted) Render(w io.Writer) error {
	*c.n++
	_, err := io.WriteString(w, "x")
	return err
}

// Copyright is the page's node: rendering one twice may cost, or render
// differently the second time
func TestFooterItsCopyrightRenderedOnce(t *testing.T) {
	var copyright int
	example.Render(site.Footer{Copyright: counted{&copyright}})
	if copyright != 1 {
		t.Errorf("copyright rendered %d times, want 1", copyright)
	}
}

func TestFooterPanics(t *testing.T) {
	example.Panics(t, func() site.Footer {
		return site.Footer{Copyright: g.Text("©")}
	}, []example.Break[site.Footer]{
		{Want: "site.Footer needs a Copyright", Change: func(p *site.Footer) { p.Copyright = nil }},
	})
}
