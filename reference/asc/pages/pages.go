// Package pages is App Store Connect's specs, by the name a command line takes.
package pages

import "github.com/at-least/ciderui/internal/signedin"

// Specs are App Store Connect's specs, each its results/<name>.json.
var Specs = map[string]signedin.Spec{
	"analytics":      analytics,
	"analytics-page": analyticsPage,
	"apps":           apps,
	"button":         button,
	"controlrow":     controlrow,
	"crashes":        crashes,
	"head":           head,
	"hover-button":   hoverButton,
	"hover-date":     hoverDate,
	"hover-link":     hoverLink,
	"hover-menu":     hoverMenu,
	"hover-row":      hoverRow,
	"media":          media,
	"testflight":     testflight,
	"tokens":         tokens,
	"version":        version,
}
