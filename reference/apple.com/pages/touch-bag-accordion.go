package pages

import "github.com/at-least/ciderui/internal/pagespec"

func touchBagAccordion() pagespec.Spec {
	return touched(bagAccordion(), []string{"accordion.open", "compact.icon-hover"}, nil)
}
