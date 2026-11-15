// Command developer-account reads the Apple Developer account's pages
// (developer.apple.com/account), read and checked by internal/signedin's command
// line, App Store Connect's (../asc):
//
//	go run ./reference/developer-account --login | --run <spec> | <spec> | --save <spec> <file> | --check <spec>
package main

import (
	"github.com/at-least/ciderui/internal/signedin"
	"github.com/at-least/ciderui/reference/developer-account/pages"
	"github.com/at-least/ciderui/reference/developer-account/results"
)

func main() {
	signedin.CLI("developer-account", "https://developer.apple.com", pages.Specs, results.Pages)
}
