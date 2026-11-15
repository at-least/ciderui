package pages

import "github.com/at-least/ciderui/internal/pagespec"

func touchHome() pagespec.Spec {
	return touched(home(), []string{"nav.link", "nav.logo", "nav.search", "nav.bag", "nav.menutrigger", "button", "button.secondary", "gallery.dot", "footer.sosumi-link", "footer.link", "footer.legal-link", "footer.locale-link", "menu.open", "menu.link"}, nil)
}
