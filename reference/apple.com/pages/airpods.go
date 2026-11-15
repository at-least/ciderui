package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

func airpods() pagespec.Spec {
	return pagespec.Spec{
		URL:  "https://www.apple.com/airpods-pro/",
		Wait: 2500,
		Probes: []pagespec.Probe{
			{Name: "stat.content", Select: "main .stat .stat-content", Props: []string{"padding-top", "padding-bottom", "margin-top", "margin-bottom"}},
			{Name: "stat.caption", Select: "main .stat .stat-caption", Props: slices.Concat(TYPE, []string{"display"})},
			{Name: "stat.value", Select: "main .stat .stat-value", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom", "display"})},
			{Name: "footnote", Select: "main sup.footnote", Props: slices.Concat([]string{"font-feature-settings"}, TYPE, []string{"vertical-align", "position", "top"})},
			{Name: "footnote.link", Select: "main sup.footnote a", Props: []string{"font-feature-settings", "color", "text-decoration-line", "font-size", "vertical-align"}},
		},
	}
}
