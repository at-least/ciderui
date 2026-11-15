package pages

import "github.com/at-least/ciderui/internal/pagespec"

func touchRicmodal() pagespec.Spec {
	return touched(ricmodal(), []string{"modal", "close.icon-states"}, nil)
}
