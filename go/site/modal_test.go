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
		site.Modal{ID: "models", Labelledby: "models-title", Variant: site.ModalVariantSheet, CloseAriaLabel: "Close", Class: "", Children: g.Group{
			h.H2(h.Class("cui-headline"), h.ID("models-title"), g.Text("Not sure which size to get?")),
			h.P(h.Class("cui-body"), g.Text("The two sizes are the same inside. The bigger one has more screen and lasts longer on a charge.")),
		}},
	}, "site/modal/examples/modal.html")
}

func TestModalRendersThePageExample(t *testing.T) {
	example.Check(t, g.Group{
		site.ModalButton{Modal: "trade-in", Class: "cui-btn", Children: g.Text("Read more about Cider Trade In")},
		site.Modal{ID: "trade-in", Labelledby: "trade-in-title", Variant: site.ModalVariantPage, CloseAriaLabel: "Close", Class: "", Children: g.Group{
			h.P(h.Class("cui-eyebrow"), g.Text("Cider Trade In")),
			h.H2(h.Class("cui-headline"), h.ID("trade-in-title"), g.Text("Get credit for the device you have.")),
			h.P(h.Class("cui-body"), g.Text("Add a trade-in when you choose something new. Once we have your device and have checked it, its value goes back to your payment method.")),
		}},
	}, "site/modal/examples/page.html")
}

func TestModalItsDarkOneTakesCuiDark(t *testing.T) {
	got := example.Render(site.Modal{ID: "m5", Labelledby: "m5-title", Variant: site.ModalVariantPageDark, CloseAriaLabel: "Close", Class: "", Children: g.Text("x")})
	if !strings.HasPrefix(got, `<dialog class="cui-modal cui-dark" data-variant="page" id="m5" aria-labelledby="m5-title">`) {
		t.Errorf("got %s", got)
	}
}

// the dialog focuses its first focusable element on opening: a tabindex="-1"
// button is one, so the curtain comes after the close button
func TestModalItsCloseComesBeforeItsCurtain(t *testing.T) {
	got := example.Render(site.Modal{ID: "m", Labelledby: "t", Variant: site.ModalVariantSheet, CloseAriaLabel: "關閉", Class: "", Children: g.Text("x")})
	first := strings.Index(got, "<button")
	if !strings.HasPrefix(got[first:], `<button class="cui-modal-close" aria-label="關閉">`) {
		t.Errorf("got %s", got)
	}
	if !strings.HasSuffix(got, `<form method="dialog"><button class="cui-modal-curtain" tabindex="-1" aria-hidden="true"></button></form></dialog>`) {
		t.Errorf("got %s", got)
	}
}

func TestModalPanics(t *testing.T) {
	example.Panics(t, func() site.Modal {
		return site.Modal{ID: "m", Labelledby: "t", Variant: site.ModalVariantSheet, CloseAriaLabel: "Close", Class: "", Children: g.Text("x")}
	}, []example.Break[site.Modal]{
		{Want: "site.Modal needs an ID", Change: func(p *site.Modal) { p.ID = "" }},
		{Want: "site.Modal needs the id of its headline (Labelledby)", Change: func(p *site.Modal) { p.Labelledby = "" }},
		{Want: "site.Modal needs a CloseAriaLabel", Change: func(p *site.Modal) { p.CloseAriaLabel = "" }},
	})
}

func TestModalButtonPanics(t *testing.T) {
	example.Panics(t, func() site.ModalButton {
		return site.ModalButton{Modal: "m", Class: "cui-btn", Children: g.Text("x")}
	}, []example.Break[site.ModalButton]{
		{Want: "site.ModalButton needs its Modal's id", Change: func(p *site.ModalButton) { p.Modal = "" }},
	})
}
