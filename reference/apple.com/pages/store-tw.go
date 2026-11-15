package pages

import (
	"strings"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The Traditional Chinese store's type scale, `store`'s probes on its page, which
// sets them apart for its language (tokens.css's [lang|="zh"], type.css's :lang(zh)).
func storeTw() pagespec.Spec {
	var probes []pagespec.Probe
	for _, p := range store().Probes {
		if p.Name == "page.body" || strings.HasPrefix(p.Name, "type.t-") {
			probes = append(probes, p)
		}
	}
	return pagespec.Spec{URL: "https://www.apple.com/tw/store", Probes: probes}
}
