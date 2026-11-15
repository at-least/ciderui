package pages

import "github.com/at-least/ciderui/internal/pagespec"

func touchOverlay() pagespec.Spec {
	return touched(overlay(), []string{"overlay", "close.circle-hover"}, nil)
}
