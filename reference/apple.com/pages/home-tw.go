package pages

import "github.com/at-least/ciderui/internal/pagespec"

func homeTw() pagespec.Spec {
	s := home()
	s.URL = "https://www.apple.com/tw/"
	return s
}
