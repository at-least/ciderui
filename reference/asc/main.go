// Command asc reads App Store Connect's pages in a Chrome the owner signed in to (no
// headless browser can sign in): its pages hold the account's data. Its command line
// is internal/signedin's, over pages' specs and results/:
//
//	go run ./reference/asc --login | --run <spec> | <spec> | --save <spec> <file> | --check <spec>
package main

import (
	"github.com/at-least/ciderui/internal/signedin"
	"github.com/at-least/ciderui/reference/asc/pages"
	"github.com/at-least/ciderui/reference/asc/results"
)

func main() { signedin.CLI("asc", "https://appstoreconnect.apple.com", pages.Specs, results.Pages) }
