package pages

import "github.com/at-least/ciderui/internal/pagespec"

func touchLocalnav() pagespec.Spec {
	return touched(localnav(), []string{"localnav.link", "localnav.button", "localnav.chevron-image", "tray.open", "tray.link"}, nil)
}
