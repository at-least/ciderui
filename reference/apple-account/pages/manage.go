// Package pages is Apple Account's specs, by the name a command line takes.
package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

// Specs are Apple Account's specs, each its results/<name>.json.
var Specs = map[string]signedin.Spec{"manage": manage}

var manage = func() signedin.Spec {
	pad := []string{"padding-top", "padding-right", "padding-bottom", "padding-left"}
	edge := []string{"border-top-width", "border-top-style", "border-top-color", "border-top-left-radius", "background-color", "box-shadow"}
	return signedin.Spec{
		Path: "/account/manage",
		Probes: []signedin.Probe{
			{Name: "page", Select: "body", Props: signedin.Page},
			{Name: "nav", Select: ".layout__nav", Props: slices.Concat([]string{"width"}, pad), Box: true},
			{Name: "person.picture", Select: ".person-info__image svg", Props: []string{"width", "height", "margin-bottom"}, Box: true},
			{Name: "person.title", Select: ".person-info__name", Props: slices.Concat(signedin.Type, []string{"margin-top"}), Box: true},
			{Name: "person.line", Select: ".person-info__email", Props: signedin.Type, Box: true},
			{Name: "list", Select: "ul.side-nav__list", Props: slices.Concat([]string{"margin-top"}, pad), Box: true},
			{Name: "list.item", Select: "li.side-nav__item:not(.side-nav__item--current)", Props: slices.Concat([]string{"height", "margin-bottom"}, pad), Box: true},
			{Name: "list.link", Select: ".side-nav__item:not(.side-nav__item--current) .side-nav__link", Props: slices.Concat(signedin.Type, signedin.Link, pad, []string{"display"}), Box: true},
			{Name: "list.current", Select: ".side-nav__item--current .side-nav__link", Props: slices.Concat(signedin.Type, signedin.Link, pad), Box: true},
			{Name: "header", Select: "header.page-header", Props: slices.Concat(pad, []string{"margin-bottom"}), Box: true},
			{Name: "title", Select: "h1.page-title", Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "description", Select: ".page-description", Props: slices.Concat(signedin.Type, []string{"margin-top"}), Box: true},
			{Name: "column", Select: ".section .row > .column", Props: slices.Concat([]string{"width"}, pad, []string{"margin-bottom"}), Box: true},
			{Name: "card.button", Select: "button.button-expand", Props: slices.Concat(pad, edge, []string{"width", "height", "cursor"}), Box: true},
			{Name: "card", Select: ".card.card-border", Props: slices.Concat(edge, []string{"border-left-color", "width", "height"}), Box: true},
			{Name: "card.top", Select: ".card .card-top", Props: slices.Concat(pad, []string{"margin-top"}), Box: true},
			{Name: "card.title", Select: ".card .card-title", Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "card.line", Select: ".card .card-line", Props: slices.Concat(signedin.Type, []string{"margin-top"}), Box: true},
			{Name: "card.glyph", Select: ".card .card-right-element svg", Props: []string{"width", "height", "color", "fill"}, Box: true},
		},
	}
}()
