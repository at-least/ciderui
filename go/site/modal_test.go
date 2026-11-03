package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestModalRendersTheExample(t *testing.T) {
	example.Check(t, g.Group{
		site.ModalButton{Modal: "models", Class: "cui-btn", Children: g.Text("Show more")},
		site.Modal{ID: "models", Labelledby: "models-title", Page: false, Dark: false, Class: "", Children: g.Group{
			h.H2(h.Class("cui-headline"), h.ID("models-title"), g.Text("Not sure which size to get?")),
			h.P(h.Class("cui-body"), g.Text("The two sizes are the same inside. The bigger one has more screen and lasts longer on a charge.")),
		}},
	}, "site/modal/examples/modal.html")
}

func TestModalRendersThePageExample(t *testing.T) {
	example.Check(t, g.Group{
		site.ModalButton{Modal: "trade-in", Class: "cui-btn", Children: g.Text("Read more about Cider Trade In")},
		site.Modal{ID: "trade-in", Labelledby: "trade-in-title", Page: true, Dark: false, Class: "", Children: g.Group{
			h.P(h.Class("cui-eyebrow"), g.Text("Cider Trade In")),
			h.H2(h.Class("cui-headline"), h.ID("trade-in-title"), g.Text("Get credit for the device you have.")),
			h.P(h.Class("cui-body"), g.Text("Add a trade-in when you choose something new. Once we have your device and have checked it, its value goes back to your payment method.")),
		}},
	}, "site/modal/examples/page.html")
}

func TestModalItsDarkOneTakesCuiDark(t *testing.T) {
	got := example.Render(site.Modal{ID: "m5", Labelledby: "m5-title", Page: true, Dark: true, Class: "", Children: g.Text("x")})
	if !strings.HasPrefix(got, `<dialog class="cui-modal cui-dark" data-variant="page" id="m5" aria-labelledby="m5-title">`) {
		t.Errorf("got %s", got)
	}
}

func TestModalButtonEscapesItsAttributes(t *testing.T) {
	got := example.Render(site.ModalButton{Modal: `a"b`, Class: "x<y&z", Children: g.Text("Open")})
	want := `<button class="x&lt;y&amp;z" type="button" commandfor="a&quot;b" command="show-modal">Open</button>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestModalAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "needs an ID", func() {
		example.Render(site.Modal{ID: "", Labelledby: "t", Page: false, Dark: false, Class: "", Children: g.Text("x")})
	})
}

func TestModalNoHeadlinePanics(t *testing.T) {
	mustPanic(t, "Labelledby", func() {
		example.Render(site.Modal{ID: "m", Labelledby: "", Page: false, Dark: false, Class: "", Children: g.Text("x")})
	})
}

func TestModalButtonWithNoModalPanics(t *testing.T) {
	mustPanic(t, "needs its Modal's id", func() {
		example.Render(site.ModalButton{Modal: "", Class: "cui-btn", Children: g.Text("x")})
	})
}

func TestModalDarkWithoutPagePanics(t *testing.T) {
	mustPanic(t, "site.Modal is Dark only with Page", func() {
		example.Render(site.Modal{ID: "m", Labelledby: "m-title", Page: false, Dark: true, Class: "", Children: nil})
	})
}
