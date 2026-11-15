// Package pages is the Apple Developer account's specs, by the name a command line
// takes.
package pages

import "github.com/at-least/ciderui/internal/signedin"

// Specs are the developer account's specs, each its results/<name>.json.
var Specs = map[string]signedin.Spec{"keys": keys, "overview": overview}
