package pages

import "github.com/at-least/ciderui/internal/pagespec"

func supportDark() pagespec.Spec {
	s := support()
	s.ColorScheme = "dark"
	return s
}
