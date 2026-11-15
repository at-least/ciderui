package pages

import "github.com/at-least/ciderui/internal/pagespec"

func touchSearch() pagespec.Spec {
	return touched(search(), []string{"nav", "nav.phone", "results.link"}, nil)
}
