package site

// @trap: what carries its own article ("an AriaLabel"): a vowel rule would
// write "an URL"
func need(part, what, value string) {
	if value == "" {
		panic(part + " needs " + what)
	}
}
