// Command apple-account reads Apple Account's pages (account.apple.com), read and
// checked by internal/signedin's command line, App Store Connect's (../asc):
//
//	go run ./reference/apple-account --login | --run <spec> | <spec> | --save <spec> <file> | --check <spec>
package main

import (
	"github.com/at-least/ciderui/internal/signedin"
	"github.com/at-least/ciderui/reference/apple-account/pages"
	"github.com/at-least/ciderui/reference/apple-account/results"
)

func main() { signedin.CLI("apple-account", "https://account.apple.com", pages.Specs, results.Pages) }
