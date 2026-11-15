package reference

import (
	"fmt"
)

// SignedInPage is what was read on one page of a signed-in site (App Store Connect,
// Apple Account, the developer account, Feedback Assistant): reference/<site>/
// results/<Name>.go, which internal/signedin writes from what the page handed back.
// It holds Probes or Tokens, as its spec read; the digests are the page's, which
// --check holds the record to.
//
// @external: App Store Connect's pages were read at one width, 1920px
type SignedInPage struct {
	Name, Path, Browser, Date string
	Window                    Window
	// Probes in the order the page handed them back; nil on a tokens record.
	Probes []SignedInProbe
	// Tokens in the order the page handed them back (by name); nil on a probes
	// record.
	Tokens []Token
	// ProbesDigest is the SHA-256 of the spec's probes (or token spec) the page was
	// given, ReaderDigest of the reader's source, SHA256 of the rest of the record:
	// each as the page's canon wrote it.
	ProbesDigest, ReaderDigest, SHA256 [32]byte
}

// Window is the page's innerWidth, innerHeight and devicePixelRatio.
type Window struct {
	Width, Height, PixelRatio float64
}

// SignedInProbe is one probe's reading: its element's name, its computed values in
// the spec's order and, where the spec read it, its box (At is nil).
type SignedInProbe struct {
	Name, HTML string
	Styles     []Style
	Box        *Box
}

// Token is one of the toolkit's tokens, --tb-<…>, as the page computed it.
type Token struct {
	Name, Value string
}

// Value is the measured value of prop on the page's probe.
func (p SignedInPage) Value(name, prop string) (string, error) {
	for _, pr := range p.Probes {
		if pr.Name != name {
			continue
		}
		if prop == "box-height" || prop == "box-width" {
			if pr.Box == nil {
				return "", fmt.Errorf("reference asc %s:%s has no box", p.Name, name)
			}
			if prop == "box-height" {
				return px(pr.Box.Height), nil
			}
			return px(pr.Box.Width), nil
		}
		for _, s := range pr.Styles {
			if s.Name == prop {
				return s.Text, nil
			}
		}
		return "", fmt.Errorf("reference asc %s:%s has no %s", p.Name, name, prop)
	}
	return "", fmt.Errorf("reference asc %s has no probe %s", p.Name, name)
}

// Token is the token --tb-<name> as the page computed it.
func (p SignedInPage) Token(name string) (string, error) {
	for _, t := range p.Tokens {
		if t.Name == "--tb-"+name {
			return t.Value, nil
		}
	}
	return "", fmt.Errorf("reference asc %s has no --tb-%s", p.Name, name)
}
