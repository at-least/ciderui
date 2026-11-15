// Command feedback-assistant reads Feedback Assistant's pages, read and checked by
// internal/signedin's command line, App Store Connect's (../asc):
//
//	go run ./reference/feedback-assistant --login | --run <spec> | <spec> | --save <spec> <file> | --check <spec>
package main

import (
	"github.com/at-least/ciderui/internal/signedin"
	"github.com/at-least/ciderui/reference/feedback-assistant/pages"
	"github.com/at-least/ciderui/reference/feedback-assistant/results"
)

func main() {
	signedin.CLI("feedback-assistant", "https://feedbackassistant.apple.com", pages.Specs, results.Pages)
}
