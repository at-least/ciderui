package pages

import "github.com/at-least/ciderui/internal/pagespec"

func touchIphone() pagespec.Spec {
	return touched(iphone(), []string{"link.more-copy", "link.inline", "paddle", "gallery.dot"}, nil)
}
