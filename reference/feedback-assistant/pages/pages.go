// Package pages is Feedback Assistant's specs, by the name a command line takes.
package pages

import "github.com/at-least/ciderui/internal/signedin"

// Specs are Feedback Assistant's specs, each its results/<name>.json.
var Specs = map[string]signedin.Spec{"item": item, "list": list}
