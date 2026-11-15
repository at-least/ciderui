package pages

import "github.com/at-least/ciderui/internal/pagespec"

func searchDark() pagespec.Spec {
	s := search()
	s.URL = "https://www.apple.com/newsroom/"
	s.Widths = []int{1600, 390}
	return s
}
