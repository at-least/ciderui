package pages

import "github.com/at-least/ciderui/internal/pagespec"

func touchStore() pagespec.Spec {
	return touched(store(), []string{"button", "button.secondary", "button.neutral", "button.secondary-neutral", "card", "card.dark", "card.link"}, []string{"hover", "focus-within"})
}
