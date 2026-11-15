package pages

import "github.com/at-least/ciderui/internal/pagespec"

func localnavDark() pagespec.Spec {
	return pagespec.Spec{URL: "https://www.apple.com/macbook-pro/", Widths: []int{1600, 1200, 900, 834, 833, 390}, Probes: localnavProbes()}
}
