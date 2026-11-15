package pages

import "github.com/at-least/ciderui/internal/pagespec"

func searchSuggestionsDark() pagespec.Spec {
	s := searchSuggestions()
	s.URL = "https://www.apple.com/newsroom/"
	s.Widths = []int{1600, 390}
	return s
}
