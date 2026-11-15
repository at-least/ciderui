package pages

import "github.com/at-least/ciderui/internal/pagespec"

// @external: appstore-ratings' bars with the system's appearance dark: the App Store's
// web page follows it.
func appstoreRatingsDark() pagespec.Spec {
	s := appstoreRatings()
	s.ColorScheme = "dark"
	s.Widths = []int{1200, 390}
	return s
}
